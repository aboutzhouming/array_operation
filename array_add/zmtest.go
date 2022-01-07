package main

import (
	"fmt"

	"math/big"
	"math/rand"
	"time"
)

func main() {

	// p(x) = x^3 + x + 5
	P := []*big.Int{

		big.NewInt(1), // x^1
		big.NewInt(2), // x^2
		big.NewInt(3), // x^3
		//big.NewInt(1), // x^4
	}
	//assert.Equal(t, "x³ + x¹ + 5", PolynomialToString(p))
	fmt.Println("p = ", P)
	//assert.Equal(t, "x³ + x¹ + 5", PolynomialToString(p))
	//fmt.Println("11111")

	var P2 [3]*big.Int
	rand.Seed(time.Now().UnixNano()) // 取纳秒时间戳，可以保证每次的随机数种子都不同
	for i := 0; i < 3; i++ {
		a := rand.Intn(10) // Intn(n)返回一个取值范围在[0,n)的伪随机int值
		b := int64(a)
		P2[i] = big.NewInt(b)

	}

	fmt.Println("p2 = ", P2)

	//res1 := big.NewInt(1)
	var p3 [3]*big.Int
	for i := 0; i < 3; i++ {
		res1 := new(big.Int) //用来计算
		a := res1.Add(P[i], P2[i])
		p3[i] = a
		//p3 = append(p3, a)
		//fmt.Println("p3 = ", p3)

	}
	fmt.Println("p3 = ", p3)

	//fmt.Println("P[1]=", P[1])
	//fmt.Println("P2[1]=", P2[1])
	//
	//res := new(big.Int) //用来计算
	//var c *big.Int      //存储计算结果
	//
	//c = res.Add(P[0], P2[0])
	//
	////c.Add(P[0], P[1])
	//fmt.Println("p4=", c)

}
