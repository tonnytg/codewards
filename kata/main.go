package main

import (
	"fmt"
	"strings"
)

var (
	count int = 0
	check bool = true
)


func correctsolution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func mysolution(a, b string) bool {
	//fmt.Printf("Checando A: %v Qt: %v %T \n", a, len(a))
	//fmt.Printf("Checando B: %v Qt: %v %T \n", b, len(b))

	for i := 0; i < len(a); i++ {
		fmt.Println("A", string(a[i]))
		for j := 0; j < len(b); j++ {
			fmt.Println("B",string(b[j]))
			if string(a[i]) == "" && string(b[j]) == "" {
				return true
			} else if a[i] == b[j] {
				fmt.Println("A e B são iguais", string(a[i]), string(b[j]))
				i++
				count++;
			} else {
				count = 0;
			}

			if len(b) == count {
				return true
			}
		}
	}
	return false
}

func main()  {
	fmt.Println(correctsolution("", ""))
	fmt.Println(mysolution("", ""))
}