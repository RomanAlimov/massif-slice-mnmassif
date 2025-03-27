package main

import (
	"fmt"
	"sort"
)

//многомерные массивы and обычные массивы and слайсы :

func massif() {
	// создал массив
	var num [3]int
	fmt.Println(num) // [0 0 0]
}

func masiff2() {
	var num [3]int
	num[0] = 3
	num[2] = 4
	num[1] = 6
	fmt.Print(num) // [3 6 4]
}

func masiff3() {
	num := [3]string{"sss", "ggg", "qqq"}
	fmt.Println(num) // [sss ggg qqq]

	for i := 0; i < len(num); i++ { // перебор массива
		fmt.Println(num[i]) // выводит элементы.
	}
}

func masiff4() {
	num := [4]float64{6, 123, 55, 66}
	var summa float64 = 0

	for i := 0; i < len(num); i++ {
		summa += num[i]
	}
	fmt.Println(summa) // складывает все элементы массива = 250.
}

//перейду к слайсам ведь это своего рода динамический массив

func slice() {
	slice := []int{12, 14, 15, 16, 17}
	slice = append(slice, 0) // append добавляет в слайс элемент в конец
	fmt.Println(slice)       // 12 14 15 16 17 0

	// так же можно изменять слайс

	slice1 := []int{11, 1231, 1414}
	slice1[0] = 777
	fmt.Println(slice1) // 777 1231 1414

	// слайс можно отсортировать

	slice3 := []int{10, 5, 3, 7, 9, 0, 4, 1, 2}
	sort.Ints(slice3)
	fmt.Println(slice3) // [0 1 2 3 4 5 7 9 10]

	// так же сортируется со строками (string)

	slice4 := []string{"b", "j", "y", "q"}
	sort.Strings(slice4)
	fmt.Println(slice4) // [b j q y]

}

func main() {
	num := [3][4]int{{3, 4, 5, 6}, {6, 7, 8, 9}, {10, 11, 12, 13}} //массив
	fmt.Println(num[2][1])                                         // 11 печатаем 2-ой элемент массива, и 1 ячейку
	massif()                                                       // функция вызвана.
	masiff2()
	masiff3()
	masiff4()
	slice()
}
