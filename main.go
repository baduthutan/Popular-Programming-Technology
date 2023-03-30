// main.go
package main

import (
    "fmt"
    "github.com/baduthutan/Popular-Programming-Language/ball"
    "github.com/baduthutan/Popular-Programming-Language/cylinder"
    "github.com/baduthutan/Popular-Programming-Language/pyramid"
)

func main() {
    ball := ball.NewBall(5)
    fmt.Println("Ball volume:", ball.Volume())
    fmt.Println("Ball area:", ball.Area())

    cylinder := cylinder.NewCylinder(3, 5)
    fmt.Println("Cylinder volume:", cylinder.Volume())
    fmt.Println("Cylinder area:", cylinder.Area())

    pyramid := pyramid.NewPyramid(4, 6, 8)
    fmt.Println("Pyramid volume:", pyramid.Volume())
    fmt.Println("Pyramid area:", pyramid.Area())
}
