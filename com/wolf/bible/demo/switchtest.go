package main

import "fmt"

func testSwitch1() {
	switch coinflip() {
	case "heads":
		fmt.Println("headers..") //不需要break
	case "tails":
		fmt.Println("tails..")
	case "a":
		fallthrough //执行下面case b
	case "b":
		fmt.Println("b..")
	default:
		fmt.Println("landed on edge!")
	}
}

func coinflip() string {
	return ""
}

//tagless switch
func testSwitch2() {
	var x int;
	switch {
	case x > 0:
		fmt.Println("x>0..")
	case x < 0:
		fmt.Println("x<0..")
	default:
		fmt.Println("x=0..")
	}
}
