package ch11

import "testing"

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"1", "Aoa", 11}
	e1 := Employee{Name: "Bob", Age: 12}
	e2 := new(Employee) // 创建的是指针
	e2.Id = "3"
	e2.Name = "Coc"
	e2.Age = 13
	t.Log(e1.Id)
	t.Logf("e is %T", e)
	t.Logf("&e is %T", &e) // 加上取值符 后也是取指针了
	t.Logf("e2 is %T", e2)
}

// 行为的封装，
// 两种 方法 1 实例  地址不同 有数据复制，  2指针 所有数据是同一块
//  %x ,unsafe.Pointer(&e.Name)
