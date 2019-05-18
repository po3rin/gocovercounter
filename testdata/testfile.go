package testfile

import (
	"bufio"
	"fmt"
	"os"
)

func Testfile() {
	fmt.Println("text: ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()

	if text != "hello" {
		return
	}
	fmt.Println("hello")
}
