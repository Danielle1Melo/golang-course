package main

import ("fmt")

type dani int // criei um tipo: dani

var x dani

func main() {
	x = 42
	fmt.Printf("%T\n", x) //%T para saber o tipo da variável
	fmt.Println(x)

}
