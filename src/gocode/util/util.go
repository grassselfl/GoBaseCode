package util

import "fmt"

var utilVersion = "1.0"

func init() {
	fmt.Println("utilVersion: ", utilVersion)
}

func Int2str(num int) string {
	return fmt.Sprintf("%c", num)
}
