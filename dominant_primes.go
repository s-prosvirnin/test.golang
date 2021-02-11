package main

import (
	"fmt"
	//"github.com/onsi/ginkgo"
)

func primeNumbersGenerator(endNumber int) []int {
	//2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199

	primeNumbersAsKeysArray := make([]int, endNumber+1)

	divisor := 2
	primeNumbersCount := 0
	for keyAsNumber, flag := range primeNumbersAsKeysArray {
		if keyAsNumber < 2 {
			primeNumbersAsKeysArray[keyAsNumber] = -1
			continue
		}
		if flag == -1 {
			continue
		} else {
			divisor = keyAsNumber
		}
		for number := divisor + 1; number <= endNumber; number++ {
			if number%divisor == 0 {
				primeNumbersAsKeysArray[number] = -1
			}
		}
		if primeNumbersAsKeysArray[keyAsNumber] == 0 {
			primeNumbersCount++
		}
	}

	primeNumbersArray := make([]int, primeNumbersCount)
	primeNumbersArrayKey := -1
	for keyAsNumber, flag := range primeNumbersAsKeysArray {
		if flag == 0 {
			primeNumbersArrayKey++
			primeNumbersArray[primeNumbersArrayKey] = keyAsNumber
		}
	}

	return primeNumbersArray
}

func dominantPrimes(start, end int) {

}

func doTestPrimeNumbersGenerator(start int, end int) {
	for _, primeNumber := range primeNumbersGenerator(end) {
		fmt.Printf("%d, ", primeNumber)
	}
}

func main() {
	doTestPrimeNumbersGenerator(0, 200)

	//doTest(0,10, 8)
	//doTest(0,50, 8)
	//doTest(0,200, 8)
	//doTest(2,200, 1080)
	//doTest(1000,100000,52114889)
	//doTest(4000,500000,972664400)
}
