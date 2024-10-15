package main

import (
	"fmt"
	"strconv"
)

func main() {
	rim := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
		"C":    100,
	}

	//Ввод данных

	var num1, num2, op, err1, err2 string
	fmt.Scanln(&num1, &op, &num2, &err1, &err2)
	var Num1, Num2 int
	var Rim, Arab bool = true, true

	//Перевод строки в числа и проверка системы на римскую или арабскую

	if value, inMap := rim[num1]; inMap {
		Num1 = value
		Arab = false

	} else {
		Num1, _ = strconv.Atoi(num1)
		Rim = false
	}
	if value, inMap := rim[num2]; inMap {
		Num2 = value
		Arab = false

	} else {
		Num2, _ = strconv.Atoi(num2)
		Rim = false
	}
	//Обработка ошибок
	if Rim == false && Arab == false {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	if Num1 > 10 || Num2 > 10 {
		panic("Выдача паники, так как калькулятор на вход принимает числа не больше десяти")
	}
	if Num1 <= 0 || Num2 <= 0 {
		panic("Выдача паники, так как калькулятор на вход принимает числа не меньше единицы")
	}
	if err1 != "" || err2 != "" {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда")
	}

	//Работа калькулятора

	res := 0

	switch {
	case op == "+" && (Rim == true || Arab == true):
		res = (Num1 + Num2)
	case op == "-" && (Rim == true || Arab == true):
		res = (Num1 - Num2)
	case op == "*" && (Rim == true || Arab == true):
		res = (Num1 * Num2)
	case op == "/" && (Rim == true || Arab == true):
		res = (Num1 / Num2)
	default:
		panic("Выдача паники, так как строка не является математической операцией.")
	}

	//Вывод результата
	if Arab == true {
		fmt.Print(res)
	}
	//Перевод результата в римскую систему
	if Rim == true {
		if res <= 0 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел или нулей.")
		}
		rimForRes := map[int]string{
			1:   "I",
			4:   "IV",
			5:   "V",
			9:   "IX",
			10:  "X",
			40:  "XL",
			50:  "L",
			90:  "XC",
			100: "C",
		}
		var rimRes string
		values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
		for _, value := range values {
			// Пока res >= value, вычитаем его и добавляем римский символ
			for res >= value {
				rimRes += rimForRes[value]
				res -= value
			}
		}
		fmt.Print(rimRes)
	}

}
