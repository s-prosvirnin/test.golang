package main

import (
	"fmt"
	"math"
)

/**
The prime number sequence starts with: 2,3,5,7,11,13,17,19.... Notice that 2 is in position one.
3 occupies position two, which is a prime-numbered position. Similarly, 5, 11 and 17 also occupy prime-numbered positions. We shall call primes such as 3,5,11,17 dominant primes because they occupy prime-numbered positions in the prime number sequence. Let's call this listA.
As you can see from listA, for the prime range range(0,10), there are only two dominant primes (3 and 5) and the sum of these primes is: 3 + 5 = 8.
Similarly, as shown in listA, in the range (6,20), the dominant primes in this range are 11 and 17, with a sum of 28.
Given a range (a,b), what is the sum of dominant primes within that range? Note that a <= range <= b and b will not exceed 500000.
*/

//2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199

func primeNumbersGenerator(endNumber int) []int {
	primeNumbersAsKeysArray := make([]int, endNumber+1)

	divisor := 2
	primeNumbersCount := 0
	stopNumber := int(math.Sqrt(float64(endNumber)))
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
		if keyAsNumber <= stopNumber {
			for number := divisor + 1; number <= endNumber; number++ {
				if number%divisor == 0 {
					primeNumbersAsKeysArray[number] = -1
				}
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

func dominantPrimesRangeSum(start, end int) int {
	primeNumbers := primeNumbersGenerator(end)
	dominantPrimeNumbersSum := 0
	for _, number := range primeNumbers {
		if number > end {
			break
		}
		if len(primeNumbers)-1 >= number {
			if primeNumbers[number-1] >= start {
				dominantPrimeNumbersSum += primeNumbers[number-1]
			}
		}
	}

	return dominantPrimeNumbersSum
}

func showPrimeNumbers(end int) {
	for _, primeNumber := range primeNumbersGenerator(end) {
		fmt.Printf("%d, ", primeNumber)
	}
}

func doTestDominantPrimes(start int, end int, expectedSum int) {
	if dominantPrimesRangeSum := dominantPrimesRangeSum(start, end); dominantPrimesRangeSum != expectedSum {
		fmt.Printf("TEST NOT PASSED. %d-%d. expected: %d, result: %d\n", start, end, expectedSum, dominantPrimesRangeSum)
	} else {
		fmt.Printf("test passed. %d-%d. expected: %d, result: %d\n", start, end, expectedSum, dominantPrimesRangeSum)
	}
}

func main() {
	//showPrimeNumbers(500000)
	//return;

	/**
	TODO: не работает на рандомных
	 */
	doTestDominantPrimes(-20, 10, 8)
	doTestDominantPrimes(0, 10, 8)
	doTestDominantPrimes(6, 20, 28)
	doTestDominantPrimes(2, 200, 1080)
	doTestDominantPrimes(1000, 100000, 52114889)
	doTestDominantPrimes(4000, 500000, 972664400)
}
