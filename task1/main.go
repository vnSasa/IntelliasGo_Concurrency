package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”

func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	var wg sync.WaitGroup
	
	for v := range n {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Printf("result of slice_%v is: %v\n", v, sum(n[v]))
		} (v)
	}
	wg.Wait()
}

func sum(input[]int)int {
	// Ваша реалізація
	var result int
	for _, v := range input {
		result += v
	}
	return result
}