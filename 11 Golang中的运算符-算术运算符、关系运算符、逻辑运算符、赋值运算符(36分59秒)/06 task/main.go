//位运算符
package main

import "fmt"

func main() {
	//练习1：有两个变量，a和b，要求将其进行交换，最终打印结果

	// var a = 34
	// var b = 10
	// t := a //t=34
	// a = b  //a=10
	// b = t  //b=34
	// fmt.Printf("a=%v b=%v", a, b)

	//练习2：有两个变量，a和b，要求将其进行交换(不能使用中间变量)，最终打印结果

	// var a = 34
	// var b = 10

	// a = a + b //a=34+10
	// b = a - b //b=34+10-10  34
	// a = a - b //a=(34+10)-(34+10-10)   10
	// fmt.Printf("a=%v b=%v", a, b)

	// 练习3：假如还有100天放假，问：xx个星期零xx天

	// var days = 100
	// var week = days / 7
	// var day = days % 7
	// fmt.Printf("距离放假还有%v周零%v天", week, day)

	/*
	 练习4：定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：
	 C=(F-32)÷1.8,摄氏温标（°C）和华氏温标（°F），请求出华氏温度对应的摄氏温度
	*/

	var F float32 = 100 //华氏温度

	C := (F - 32) / 1.8

	fmt.Printf("华氏温度对应的摄氏温度是%.2f", C)

}
