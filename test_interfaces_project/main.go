package main

//defining two custom structs called square and triangle
type square struct {
	sideLength float64
}
type triangle struct {
	hight float64
	base  float64
}

//both structs have an area method with the same signature
func (s square) area() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) area() float64 {
	return 0.5 * t.hight * t.base
}
func main() {
	s1 := square{sideLength: 5}
	t1 := triangle{hight: 4, base: 6}

	//we can create a slice of empty interfaces and store both structs in it
	shapes := []interface{}{s1, t1}
	// fmt.Println(shapes)

	//we can use type assertion to access the area method of each struct
	for _, shape := range shapes {
		switch shape.(type) {
		case square:
			s := shape.(square)
			println("Square area:", s.area())
		case triangle:
			t := shape.(triangle)
			println("Triangle area:", t.area())
		default:
			println("Unknown shape")
		}
	}
}
