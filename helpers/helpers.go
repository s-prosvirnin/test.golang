package helpers

import "fmt"

func printType(a interface{}, description string) {
	//fmt.Print("(" + reflect.TypeOf(a).String() + ")")
	fmt.Printf("%v: [%T] %+v\n", description, a, a)
}
