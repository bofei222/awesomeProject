package main

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/google/martian/log"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "10.162.138.10:6379", // 你的Redis服务器地址
		Password: "zTGBBRP3dx2VG8su",   // 如果有密码，添加密码
		DB:       2,                    // 默认数据库
	})

	sourceDirectory = "/home/scada/data/geli/cms"
	listName        = "event-center-list"
	watcher         *fsnotify.Watcher
	mutex           sync.Mutex
)

func init() {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		os.Exit(1)
	}
}

func watchDirectory() {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			switch {
			case event.Op&fsnotify.Create == fsnotify.Create:
				handleFile(event.Name)
			case event.Op&fsnotify.Rename == fsnotify.Rename:
				handleFile(event.Name)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error watching directory:", err)
		}
	}
}

func handleFile(filePath string) {
	mutex.Lock()
	defer mutex.Unlock()

	// 等待文件完全生成
	time.Sleep(1 * time.Second)

	// 获取文件名
	_, fileName := filepath.Split(filePath)

	// 读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 解析JSON内容为map
	var contentMap map[string]interface{}
	err = json.Unmarshal(content, &contentMap)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	myUUID := uuid.New().String()
	log.Infof("fileName:%s uuid:%s", fileName, myUUID)
	// 生成UUID作为message_id
	contentMap["message_id"] = myUUID
	contentMap["relevant_id"] = myUUID
	contentMap["part"] = "cms部件"

	// 将map转为JSON字符串
	updatedContent, err := json.Marshal(contentMap)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// 写入Redis List
	key := listName
	err = redisClient.RPush(context.Background(), key, string(updatedContent)).Err()
	if err != nil {
		fmt.Println("Error pushing to Redis List:", err)
		return
	}

	// 删除文件
	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Printf("File '%s' processed and moved to Redis.\n", fileName)
}

func main() {
	// 检查并创建sourceDirectory目录
	if _, err := os.Stat(sourceDirectory); os.IsNotExist(err) {
		err = os.MkdirAll(sourceDirectory, 0755)
		if err != nil {
			fmt.Println("Error creating source directory:", err)
			os.Exit(1)
		}
	}

	err := watcher.Add(sourceDirectory)
	if err != nil {
		fmt.Println("Error watching directory:", err)
		os.Exit(1)
	}

	go watchDirectory()

	// 阻塞主程序
	select {}
}
