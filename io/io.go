package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// функция ReadArray считывает размер и элементы массива из ввода
func ReadArray(sizePrompt, elementsPrompt string) []string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(sizePrompt)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())

	array := make([]string, size)
	fmt.Println(elementsPrompt)
	for i := 0; i < size; i++ {
		scanner.Scan()
		array[i] = scanner.Text()
	}

	return array
}

// функция PrintCommonElements выводит общие элементы массива
func PrintCommonElements(commonElements []string) {
	fmt.Println("Результат:")
	if len(commonElements) == 0 {
		fmt.Println("Общих элементов не найдено")
	} else {
		for _, element := range commonElements {
			fmt.Println(element)
		}
	}
}