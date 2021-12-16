package main

import "fmt"

func main() {
	toChange := "hello"
	changeValue2(toChange)
	fmt.Println(toChange)
}

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) string {
	str = "changed!"
	return str
}
