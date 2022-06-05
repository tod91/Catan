package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, gl.TRUE)
}

func CleanUp() {
	glfw.Terminate()
}

func CreateWindow(w int, h int, title string) {
	var err error

	window, err = glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
}

//func ProcessInput(window *glfw.Window) {
//	if window.GetKey(glfw.KeyEscape) == glfw.Press {
//		window.SetShouldClose(true)
//	}
//}

func Loop(sh *Shader) {

	for !window.ShouldClose() {
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		// draw our first triangle
		gl.UseProgram(sh.ShaderProgram)
		gl.BindVertexArray(sh.VAO) // seeing as we only have a single VAO there's no need to bind it every time, but we'll do so to keep things a bit more organized
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		//SetBackgroundColor(0.2, 0.1, 1, 1)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var window *glfw.Window
