package greeting

import (
	"io"
	"fmt"
)

// Greet does greeting things
func Greet(writer io.Writer, name string)  {
	fmt.Fprintf(writer, "Hello, %s\n", name)
	
}