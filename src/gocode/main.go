package main

import (
	"GoBaseCode/src/gocode/util"
	"fmt"
)

var projectVersion = "1.0"

func init() {
	fmt.Println("mainVersion: ", projectVersion)
}

func main() {
	fmt.Println("运行main()", util.Int2str(1))

}
