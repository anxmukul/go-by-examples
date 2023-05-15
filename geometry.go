// Another example of Interface

package main
import ("fmt"
		"math"
)
type geometry interface {
	calculateArea() float64
	calculatePerimeter() float64
}
type rect struct{
	length float64
	width float64

}
type circle struct {
	radius float64
}
func (r rect) calculateArea() float64  {
	area := r.length * r.width
	// fmt.Println("Area of rectangle: ", area)
	return area;

}
func (r rect) calculatePerimeter() float64 {
	peri := 2*(r.length + r.width)
	// fmt.Println("Perimeter of rectangle: ", peri)
	return peri
}
func (c circle) calculateArea() float64 {
	area := math.Pi * c.radius * c.radius
	// fmt.Println("Area of circle: ", area)
	return area
	
}
func (c circle) calculatePerimeter() float64 {
	peri := math.Pi*c.radius*2
	// fmt.Println("Perimeter of circle is: ", peri)
	return peri
}

func measure (g []geometry) {
	for _, shp := range g {
		fmt.Printf("Area of geometry %+v is: %v\n",shp, shp.calculateArea())
		fmt.Printf("Perimeter of geometry %+v is: %v\n",shp, shp.calculatePerimeter())
		fmt.Println("----------")
	}
	
}
func main(){
	rect1 := rect{length:5, width: 9}
	rect2 := rect{length:6, width: 3}

	circle1 := circle{radius: 7}
	circle2 := circle{radius: 2}

	shapes := []geometry{rect1, rect2, circle1, circle2}
	measure(shapes)

}