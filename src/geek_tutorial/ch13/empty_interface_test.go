package ch13

import "fmt"

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
	}
	if i, ok := p.(string); ok {
		fmt.Println("String", i)
	}

	fmt.Println("un know Type")

	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Un know Type")
	}

}

// 倾向于 定义 较小的接口，  共用性更强
// 较大的接口定义，可以由多个小接口组合
