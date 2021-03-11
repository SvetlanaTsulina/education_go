package main

import "fmt"

func main() {
	fmt.Print("Введите температуру по Фаренгейту: ")
	var f int
	fmt.Scanf("%i", &f)
	c := (f - 32) * 5 / 9

	fmt.Println("Температура по Цельсию ", c)

}
