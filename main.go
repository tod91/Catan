package main

import (
	"Catan/gl"
	"runtime"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	gl.CreateWindow(640, 500, "testing")
	gl.SetViewPort(0, 0, 500, 400)

	shader := gl.NewShader()
	shader.LoadAndCompile("gl/shaders/vertex.sh", "gl/shaders/fragment.sh")
	defer gl.CleanUp()

	gl.Loop(shader)
}
