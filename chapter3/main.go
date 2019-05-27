package main

//currency calc
import "fmt"

var rubls float64

func main() {
	fmt.Println("Сколько рублей вы хотите поменять, введите сумму?")
	fmt.Scanf("%f", &rubls)
	const dollarkurs float64 = 64
	summa := rubls / dollarkurs
	fmt.Fprintf("Сумма конвертации, %.2f", summa)
}
