package main

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	length, width float64
}
type Circle struct {
	radius float64
	pia    float64
}

func (ci *Circle) Area() float64 {
	if ci.radius <= 0 {
		return 0
	}
	return ci.pia * ci.radius * ci.radius
}
func (ci *Circle) Perimeter() float64 {
	if ci.radius <= 0 {
		return 0
	}
	return 2 * ci.pia * ci.radius
}
func (re *Rectangle) Area() float64 {
	if re.length <= 0 || re.width <= 0 {
		return 0
	}
	return re.length * re.width
}

func (re *Rectangle) Perimeter() float64 {
	if re.length <= 0 || re.width <= 0 {
		return 0
	}
	return 2 * (re.length + re.width)
}

func main16() {
	Circle1 := Circle{radius: 10, pia: 3.14}
	Rectangle1 := Rectangle{length: 10, width: 5}
	var c Shape
	var r Shape
	c = &Circle1
	r = &Rectangle1
	println("Circle Area:", c.Area())
	println("Circle Perimeter:", c.Perimeter())
	println("Rectangle Area:", r.Area())
	println("Rectangle Perimeter:", r.Perimeter())
}
