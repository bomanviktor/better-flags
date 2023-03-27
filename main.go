package main

import (
	"betterflags/betterflags"
	"fmt"
)

func main() {
	Test := betterflags.Create("test", "usage msg", true, false)
	Test2 := betterflags.Create("test2", "help ffs!", "Test2", false)
	betterflags.Parse(false)
	fmt.Println(*Test, *Test2)
}
