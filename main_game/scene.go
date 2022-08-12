package main_game

import (
	"Catan/resources"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

func LoadScene(scene *core.Node) {
	loadTile(scene, "desert", 6, 6, &math32.Vector3{0, 0, 0})
	loadTile(scene, "wood_field", 6, 6, &math32.Vector3{5.8, 0, 0})
	loadTile(scene, "rock_field", 6, 6, &math32.Vector3{11.5, 0, 0})
	loadTile(scene, "grain_field", 6, 6, &math32.Vector3{2.9, 4.3, 0})
	loadTile(scene, "sheep_field", 6, 6, &math32.Vector3{-5.7, 0, 0})
	loadTile(scene, "wood_field", 6, 6, &math32.Vector3{-11.4, 0, 0})
	loadTile(scene, "rock_field", 6, 6, &math32.Vector3{2.9, -4.3, 0})
	loadTile(scene, "grain_field", 6, 6, &math32.Vector3{-2.9, -4.3, 0})
	loadTile(scene, "sheep_field", 6, 6, &math32.Vector3{-2.9, 4.3, 0})
	loadTile(scene, "rock_field", 6, 6, &math32.Vector3{-8.6, -4.3, 0})
	loadTile(scene, "wood_field", 6, 6, &math32.Vector3{8.6, -4.3, 0})
	loadTile(scene, "clay_field", 6, 6, &math32.Vector3{-8.6, 4.3, 0})
	loadTile(scene, "grain_field", 6, 6, &math32.Vector3{8.6, 4.3, 0})
	loadTile(scene, "sheep_field", 6, 6, &math32.Vector3{5.8, 8.6, 0})
	loadTile(scene, "wood_field", 6, 6, &math32.Vector3{-5.8, 8.6, 0})
	loadTile(scene, "grain_field", 6, 6, &math32.Vector3{-5.8, -8.6, 0})
	loadTile(scene, "sheep_field", 6, 6, &math32.Vector3{5.8, -8.6, 0})
	loadTile(scene, "clay_field", 6, 6, &math32.Vector3{0, -8.6, 0})
	loadTile(scene, "clay_field", 6, 6, &math32.Vector3{0, 8.6, 0})

	//loadCubeMap(scene, 18)
}

func loadTile(scene *core.Node, name string, w float32, h float32, pos *math32.Vector3) {
	geom := geometry.NewPlane(w, h)
	mat := material.NewStandard(math32.NewColor("white"))
	mat.AddTexture(resources.Get[name])
	mat.SetBlending(material.BlendNormal)
	mat.SetSide(material.SideDouble)
	mat.SetDepthMask()
	mesh := graphic.NewMesh(geom, mat)
	if name == "desert" {
		scaleFactor := float32(0.95)
		mesh.SetScale(scaleFactor, scaleFactor, scaleFactor)
	}
	mesh.SetPositionVec(pos)
	scene.Add(mesh)
}

func loadCubeMap(scene *core.Node, posX float32) {
	cube := geometry.NewCube(6)
	vbo := cube.GetGeometry().VBOs()
	uvCoords := vbo[2].Buffer()
	uvCoords.SetVector2(0, &math32.Vector2{X: 0.25, Y: 0.66})
	uvCoords.SetVector2(2, &math32.Vector2{X: 0.25, Y: 0.33})
	uvCoords.SetVector2(4, &math32.Vector2{X: 0, Y: 0.66})
	uvCoords.SetVector2(6, &math32.Vector2{X: 0, Y: 0.33})

	uvCoords.SetVector2(8, &math32.Vector2{X: 0.5, Y: 0.66})
	uvCoords.SetVector2(10, &math32.Vector2{X: 0.5, Y: 0.33})
	uvCoords.SetVector2(12, &math32.Vector2{X: 0.25, Y: 0.66})
	uvCoords.SetVector2(14, &math32.Vector2{X: 0.25, Y: 0.33})

	uvCoords.SetVector2(16, &math32.Vector2{X: 0.75, Y: 0.66})
	uvCoords.SetVector2(18, &math32.Vector2{X: 0.75, Y: 0.33})
	uvCoords.SetVector2(20, &math32.Vector2{X: 0.5, Y: 0.66})
	uvCoords.SetVector2(22, &math32.Vector2{X: 0.5, Y: 0.33})

	uvCoords.SetVector2(24, &math32.Vector2{X: 1, Y: 0.66})
	uvCoords.SetVector2(26, &math32.Vector2{X: 1, Y: 0.33})
	uvCoords.SetVector2(28, &math32.Vector2{X: 0.75, Y: 0.66})
	uvCoords.SetVector2(30, &math32.Vector2{X: 0.75, Y: 0.33})

	uvCoords.SetVector2(32, &math32.Vector2{X: 0.5, Y: 1.0})
	uvCoords.SetVector2(34, &math32.Vector2{X: 0.5, Y: 0.66})
	uvCoords.SetVector2(36, &math32.Vector2{X: 0.25, Y: 1.0})
	uvCoords.SetVector2(38, &math32.Vector2{X: 0.25, Y: 0.66})

	uvCoords.SetVector2(40, &math32.Vector2{X: 0.5, Y: 0.33})
	uvCoords.SetVector2(42, &math32.Vector2{X: 0.5, Y: 0.0})
	uvCoords.SetVector2(44, &math32.Vector2{X: 0.25, Y: 0.33})
	uvCoords.SetVector2(46, &math32.Vector2{X: 0.25, Y: 0.0})

	mat := material.NewStandard(math32.NewColor("white"))
	mat.AddTexture(resources.Get["die"])
	mat.SetBlending(material.BlendNormal)
	mat.SetSide(material.SideFront)
	mesh3 := graphic.NewMesh(cube, mat)
	mesh3.SetPositionX(posX)
	scene.Add(mesh3)
}
