package main

import "fmt"

type Point struct { x,y float64}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y+q.y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.x-q.x, p.y-q.y}
}

type Path []Point

func main() {
	var path = Path{
		{1,3},
		{2,4},
		{3,5},
	}
	path.pathByFuncVal(Point{1,1}, true)
	for i := range path{
		fmt.Println(path[i])
	}
}

func (path Path) pathByFuncVal(offset Point, flag bool) {
	var op func(p, q Point) Point
	if flag {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range  path{
		path[i] = op(path[i], offset)
	}
}
