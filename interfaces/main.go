package main

import "fmt"

type triangle struct {
	base float64
	hight float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main()  {

	 s := square{sideLength: 4}
	 t := triangle{base: 5, hight: 4}
	 printArea(s)
	 printArea(t)

}


func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.hight
}

func (s square) getArea()  float64{
	return s.sideLength * s.sideLength
}

func printArea(s shape)  {
	fmt.Println(s.getArea())
}
