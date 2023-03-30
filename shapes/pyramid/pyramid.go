// pyramid/pyramid.go
package pyramid

import (
    "math"
    "github.com/baduthutan/Popular-Programming-Language/shapes"
)

type Pyramid struct {
    BaseLength float64
    BaseWidth  float64
    Height     float64
}

func (p Pyramid) Volume() float64 {
    return (p.BaseLength * p.BaseWidth * p.Height) / 3
}

func (p Pyramid) Area() float64 {
    baseArea := p.BaseLength * p.BaseWidth
    slantHeight := math.Sqrt(math.Pow(p.BaseLength/2, 2) + math.Pow(p.Height, 2))
    triangleArea := (p.BaseLength * slantHeight) / 2
    return baseArea + (4 * triangleArea)
}

func NewPyramid(baseLength, baseWidth, height float64) shapes.Shape {
    return Pyramid{BaseLength: baseLength, BaseWidth: baseWidth, Height: height}
}
