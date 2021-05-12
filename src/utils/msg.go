package utils

import "fmt"

// Log ...
// log msg
func Log(v string) {
	fmt.Printf("\033[36mLog > \033[0m%s\n", v)
}

// Wan ...
// warn msg
func Wan(v string) {
	fmt.Printf("\033[33mWarning > \033[0m%s\n", v)
}

// Err ...
// output error msg
func Err(v string, e error) {
	fmt.Printf("\033[31mError > \033[0m%s \033[31m: at\033[0m %s\n", v, e)
}
