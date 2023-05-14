package main

import (
	"fmt"
	"reflect"
	"strings"
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
	fmt.Println(2 + 2)
	fmt.Println(5 - 1)
	fmt.Println(2 * 2)
	fmt.Println(8 / 2)
	fmt.Println(10 % 2)

	fmt.Println(word, num - 4)

	// Strings

	var cadena string = "Hola mundo"
	fmt.Println(cadena[0:4])
	fmt.Println(cadena[:4])
	fmt.Println(cadena[5:])

	fmt.Println(strings.ReplaceAll(cadena, "o", "-"))
	fmt.Println(strings.ToUpper(cadena))

	// Inputs - Scan
	var inputName string
	var inputAge int

	fmt.Println("Introduce tu nombre:")
	fmt.Scanln(&inputName)
	fmt.Println("Introduce tu edad:")
	fmt.Scanln(&inputAge)

	fmt.Println("Nombre:", inputName, "\nEdad:", inputAge)

	// Verbos en go, %s = string, %d = int, %v = var
	fmt.Printf("Nombre: %s \nEdad: %d \n", inputName, inputAge)
	fmt.Printf("Nombre: %v \nEdad: %v \n", inputName, inputAge)

	// Booleanos

	var booleano bool = true;
	fmt.Println(booleano) // true
	fmt.Println(4 > 1) // true

	// Condicionales

	// Podemos declarar variables antes de evaluar la condición, pero solo viven en el scope del if
	// No hacen falta paréntesis
	if data := 40; data == 40 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	
	switch data:= 40; data {
	case 1:
		fmt.Println(data)
	case 40:
		fmt.Println(data)
	}

	// Bucles
	// No existen los while en go
	for i := 0; i<10; i++ {
		// fmt.Println(i)
	}

	// Keywords continue, break, defer

	// Continue
	for i := 0; i <5; i++ {
		if i == 2 {
			fmt.Println("Condicional true")
			continue // Si no lo añadimos, printea el 2 y el mensaje. De esta forma, vuelve a evaluar la condición del bucle
		}
		fmt.Println(i)
	}
	fmt.Println("Fin del bucle 1")

	// Break
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("Condicional true")
			break // Finaliza la iteración
		}
		fmt.Println(i)
	}
	fmt.Println("Fin del bucle 2")

	// Defer

	defer fmt.Println("Segundo") // Al añadir esta kw, el código se ejecuta en último lugar
	fmt.Println("Primero")
}