package main

import "fmt"

func printType(a interface{}, description string) {
	//fmt.Print("(" + reflect.TypeOf(a).String() + ")")
	fmt.Printf("%v: [%T] %+v\n", description, a, a)
}

func Snail(snaipMap [][]int) []int {

	//var resultArray = (snaipMap[0])[:]
	linesCount := len(snaipMap)
	var resultArray = make([]int, linesCount*linesCount)
	printType(snaipMap, "snaipMap")
	for lineKey, lineArray := range snaipMap {
		for cellKey, cell := range lineArray {
			if lineKey == 0 || lineKey == linesCount {
				resultArray[lineKey*cellKey] = cell
			}
		}
	}

	return resultArray
}

func snailTest() {
	snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	printType(Snail(snailMap), "snail")
}

func main() {
	snailTest()
}
