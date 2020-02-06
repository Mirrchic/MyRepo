package convert

import "fmt"

//Base to Dec taking value and it base, and returning
//it in base:10
func BaseToDec(value string, base int) int {
	var res int
	var val int
	multrip := 1
	for i := len(value) - 1; i >= 0; i-- {
		fmt.Sscanf(string(value[i]), "%X", &val)
		res += multrip * val
		multrip *= base
	}
	return res
}

//Dec to Base converting dec value
//into base you need
func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var res string
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec = dec / base
	}
	return res
}

//Taking base and it value and retturning newbase value
func BaseToBase(number string, oldbase int, newbase int) string {
	var res string
	res = DecToBase(BaseToDec(number, oldbase), newbase)
	return res
}
