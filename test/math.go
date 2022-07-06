package main

import (
	"encoding/json"
	"fmt"
	"github.com/baetyl/baetyl-go/v2/dmcontext"
	"strconv"
)

func main() {

}

func test1() {
	num := float32(-0.0000000000000000000000000000000000005414906)
	args := map[string]interface{}{"x1": num}

	res := num * 10

	bie := make(map[string]interface{})
	res2, _ := dmcontext.ExecExpression("x1*10", args, dmcontext.MappingCalculate)
	bie["x1"] = res2
	res2B, _ := json.Marshal(bie)

	println(num)
	println(res)
	println(res2)
	println(res2.(float64))
	println(string(res2B))
}

func test2() {
	num := float32(-0.0000000000000000000000000000000000005414906)
	num2 := float64(num)

	str := strconv.FormatFloat(float64(num), 'e', -1, 32)
	str2 := fmt.Sprint(num)
	num3, _ := strconv.ParseFloat(str, 64)
	num4, _ := strconv.ParseFloat(str2, 64)

	println(num)
	println(num2)
	println(str)
	println(str2)
	println(num3)
	println(num4)
}

func test3() {
	num := int(2)
	num2 := int16(2)
	num3 := int32(2)
	num4 := int64(2)

	res := float64(num)
	res2 := float64(num2)
	res3 := float64(num3)
	res4 := float64(num4)

	println(res)
	println(res2)
	println(res3)
	println(res4)
}

func test4() {
	num := float32(0.8213302)
	str := strconv.FormatFloat(float64(num), 'e', -1, 32)
	num2, _ := strconv.ParseFloat(str, 64)

	num3 := num2 + 1.12121
	num4 := num3 * 10

	num5 := 1.9425402
	num6 := float64(10)
	num7 := num5 * num6

	num8 := 1.12121
	num9 := (num2 + num8) * 10
	num10 := num2*10 + num8

	str9 := strconv.FormatFloat(num9, 'e', -1, 64)
	str10 := strconv.FormatFloat(num10, 'e', -1, 64)

	num11, _ := strconv.ParseFloat(str9, 64)
	num12, _ := strconv.ParseFloat(str10, 64)

	str11 := fmt.Sprint(num9)

	println(num4)
	println(num7)
	println(num8)
	println(num9)
	println(num10)
	println(num11)
	println(num12)
	println(str11)
}

func test5() {
	num := float32(+1.313811e-002)
	str := strconv.FormatFloat(float64(num), 'e', -1, 32)
	num2, _ := strconv.ParseFloat(str, 64)

	num3 := num2 * 10

	println(num3)
}
