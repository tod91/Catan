package main

import (
	"Catan/gl"
)

func main() {
	gl.CreateWindow(640, 500, "testing")
	shader := gl.NewShader()
	shader.LoadAndCompile("gl/shaders/vertex.sh", "gl/shaders/fragment.sh")
	defer gl.CleanUp()

	gl.SetViewPort(0, 0, 500, 400)
	gl.Loop(shader)
}
