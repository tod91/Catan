package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"log"
)

func init() {
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
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
