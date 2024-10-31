package main

import "fmt"

type Point struct {
	X, Y float64
}

func main() {
	/*
		调用函数时会对每一个参数值进行拷贝，如果一个函数需要更新一个变量或参数太大就需要指针
		声明方法

		Nil可以作为一个合法的接收器
	*/
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p := Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
} //方法名为(*Point).ScaleBy 调用时只需要提供一个Point类型的指针即可
