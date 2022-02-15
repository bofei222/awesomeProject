package client

import (
	"awesomeProject/src/geek_tutorial/ch15/series"
	cm "github.com/orcaman/concurrent-map"
	"testing"
)

func TestPackage(t *testing.T) {

	t.Log(series.GetFibonacci(5))
	concurrentMap := cm.New()
	concurrentMap.IsEmpty()
}

// go get -u 本地有了 也要从网上更新 最新版本
// 不要把src 提交到 github
//https://github.com/orcaman/concurrent-map.git
// 不能管理特定版本 依赖   1.5之后 vendor
