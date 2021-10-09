package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	var (
		mac   string
		cpuid string
	)
	mac = getMac()
	addrs := getMacAddrs()
	fmt.Println(addrs)
	fmt.Println("mac:" + mac)
	fmt.Println("cpuid:" + cpuid)
	myMap := make(map[string]string)
	myMap["cpuid"] = cpuid

	// 将结构体解析为字符串
	b, e := json.Marshal(myMap)
	if nil != e {
		fmt.Println(e)
	}

	// 将字符串反解析为结构体
	json.Unmarshal(b, &myMap)

	// 写入文件
	f, err := os.Create("sany-cpu.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(string(b))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getMac() string {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	inter := interfaces[0]
	mac := inter.HardwareAddr.String() //获取本机MAC地址
	//	fmt.Println("MAC = ", mac)
	return mac
}

func getMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

func getCpuId() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	//	fmt.Println(string(out))
	str := string(out)
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	return str[11:]
}
