package main

import (
	"fmt"
	"math"
)

func main() {
	//Первая задача
	fmt.Println("№1")
	var Cost float64
	var rent float64 = 95000
	var perc float64 = 10
	Cost = rent + (rent / 100 * perc)
	fmt.Println("Ответ на первую задачу = ", Cost)
	fmt.Println()
	//Вторая задача
	fmt.Println("№2")
	var laptop float32 = 55480
	var monitor float32 = 21830
	var mouse float32 = 890
	var keyboard float32 = 1560

	var check float32 = (6 * laptop) + (3 * monitor) + (11 * mouse) + (5 * keyboard)
	fmt.Println("Ответ на вторую задачу = ", check)
	fmt.Println()
	//Третья задача
	fmt.Println("№3")
	var storage float64 = 5000
	var file float64 = 256
	var x1 float64 = math.Floor(storage / file)
	var filesSize float64 = file * x1
	var x2 float64 = storage - filesSize
	fmt.Println("Можно разместить ", x1, " файлов и останется ", x2, " гб свободного места")
	fmt.Println()
	//Четвертая задача
	fmt.Println("№4")
	fmt.Println("Введите число")
	var f float32
	_, err := fmt.Scan(&f)

	if err != nil {
		fmt.Println("Введите ЧИСЛО")
		return
	}

	var celc float32 = 5.0 / 9.0 * (f - 32)

	fmt.Println(celc, " градусов")
	fmt.Println()
	//Пятая задача
	fmt.Println("№5")
	fmt.Println("Введите радиус")
	var r float32
	_, err1 := fmt.Scan(&r)
	if err1 != nil {
		fmt.Println("Введите РАДИУС")
		return
	}
	var l float32 = 2 * math.Pi * r
	var s float32 = math.Pi * r * r
	fmt.Println("Длина = ", l, ",площадь равна = ", s)
	fmt.Println()
	//Шестая задача
	fmt.Println("№6")
	fmt.Println("Введите кол-во лет")
	var y float64
	_, err2 := fmt.Scan(&y)
	if err2 != nil {
		fmt.Println("Введите корректное кол-во лет")
		return
	}

	fmt.Println("Введите годовую ставку")
	var a float64
	_, err20 := fmt.Scan(&a)
	if err20 != nil {
		fmt.Println("Введите корректную годовую ставку")
		return
	}

	fmt.Println("Введите начальную сумму")
	var i float64
	_, err21 := fmt.Scan(&i)
	if err21 != nil {
		fmt.Println("Введите корректную начальную сумму")
		return
	}

	var calculator float64 = i * (math.Pow((1 + a/100), y))
	fmt.Println(calculator, " единиц валюты")
	fmt.Println()
}
