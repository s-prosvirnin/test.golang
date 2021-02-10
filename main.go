package main

import (
	"fmt"
	//"golang.org/x/tour/wc"
	"math"
	"reflect"
	"strings"
)

func main() {
	//wc.Test(WordCount)

	/*f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}*/

	//collections()

	//extend()

	//encapsulation()

	//cast()

	//stringable()

	//halo()

	//interfaces1()

	//pointers()

	//structPointers()

	errors()
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func pow(x, n, lim float64) float64 {
	fmt.Println(math.Pow(x, n))
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func arrays() {
	type A struct {
		a int
	}
	a := A{1}

	c := [2]A{a, a}
	c[0].a = 2
	fmt.Println(c)
}

func pointers() {
	a := 1
	//b := 2

	p1 := &a
	p2 := &a
	p3 := p2
	*p1 = 3
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(p1)
	fmt.Println(*p2)
	fmt.Println(*p3)
	fmt.Println(*p1)
	fmt.Printf("[%T] %+v\n", a, a)
	fmt.Printf("[%T] %+v\n", &a, &a)
	fmt.Printf("[%T] %+v\n", p1, p1)
	fmt.Printf("[%T] %+v\n", &p1, &p1)
	fmt.Printf("[%T] %+v\n", *p2, *p2)
}

func WordCount(s string) map[string]int {
	fmt.Println(strings.Fields(s))
	arStr := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range arStr {
		m[v] = m[v] + 1
	}
	return m
}

func fibonacci() func() int {
	previous := 0
	current := 0
	sequence := 0

	return func() int {
		if sequence == 1 {
			current = 1
		} else if sequence > 1 {
			savedPrevious := previous
			previous = current
			current = current + savedPrevious
		}
		sequence++

		return current
	}
}

type A1 struct {
	a int
	b int
}

func createA1(A int, B int) A1 {
	return A1{A, B}
}

func (a1 *A1) A() int {
	return a1.a
}

func (a1 *A1) B() int {
	return a1.b
}

func (a1 *A1) update(A int, B int) {
	a1.a = A
	a1.b = B
}

type A1Collection []*A1

func (a1Collection A1Collection) Update() {
	for index, _ := range a1Collection {
		a1Collection[index].update(100, 100)
	}
}

func collections() {
	a1 := A1{10, 10}
	a2 := createA1(20, 20)
	a2.update(21, 21)

	a1Collection := make(A1Collection, 2)
	a1Collection[0] = &a1
	a1Collection[1] = &a2
	a1Collection.Update()
	//a1Collection[0].update(10, 10)

	fmt.Println(a1)
	fmt.Println(a1Collection[0])
	fmt.Println(a1Collection[1])

	var a2Collection [2]*A1
	a2Collection[0] = &a1
	a2Collection[1] = &a2
	a2Collection[1].update(200, 200)

	fmt.Println(a2Collection[0])
	fmt.Println(a2Collection[1])
	fmt.Println(a2)
}

type parent struct {
	x int
	z int
}

func createParent() parent {
	return parent{10, 10}
}

func (parent *parent) updateX() {
	parent.x = 10
}

func (parent *parent) updateZ() {
	parent.z = 10
}

type child struct {
	parent
	y int
	p parent
}

func (child *child) updateX() {
	child.parent.updateX()
	//child.x = 20
}

func (child *child) updateY() {
	child.y = 20
}

func extend() {
	child := child{}
	fmt.Println(child)
	child.updateY()
	fmt.Println(child)
	child.updateX()
	fmt.Println(child)
	child.updateZ()
	fmt.Println(child)
	child.p.updateZ()
	fmt.Println(child)
}

type private struct {
	x int
	Y int
}

func (p *private) setX(x int) {
	p.x = x
}

func (p *private) setY(y int) {
	p.Y = y
}

func encapsulation() {
	p := private{}
	p.x = 2
	p.Y = 2

	p.setX(3)
	p.setY(3)

	fmt.Println(p)
}

func cast() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	//fmt.Println(f)
}

type Str struct {
	s string
}

func (str Str) String() string {
	return str.s
}

func printStringable(stringer fmt.Stringer) {
	fmt.Println(stringer)
}

func stringable() {
	//s := Str{"123"}
	//printStringable(s)
	i := IPAddr{127, 0, 0, 1}
	printStringable(i)
	fmt.Println(i)
	i2 := [4]int{1, 2, 3, 4}
	fmt.Println(i2)
}

type IPAddr [4]int

func (ipAddr IPAddr) String() string {
	for _, v := range ipAddr {
		return string(v)
	}
	return "000"
}

func halo() {
	fmt.Println("Halo!")
}

type animal interface {
	makeSound()
}

type cat struct{}
type dog struct{}

func (cat *cat) makeSound() {
	fmt.Println("meow")
}

func (dog *dog) makeSound() {
	fmt.Println("woof")
}

func animalSound(a animal) {
	a.makeSound()
	printType(a, "animalSound(a)")
	printType(&a, "animalSound(&a)")
}

func animalSoundP(a *animal) {
	//*a.makeSound()
	//*a.makeSound()
	printType(a, "animalSoundP(a)")
	printType(&a, "animalSoundP(&a)")
	printType(*a, "animalSoundP(*a)")
}

func structPointers() {
	var c, d animal = &cat{}, &dog{}
	c2 := &cat{}
	var c3 animal = &cat{}
	c1 := cat{}
	printType(c, "c")
	printType(&c, "&c")
	printType(c2, "c2")
	printType(*c2, "*c2")
	printType(d, "d")
	printType(c1, "c1")
	printType(&c1, "&c1")
	printType(c3, "c3")
	printType(&c3, "&c3")
	printType(animal(&c1), "animal(&c1)")
	animalSound(c)
	animalSoundP(&c)
	animalSound(c2)
	//animalSound(c3)
	//animalSoundP(&c3)
	//animalSoundP(&c2)
	//animalSound(&c1)
	//animalSoundP(&c1)
}

func interfaces1() {
	var c, d animal = &cat{}, &dog{}
	c1 := cat{}
	c.makeSound()
	d.makeSound()
	c1.makeSound()
	fmt.Println(reflect.TypeOf(&c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(&c1))

	animalSound(c)
	animalSound(d)
	animalSound(&c1)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(&c1)
}

func printType(a interface{}, description string) {
	//fmt.Print("(" + reflect.TypeOf(a).String() + ")")
	fmt.Printf("%v: [%T] %+v\n", description, a, a)
}

type MyError struct {
	code string
}

func (error *MyError) Error() string {
	return error.code
}

func giveError() (string, error) {
	return "result", &MyError{
		"crush!",
	}
}

func errors() {
	result, error := giveError()
	println(result)
	printType(error, "error")
}
