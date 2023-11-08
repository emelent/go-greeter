package greeter

import "fmt"

// Hello returns a string with a message saying hello
// to the given target.
func Hello(target string) string {
	return fmt.Sprintf("Hello %s.", target)
}
