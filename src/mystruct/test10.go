package main

func main() {
	i := 10
	j := i
	println(j)
	j = 11
	println(i)

	println("-----------------")
	p := &i
	println(p)
	println(*p)
	*p = 12
	println(i)

	println("-----------------")
	j = 2701
	p = &j
	println(*p)
	*p = *p / 37
	println(j)
}
