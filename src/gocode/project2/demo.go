package main

func func6() {
	panic("division by zero")
}
func func5() {
	func6()
}
func func4() {
	func5()
}
func func3() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	func4()
}
func func2() {
	func3()
}

func func1() {
	func2()
}

func main() {
	func1()
}
