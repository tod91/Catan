package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"log"
)

func Init1() {
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
}

func SetVBO() {
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
}

func SetViewPort(x, y, w, h int32) {
	gl.Viewport(x, y, w, h)
}

func SetBackgroundColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

//func ProcessInput () {
//	gl.ProcessInput(window)
//}

var vbo uint32

var vertices = []float32{
	-0.5, -0.5, 0.0,
	0.5, -0.5, 0.0,
	0.0, 0.5, 0.0,
}
