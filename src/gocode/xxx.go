package main

import (
	"log"
	"os"
)

func main() {
	//输出日志
	//log.Print()
	//log.Println()
	//log.Printf()

	//输出日志并panic异常
	//log.Panic()
	//log.Panicln()
	//log.Panicf()

	//输出日志并退出程序
	//log.Fatal()
	//log.Fatalln()
	//log.Fatalf()

	log.Flags()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.New(os.Stdout, "xxx", 1)
}
