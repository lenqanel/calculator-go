package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("Деление на ноль невозможно")
		}
		return num1 / num2
	default:
		panic("Неверная операция")
	}
}

func main() {
	var input string
	fmt.Print("Введите выражение: ")
	scanner :=
		bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	c := strings.Fields(input)
	if len(c) != 3 {
		panic("Неверный формат ввода") // ожидается формат в 3 строки
	}

	x, operator, y := c[0], c[1], c[2]

	num1, err1 := strconv.Atoi(x)
	num2, err2 := strconv.Atoi(y)

	roman1 := isRoman(x)
	roman2 := isRoman(y)

	//первая проверка
	if (err1 != nil && !roman1) || (err2 != nil && !roman2) {
		panic("Числа могут быть только римскими или арабскими и должны быть в допустимых пределах")
	}
	//вторая проверка
	if roman1 != roman2 {
		panic("Недопустимая комбинация ввода")
	}
	//допуск к конвертации числа, условие меньше 10
	if roman1 {
		num1 = RomanToInt(x)
		num2 = RomanToInt(y)
	}
	if num1 > 10 || num2 > 10 {
		panic("Недопустимое арабское число")

	}
	result := calculate(num1, num2, operator)
	//значения равны и преобразованы, результат, условие для римских чисел
	if roman1 {
		if result < 1 {
			panic("Результат не может быть меньше единицы")
		}
		fmt.Print("Результат: ", IntToRoman(result))
	} else {
		fmt.Println("Результат: ", result)
	}

}

// функция на соответсвтие римских чисел
func isRoman(num string) bool {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, romanNumeral := range romanNumerals {
		if num == romanNumeral {
			return true
		}
	}
	return false

}

// конвертация
func RomanToInt(s string) int {
	var v, lv, cv int
	h := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}
	//перевод римских чисел в арабские
	for k := len(s) - 1; k >= 0; k-- {
		cv = h[s[k]] //текущее значение
		if cv < lv {
			v -= cv
		} else {
			v += cv
		}
		lv = cv
	}

	return v
}

// из арабских в римские, использование строки
func IntToRoman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value { //число больше преобразования
			roman += conversion.digit  //добавление символов
			number -= conversion.value //вычитание значения
		}
	}
	return roman
}
