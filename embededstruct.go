package main
import "fmt"

type Point struct {
	x int
	y int
}

type Circle struct {
	radius int
	center *Point
}
func changeX (p *Point, xcor int){
	p.x = xcor
}
func main(){
	p1 := Point{4,6}
	fmt.Printf("P1: %+v\n", p1);
	p2 := &Point{y:3}
	fmt.Printf("P2: %+v\n", p2);

	changeX(p2, 10)
	fmt.Printf("P2: %+v\n", *p2);

	p3 := &Point{3,3}
	c1 := Circle{6, p3}
	fmt.Println(c1)
	fmt.Printf("Circle\n Radius: %v\n Centre:%+v\n", c1.radius, *c1.center)

}