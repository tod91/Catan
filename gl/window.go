package gl

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
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

func CheckGLErrors() {
	glerror := gl.GetError()
	if glerror == gl.NO_ERROR {
		return
	}

	fmt.Printf("gl.GetError() reports")
	for glerror != gl.NO_ERROR {
		fmt.Printf(" ")
		switch glerror {
		case gl.INVALID_ENUM:
			fmt.Printf("GL_INVALID_ENUM")
		case gl.INVALID_VALUE:
			fmt.Printf("GL_INVALID_VALUE")
		case gl.INVALID_OPERATION:
			fmt.Printf("GL_INVALID_OPERATION")
		case gl.STACK_OVERFLOW:
			fmt.Printf("GL_STACK_OVERFLOW")
		case gl.STACK_UNDERFLOW:
			fmt.Printf("GL_STACK_UNDERFLOW")
		case gl.OUT_OF_MEMORY:
			fmt.Printf("GL_OUT_OF_MEMORY")
		default:
			fmt.Printf("%d", glerror)
		}
		glerror = gl.GetError()
	}
	fmt.Printf("\n")
}

func KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyE && action == glfw.Press {
		fmt.Println("E WAS PRESSED")
	}

}

func Loop(sh *Shader) {

	window.SetKeyCallback(KeyCallback)
	//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	var elapsed float32 = 0.1
	for !window.ShouldClose() {
		CheckGLErrors()
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		rotate := mgl32.Ident4()
		rotate = mgl32.HomogRotate3D(float32(glfw.GetTime()), mgl32.Vec3{1, 1, 1})
		gl.UseProgram(sh.ShaderProgram)
		rotateUniform := gl.GetUniformLocation(sh.ShaderProgram, gl.Str("rotate\x00"))
		gl.UniformMatrix4fv(rotateUniform, 1, false, &rotate[0])

		translate := mgl32.Ident4()

		translate = mgl32.Translate3D(elapsed, 0, 0)
		elapsed += 0.0001
		gl.UseProgram(sh.ShaderProgram)
		translateUniform := gl.GetUniformLocation(sh.ShaderProgram, gl.Str("translate\x00"))
		gl.UniformMatrix4fv(translateUniform, 1, false, &translate[0])

		gl.BindTexture(gl.TEXTURE_2D, sh.TextureID)
		gl.BindVertexArray(sh.VAO) // seeing as we only have a single VAO there's no need to bind it every time, but we'll do so to keep things a bit more organized
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
		gl.BindVertexArray(0)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var window *glfw.Window
