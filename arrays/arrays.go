package arrays

// функция FindCommonElements возвращает слайс общих элементов двух массивов
func FindCommonElements(arr1, arr2 []string) []string {
	common := make([]string, 0)
	mapping := make(map[string]bool)

	for _, item := range arr1 {
		mapping[item] = true
	}

	for _, item := range arr2 {
		if _, exists := mapping[item]; exists {
			common = append(common, item)
		}
	}

	return common
}