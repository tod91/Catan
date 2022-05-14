package main

import (
	"Catan/gl"
)

func main() {

	gl.CreateWindow(640, 500, "testing")
	defer gl.CleanUp()

	gl.Loop()
}
