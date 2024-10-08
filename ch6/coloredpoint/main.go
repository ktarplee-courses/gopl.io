// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 161.

// Coloredpoint demonstrates struct embedding.
package main

import (
	"fmt"
	"image/color"
	"math"
)

//!+decl

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

//!-decl

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (cp ColoredPoint) Colored() string {
	return "hola"
}

// var p *Point

// p.Distance(p2)
// (*p).Distance(p2)

// f := (*Point).ScaleBy
// f(p, 4)

func main() {
	//!+main
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"

	pp := &p
	pp.Colored()
	//!-main

	fmt.Println(q.Color)
}

/*
//!+error
	p.Distance(q) // compile error: cannot use q (ColoredPoint) as Point
//!-error
*/

var DistanceBetweenPoints = Point.Distance

func init() {
	//!+methodexpr
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)            // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
	//!-methodexpr

	dis := p.Distance // method value
	dis(q)

	findMinDistance(func(q Point) float64 { return p.Distance(q) }, []Point{})
	findMinDistance(p.Distance, []Point{})
}

// func findMinDistance(func(q Point) float64, []Point) *Point

func findMinDistance(dis func(q Point) float64, points []Point) *Point {
	var point *Point
	var minDis *float64
	for i, p := range points {
		d := dis(p)
		if minDis == nil || *minDis < d {
			minDis = &d
			point = &points[i]
		}
	}
	return point
}

func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	//!+indirect
	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}

	// no possible
	// func (p *ColoredPoint) MyDistance() {}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point                 // p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
	//!-indirect

	fmt.Println(q.Color)
}
