package main

import (
	"os"
)

//func main() {
//	outputfile, err := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
//	if err != nil {
//		fmt.Printf("An error occurred with file opening or creation\n")
//		return
//	}
//	defer outputfile.Close()
//
//	outputWriter := bufio.NewWriter(outputfile)
//	outputString := "hello world!\n"
//	for i := 0; i < 10; i++ {
//		outputWriter.WriteString(outputString)
//	}
//	outputWriter.Flush()
//}

func main() {
	os.Stdout.WriteString("hello, world\n")
	f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	f.WriteString("hello, world in a file\n")
}
