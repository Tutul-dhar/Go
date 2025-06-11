package main

import "fmt"

func main() {

 	fmt.Println(len("tutul dhar"))
	fmt.Println("hello world"[1]) // print ascii value of 'e'
	fmt.Println("%c", "hello world"[1])
	/* 
		this lione does not work correctly. 
		go sees this line like fmt.Println(arg1, arg2)
		so it print 
		%c 101
	*/
	fmt.Printf("%c\n", "hello world"[1]) // this line work correctly 
	fmt.Println("hello " + "world")
}