package tool

import (
	"fmt"
	"testing"
)

// 将不删除的 前移 ， 记录有效的元素坐标， 最后截取
func TestSliceRemove(t *testing.T) {
	a := []int{1, 5, 4, 3, 2}
	j := 0

	for _, v := range a {
		if v <= 3 {
			a[j] = v
			j++
		}
	}
	a = a[:j]
	fmt.Println(a)
}

func TestDiff(t *testing.T) {
	slice1 := []string{"5", "20", "c", "100"}
	slice2 := []string{"ff20", "c", "ff5"}
	difference := Difference(slice1, slice2)
	fmt.Println(difference)
	j := 0
	for _, e := range difference {
		if !Contains(slice2, "ff"+e) {
			difference[j] = e
			j++
		}
	}
	difference = difference[:j]
	fmt.Println(difference)
}
func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//求差集 slice1-并集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++ // m[v]=1
	}

	for _, value := range slice1 {
		times := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//求交集
func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}
