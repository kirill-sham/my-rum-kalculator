package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		num1, num2, oper string //вводимые данные
	)
	for {
		fmt.Println("Введите данные (операнд, оператор (+, -, /, *), операнд):")

		_, err := fmt.Scanf("%s %s %s", &num1, &oper, &num2)
		if err != nil {
			panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *). ")
		}

		numint1, err1 := strconv.Atoi(num1)
		numint2, err2 := strconv.Atoi(num2)

		//если оба числа инт переходит в элс и считает
		if err1 != nil || err2 != nil {
			if !(err1 != nil && err2 != nil) {
				panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
			}
			//обработка римских цифр

			switch oper {
			case "+":
				fmt.Println(intToRoman(add(romanToInt(num1), romanToInt(num2))))
			case "-":
				fmt.Println(intToRoman(subtr(romanToInt(num1), romanToInt(num2))))
			case "*":
				fmt.Println(intToRoman(multip(romanToInt(num1), romanToInt(num2))))
			case "/":
				fmt.Println(intToRoman(divide(romanToInt(num1), romanToInt(num2))))
			default:
				panic("Несоответствующая операция.")

			}

		} else {
			switch oper {
			case "+":
				fmt.Println(add(numint1, numint2))
			case "-":
				fmt.Println(subtr(numint1, numint2))
			case "*":
				fmt.Println(multip(numint1, numint2))
			case "/":
				fmt.Println(divide(numint1, numint2))
			default:
				panic("Несоответствующая операция.")

			}
		}

	//	fmt.Printf("Value: %v %v %v\n", num1, oper, num2)

	}
}

func add(a int, b int) int {
	return a + b
}

func subtr(a int, b int) int {
	return a - b
}

func multip(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

// римские цифры переводит в десятичные
func romanToInt(s string) int {
	var (
		result int
		cur    = 1001
	)
	for i := 0; i < len(s); i++ {
		prev := cur
		switch s[i] {
		case 'I':
			cur = 1
		case 'V':
			cur = 5
		case 'X':
			cur = 10
		case 'L':
			cur = 50
		case 'C':
			cur = 100
		case 'D':
			cur = 500
		case 'M':
			cur = 1000
		default:
			panic("Введина не римская цифра.")

		}
		if cur > prev {
			result += cur - 2*prev
		} else {
			result += cur
		}
	}
	return result
}

func intToRoman(num int) string {
	var (
		result string
	)
	if num <= 0 {
		panic("Римские цифры не могут равнятся нулю и быть отрицательными")
	}
	for num > 0 {
		switch {
		case num/1000 > 0:
			result += "M"
			num -= 1000
		case num/900 > 0:
			result += "CM"
			num -= 900
		case num/500 > 0:
			result += "D"
			num -= 500
		case num/400 > 0:
			result += "CD"
			num -= 400
		case num/100 > 0:
			result += "C"
			num -= 100
		case num/90 > 0:
			result += "XC"
			num -= 90
		case num/50 > 0:
			result += "L"
			num -= 50
		case num/40 > 0:
			result += "XL"
			num -= 40
		case num/10 > 0:
			result += "X"
			num -= 10
		case num/9 > 0:
			result += "IX"
			num -= 9
		case num/5 > 0:
			result += "V"
			num -= 5
		case num/4 > 0:
			result += "IV"
			num -= 4
		case num/1 > 0:
			result += "I"
			num -= 1
		}
	}
	return result
}
