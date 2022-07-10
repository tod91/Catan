package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"neilpa.me/go-stbi"
)

var Texture uint32

func InitTextures() {

	gl.GenTextures(1, &Texture)
	gl.BindTexture(gl.TEXTURE_2D, Texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	image, err := stbi.Load("resource/wall.jpeg")
	if err != nil {
		panic(err)
	}

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, 512, 512, 0, gl.RGB, gl.UNSIGNED_BYTE, gl.Ptr(image.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
}
