package main

import "fmt"

var (
	count int = 0
	check bool = true
)

func solution(a, b string) bool {
	fmt.Println("Checando A:", a, len(a))
	fmt.Println("Checando B:", b, len(b))

	for i := 0; i < len(a); i++ {
		fmt.Println("A", string(a[i]))
		for j := 0; j < len(b); j++ {
			fmt.Println("B",string(b[j]))
			if a[i] == b[j] {
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
	fmt.Println(solution("1oo", "100"))
}