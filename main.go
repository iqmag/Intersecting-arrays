package main

import (
	"Intersecting-arrays/arrays"
	"Intersecting-arrays/io"
)

func main() {
	array1 := io.ReadArray("Введите размер первого массива: ", "Введите элементы первого массива:")
	array2 := io.ReadArray("Введите размер второго массива: ", "Введите элементы второго массива:")

	commonElements := arrays.FindCommonElements(array1, array2)

	io.PrintCommonElements(commonElements)
}