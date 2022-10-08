package main

import "fmt"

type Sayer interface {
	say()
}
type Mover interface {
	move()
}
type Dog struct {
}
type Cat struct {
}

func (d Dog) say() {
	fmt.Println("汪汪汪")
}
func (c Cat) say() {
	fmt.Println("喵喵喵")
}

func (d Dog) move() {
	fmt.Println("狗会动")
}
func main() {
	var say Sayer
	a := Cat{}
	b := Dog{}
	say = a
	say.say()
	say = b
	say.say()

	var mover Mover
	fmt.Println(mover)
	var wangcai = Dog{}
	mover = wangcai
	wangcai.move()
	var fugui = &Dog{}
	mover = fugui
	fugui.move()
}
