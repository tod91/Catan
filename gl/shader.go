package gl

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"log"
	"os"
	"strings"
)

func NewShader() *Shader {
	sh := &Shader{vbo: 0}
	gl.GenBuffers(1, &sh.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, sh.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	return sh
}

func (sh *Shader) LoadAndCompile(vertexFile, fragmentFile string) {

	// build vertex shader
	shaderSrc, err := readShaderFile(vertexFile)
	if err != nil {
		panic(err)
	}

	gl.GenVertexArrays(1, &sh.VAO)
	gl.BindVertexArray(sh.VAO)

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	content, freeVert := gl.Strs(shaderSrc)
	gl.ShaderSource(vertexShader, 1, content, nil)
	freeVert()
	gl.CompileShader(vertexShader)

	var status int32
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(vertexShader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(vertexShader, logLength, nil, gl.Str(log))

		fmt.Errorf("failed to compile %v: %v", shaderSrc, log)
	}

	// build fragment shader
	var fragmentSrc string
	fragmentSrc, err = readShaderFile(fragmentFile)
	if err != nil {
		panic(err)
	}

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	contentFragment, freeFrag := gl.Strs(fragmentSrc)
	gl.ShaderSource(fragmentShader, 1, contentFragment, nil)
	freeFrag()
	gl.CompileShader(fragmentShader)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(fragmentShader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(fragmentShader, logLength, nil, gl.Str(log))

		fmt.Errorf("failed to compile %v: %v", fragmentSrc, log)
	}

	// build shader program
	sh.ShaderProgram = gl.CreateProgram()
	gl.AttachShader(sh.ShaderProgram, vertexShader)
	gl.AttachShader(sh.ShaderProgram, fragmentShader)
	gl.LinkProgram(sh.ShaderProgram)

	gl.GetProgramiv(sh.ShaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(sh.ShaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(sh.ShaderProgram, logLength, nil, gl.Str(log))

		fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
	gl.EnableVertexAttribArray(0)
}

func readShaderFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content), err
}

type Shader struct {
	vbo           uint32
	VAO           uint32
	ShaderProgram uint32
}

var vertices = []float32{
	-0.5, -0.5, 0.0,
	0.5, -0.5, 0.0,
	0.0, 0.5, 0.0,
}
