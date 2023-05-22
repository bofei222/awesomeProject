package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

// 省略了err的判断，正常情况需要有，否则可能panic
func main() {

	expr, _ := govaluate.NewEvaluableExpression("foo > 0")
	parameters := make(map[string]interface{})
	parameters["foo"] = -1
	result, _ := expr.Evaluate(parameters) // 即 -1 > 0
	fmt.Println(result)                    // false

	expr, _ = govaluate.NewEvaluableExpression("pingPassFlag=='1'")
	parameters = make(map[string]interface{})
	parameters["pingPassFlag"] = "1"
	result, _ = expr.Evaluate(parameters) // 即 -1 > 0
	fmt.Println(result)

	//expr, _ = govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")
	//parameters = make(map[string]interface{})
	//parameters["requests_made"] = 100
	//parameters["requests_succeeded"] = 80
	//result, _ = expr.Evaluate(parameters) // 即 80 >= 90
	//fmt.Println(result)                   // false
	//
	//expr, _ = govaluate.NewEvaluableExpression("(mem_used / total_mem) * 100")
	//parameters = make(map[string]interface{})
	//parameters["total_mem"] = 1024
	//parameters["mem_used"] = 512
	//result, _ = expr.Evaluate(parameters) // 即 512/1024 * 100
	//fmt.Println(result)                   // 50
}
