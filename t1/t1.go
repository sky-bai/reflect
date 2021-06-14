package main

import "fmt"

type student struct {
	name string
	age string
}

func (s student)changename(name string)  {
	s.name = name
	fmt.Println(s.name)
}
func (s *student)changeage(age string)  {
	s.age = age
	fmt.Println(s.age)

}
func main()  {
	s1:=student{
		name: "bai",
		age: "18",
	}
	s2:=&student{
		name: "wxy",
		age: "18",
	}
	fmt.Println(s1.name)
	fmt.Println(s1.age)

	s1.changename("baibai")
	s1.changeage("20")


	fmt.Println(s2.name)
	fmt.Println(s2.age)
	s2.changename("wxywxy")
	s2.changeage("20")
}

