package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"log"
	"os"
)

func NewShader() *Shader {
	sh := &Shader{vbo: 0}
	gl.GenBuffers(1, &sh.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, sh.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	return sh
}

func (sh *Shader) LoadAndCompile(filename string) {

	shaderSrc, err := readShaderFile(filename)
	if err != nil {
		panic(err)
	}

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	content, free := gl.Strs(shaderSrc)
	gl.ShaderSource(vertexShader, 1, content, nil)
	free()
	gl.CompileShader(vertexShader)
}

func readShaderFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content), err
}

type Shader struct {
	vbo uint32
}

var vertices = []float32{
	-0.5, -0.5, 0.0,
	0.5, -0.5, 0.0,
	0.0, 0.5, 0.0,
}
