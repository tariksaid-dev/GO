package main

import (
	"fmt"
	"reflect"
)

// Devuelve 0 si no se le asigna valor
var ext int

func main() {
	fmt.Println("Hi mom")

	var num int = 10
	const word string = "palabra"

	fmt.Println(num, word, ext)
	fmt.Println(reflect.TypeOf(num))

	// En go se le puede asignar bit-size a los números
	var numero8bits int8 = 4
	fmt.Println(numero8bits, reflect.TypeOf(numero8bits))

	numeroEntero := 60
	fmt.Println(numeroEntero, reflect.TypeOf(numeroEntero))
	
	// Para los flotantes solo existen 32 y 64 bits
	numeroFlotante := 1.1
	var numeroFlotante32 float32 = 3.2
	fmt.Println(numeroFlotante, reflect.TypeOf(numeroFlotante))
	fmt.Println(numeroFlotante32, reflect.TypeOf(numeroFlotante32))

	// Operadores
}