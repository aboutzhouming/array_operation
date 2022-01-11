func Init_array_change(row, col, row_1, col_1 int) ([][]*big.Int, []*big.Int, [][]*big.Int) {

	//二维数组初始化为0或nil
	Init_2d_arr := make([][]*big.Int, row)
	for i := range Init_2d_arr {
		Init_2d_arr[i] = make([]*big.Int, col)
	}
	fmt.Println(Init_2d_arr)

	//二维数组随机赋值为big.Int类的值
	for i := 0; i < len(Init_2d_arr); i++ {
		for j := 0; j < len(Init_2d_arr[i]); j++ {
			a := rand.Intn(10) // Intn(n)返回一个取值范围在[0,n)的伪随机int值
			b := int64(a)
			Init_2d_arr[i][j] = big.NewInt(b)
			//p5[i][j] = new(big.Int).SetInt64(b)
			//p5 = p4
		}

	}
	fmt.Println("Init_2d_array = ", Init_2d_arr)

	//二维数据转一维数组
	To_1d_arr := make([]*big.Int, row*col)
	var cnt = 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			To_1d_arr[cnt] = Init_2d_arr[i][j]
			cnt++
		}

		//p4i[i][j] = p4
	}
	fmt.Println("1d_array =", To_1d_arr)

	//一维数组变想要的二维数组
	//const row_1 = 20
	//const col_1 = 5
	//var p5 [row_1][col_1]*big.Int
	wan_2d_arr := make([][]*big.Int, row_1)
	for i := range wan_2d_arr {
		wan_2d_arr[i] = make([]*big.Int, col_1)
	}
	fmt.Println("wan_2d_arr=", wan_2d_arr)

	for i := 0; i < row_1; i++ {
		for j := 0; j < col_1; j++ {
			wan_2d_arr[i][j] = To_1d_arr[i*col_1+j]
		}
	}
	fmt.Println("wan_2d_arr=", wan_2d_arr)

	return Init_2d_arr, To_1d_arr, wan_2d_arr

}
