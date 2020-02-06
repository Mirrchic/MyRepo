package calculator

import (
	"calk/convert"
	"math"
)

//Simple calculatore function
func CalculatorInt(val string, a int, b int) int {
	var res int
	if val == "/" {
		res = a / b
	}
	if val == "*" {
		res = a * b
	}
	if val == "+" {
		res = a + b
	}
	if val == "-" {
		res = a - b
	}
	if val == "^" {
		res = int(math.Pow(float64(a), float64(b)))
	}
	return res
}

//function getting two numbers and their base, and algebratic method,
//and retturning result in base >= 16
func GetResult(num1 string, num2 string, base1 int, base2 int, baseRet int, char string) string {
	var result string
	value1 := convert.BaseToDec(num1, base1)
	value2 := convert.BaseToDec(num2, base2)
	valueReturn := CalculatorInt(char, value1, value2)
	if baseRet == 10 {
		result = string(valueReturn)
	} else {
		result = convert.DecToBase(valueReturn, baseRet)
	}
	return result

}
