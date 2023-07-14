package main
import "fmt"

type Teacher struct{
	Name string
	Age int
	School string
}

func main() {
	var t1 Teacher
	fmt.Println(t1)
	t1.Name = "xxx"
	t1.Age = 18
	t1.School = "yyy"
	fmt.Println(t1)

	var t2 Teacher = Teacher{}
	fmt.Println(t2)


	var t3 Teacher = Teacher{"xxx", 18, "yyy"}
	fmt.Println(t3)


	var t4 *Teacher = new(Teacher)
	(*t4).Name = "xxx"
	(*t4).Age = 18
	t4.School = "yyy" // p->school = "xxx" go语言的简化方式
	fmt.Println(t4)

	var t5 *Teacher = &Teacher{"xxx", 18, "yyy"}
	fmt.Println(t5)
}