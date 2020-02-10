package main

import (
	"bufio"
	"fmt"
	"os"
)

//var (
//	firstName, lastName, s string
//	i                      int
//	f                      float32
//	input                  = "56.12 / 5212 / Go"
//	format                 = "%f / %d / %s"
//)
//
//func main() {
//	fmt.Println("Please enter your full name: ")
//	fmt.Scanln(&firstName, &lastName)
//	//fmt.Scanf("%s %s", &firstName, &lastName)
//	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
//	fmt.Sscanf(input, format, &f, &i, &s)
//	fmt.Println("From the string we read: ", f, i, s)
//	// 输出结果: From the string we read: 56.12 5212 Go
//}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	//switch input {
	//case "lisai\n":
	//	fmt.Printf("welcome %s", input)
	//case "anfeng\n":
	//	fmt.Printf("welcome %s", input)
	//case "llussy\n":
	//	fmt.Printf("welcome %s", input)
	//default:
	//	fmt.Println("You are not welcome here! Goodbye!")
	//}
	switch input {
	case "lisai\n":
		fallthrough
	case "anfeng\n":
		fallthrough
	case "llussy\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}

// fallthrough：Go里面switch默认相当于每个case最后带有break，
// 匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。
