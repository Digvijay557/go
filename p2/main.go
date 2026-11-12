package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() 
	Perimeter() 
}
type Circle struct{
	radius float64
}

func (circle *Circle) Area() {
	fmt.Printf("Area: %f\n", circle.radius*circle.radius*math.Pi)
}

func (circle *Circle) Perimeter() {
	fmt.Printf("Perimeter: %f\n", 2*math.Pi*circle.radius)
}
func (circle *Circle) Scale(factor float64){
	circle.radius = circle.radius*factor
}

type Rectangle struct{
	length float64
	width float64
}

func (rectangle *Rectangle) Area(){
	fmt.Printf("Area: %f\n", rectangle.length*rectangle.width)
}

func (rectangle *Rectangle) Perimeter(){
	fmt.Printf("Perimeter: %f\n", 2*(rectangle.length+rectangle.width))
}
func (rectangle *Rectangle) Resize(nlength float64 , nwidth float64){
	rectangle.length = nlength
	rectangle.width = nwidth
}
func getType(shape Shape) {
	switch shape.(type) {
	case *Circle:
		fmt.Println("This is a Circle")
	case *Rectangle:
		fmt.Println("This is a Rectangle")
	default:
		fmt.Println("This is some other shape")
	}
}

func printShape(shape Shape){
	shape.Area()
	shape.Perimeter()
}

func main() {
	circle := Circle{4.0}
	// circle.Area()
	// circle.Perimeter()
	// circle.Scale(3)
	// circle.Area()
	rectangle := Rectangle{3,4}
	rectangle2 := Rectangle{8,6}
	// rectangle.Area()
	// rectangle.Perimeter()
	// rectangle.Resize(9,10)
	// rectangle.Area()
	// printShape(&rectangle)
	// printShape(&rectangle2)
	// printShape(&circle)


	shapes := []Shape{
		&circle,
		&rectangle,
		&rectangle2,
	}
	for i := 0; i < len(shapes); i++ {
		printShape(shapes[i])
		getType(shapes[i])
	}

}
