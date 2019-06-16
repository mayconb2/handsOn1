// HANDS ON 1
// create a type square
// create a type circle
// attach a method to each that calculates area and returns it
// create a type shape which defines an interface as anything which has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle


package main

import "fmt"

const pi float64 = 3.14

type square struct {

		lado float64

}

type circle struct {

	raio float64

}

func (s square) area() float64{

	areaSquare := s.lado * s.lado

	return areaSquare
	

}

func (c circle) area() float64{

	areaCirculo := pi * c.raio * c.raio

	return areaCirculo
	

}

type shape interface {

	area () float64

}

func info (s shape) {
	fmt.Println (s.area())
}

func main () {

	s1 := square {
		 2,
	}


	c1:= circle {
		4.2,

	}

	s2 := square {3}
	c2 := circle {4}

	fmt.Println(c1.area())
	fmt.Println(s1.area())
	info(s2)
	info(c2)



}