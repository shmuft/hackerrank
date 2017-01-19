package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	arr := make(map[string]bool)
	for _, i := range s {
		s2 := strings.ToLower(string(i))
		if s2 != " " {
			_, ok := arr[s2]
			if !ok {
				arr[s2] = true
			}
		}
	}
	if len(arr) == 26{
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
