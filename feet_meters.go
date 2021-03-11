package main

import "fmt"

func main() {
fmt.Print("Введите футы: ")
var foot float64
fmt.Scanf("%f", &foot)

meter := foot*0.3048

fmt.Println("Получается", meter, "метров.")



}

