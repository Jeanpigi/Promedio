package main

import "fmt"

func main()  {
	var n, sum, num, promedio int

	fmt.Println("Escriba un número: ")
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Println("Escriba el número que participa en el promedio")
		fmt.Scanf("%d", &num)
		sum = sum + num
	}

	promedio = sum / n

	fmt.Println("El promedio es ", promedio)
}