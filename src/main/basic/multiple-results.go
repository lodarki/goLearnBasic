package basic

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	s, i := swap("hello", "world")
	fmt.Print(s, " ", i)
}
