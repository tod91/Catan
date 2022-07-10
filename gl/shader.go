package gl

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"log"
	"os"
	"strings"
)

func NewShader() *Shader {
	sh := &Shader{VBO: 0,
		VAO:           0,
		ShaderProgram: 0}

	gl.GenVertexArrays(1, &sh.VAO)
	gl.BindVertexArray(sh.VAO)
	gl.GenBuffers(1, &sh.VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, sh.VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.GenBuffers(1, &sh.EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, sh.EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3)
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	return sh
}

func (sh *Shader) LoadAndCompile(vertexFile, fragmentFile string) {

	// build vertex shader
	shaderSrc, err := readShaderFile(vertexFile)
	if err != nil {
		panic(err)
	}

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	content, freeVert := gl.Strs(shaderSrc + "\x00")
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
	contentFragment, freeFrag := gl.Strs(fragmentSrc + "\x00")
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

	InitTextures()
}

func readShaderFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content), err
}

type Shader struct {
	VBO           uint32
	VAO           uint32
	EBO           uint32
	ShaderProgram uint32
}

var vertices = []float32{
	//0.0, 0.0, 0.0, //center
	//-0.5, 1.0, 0.0, // left top
	//0.5, 1.0, 0.0, // right top
	//1.0, 0.0, 0.0, // right
	//0.5, -1.0, 0.0, // right bottom (notice sign)
	//-0.5, -1.0, 0.0, // left bottom
	//-1.0, 0.0, 0.0, // left
	0.5, 0.5, 0.0, 1.0, 1.0,
	0.5, -0.5, 0.0, 1.0, 0.0,
	-0.5, -0.5, 0.0, 0.0, 0.0,
	-0.5, 0.5, 0.0, 0.0, 1.0,
}
var indices = []uint32{0, 1, 2, 3, 4, 5, 6, 1}
