package gl

import "github.com/go-gl/glfw/v3.3/glfw"

func init() {
	err := glfw.Init()

	if err != nil {
		panic(err)
	}
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

func Loop() {
	for !window.ShouldClose() {
		//ProcessInput()
		SetBackgroundColor(0.2, 0.1, 1, 1)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var window *glfw.Window
