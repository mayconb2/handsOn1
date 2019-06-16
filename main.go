package main

import "fmt"

const pi float64 = 3.14

type square struct {

		lado float64

}

type circle struct {

	raio float64

}

func (s square) areaS() float64{

	areaSquare := s.lado * s.lado

	return areaSquare
	

}

func (c circle) areaC() float64{

	areaCirculo := pi * c.raio * c.raio

	return areaCirculo
	

}


func main () {

	s1 := square {
		 2,
	}


	c1:= circle {
		4.2,

	}

	fmt.Println(c1.areaC())
	fmt.Println(s1.areaS())

}