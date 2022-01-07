/*针对大整数的运算：主要实现两个数组对应元素的相加操作*/
package main

import (
	"fmt"

	"math/big"
	"math/rand"
	"time"
)

func main() {

	// p(x) = x^3 + x + 5
	//定义p1为一个大整数类型的数组
	p1 := []*big.Int{

		big.NewInt(1), // x^1
		big.NewInt(2), // x^2
		big.NewInt(3), // x^3
		//big.NewInt(1), // x^4
	}

	fmt.Println("p = ", p1)

	//定义p2是一个大整数类型的数组，含有3个元素
	//并给该数组随机赋值
	var p2 [3]*big.Int
	rand.Seed(time.Now().UnixNano()) // 取纳秒时间戳，可以保证每次的随机数种子都不同
	for i := 0; i < 3; i++ {
		a := rand.Intn(10) // Intn(n)返回一个取值范围在[0,n)的伪随机int值
		b := int64(a)
		p2[i] = big.NewInt(b)

	}

	fmt.Println("p2 = ", p2)

	//定义p3是一个大整数类型的数组，用来存放两个数组对应元素相加的结果
	//注意：res1 := new(big.Int) 的位置，应该在循环内
	var p3 [3]*big.Int
	for i := 0; i < 3; i++ {
		res1 := new(big.Int) //用来计算
		a := res1.Add(p1[i], p2[i])
		p3[i] = a
		//p3 = append(p3, a)
		//fmt.Println("p3 = ", p3)

	}
	fmt.Println("p3 = ", p3)

}
