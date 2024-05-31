//package main
//
//import (
//	"GoBaseCode/src/gocode/util"
//	"fmt"
//)
//
//var projectVersion = "1.0"
//
//func init() {
//	fmt.Println("mainVersion: ", projectVersion)
//}
//
//func main() {
//	fmt.Println("运行main()", util.Int2str(1))
//
//}

package main // 声明文件所在的包，每个go文件必须有归属的包
import "fmt" // 依赖的其他包

// 主函数，程序入口
func main() {
	fmt.Println("Hello World")
}
