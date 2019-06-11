package main

import (
	"github.com/phillipahereza/testing_with_go/dependency_injection/greeting"
	"os"
)


func main() {
	greeting.Greet(os.Stdout, "ELodie")
}