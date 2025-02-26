package main

import "fmt"

type Diy struct {
	A int64
	b float64
}

// 自定义String方法，显示未导出字段
func (d Diy) String() string {
	return fmt.Sprintf("{A: %d, b: %f}", d.A, d.b)
}

func (diy *Diy) Set(a int64, b float64) {
	diy.A = a
	diy.b = b
}

func (diy Diy) Set2(a int64, b float64) {
	diy.A = a
	diy.b = b
}

func main() {
	// 正确初始化未导出字段（同一包内允许）
	g := Diy{
		A: 2,
		b: 4.0,
	}
	fmt.Println("Initial g:", g) // 使用自定义String方法输出

	g.Set(1, 1)
	fmt.Printf("After Set(1,1): %v\n", g)

	g.Set2(3, 3)
	fmt.Printf("After Set2(3,3): %v\n", g) // 值未改变

	k := &Diy{
		A: 2,
		b: 5.0,
	}
	fmt.Println("Initial k:", k)

	k.Set(1, 1)
	fmt.Printf("After Set(1,1): %v\n", k)

	k.Set2(3, 3)
	fmt.Printf("After Set2(3,3): %v\n", k) // 值未改变
}
