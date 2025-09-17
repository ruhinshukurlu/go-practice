package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	fmt.Println(transformNumbers(&numbers, double))
	fmt.Println(transformNumbers(&numbers, triple))
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dnumbers := []int{}

	for _, val := range *numbers {
		dnumbers = append(dnumbers, transform(val))
	}

	return dnumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
