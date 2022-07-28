package main

import (
	"fmt"
)

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”

func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}


	// Ваша реалізація
	ch := make(chan int, 3)
	var result int

	for _, v := range n {
		go sum(v, ch)
		fmt.Printf("result of slice %v is - %v\n", v, sumSlice(v))
	}
	
	for i := 0; i < len(n); i++ {
		result += <- ch
	}

	fmt.Printf("result: %v\n", result)

}

func sum(input[]int, ch chan int){
	var result int
	for _, v := range input {
		result += v
	}
	ch <- result
}

func sumSlice(input[]int)int {
	var result int
	for _, v := range input {
		result += v
	}
	return result
}