package hello

import "fmt"

// Hello - greet
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
