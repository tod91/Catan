package main_game

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
)

func LoadScene(scene *core.Node) {
	loadDessert(scene)
	loadTile(scene, "sheep", &math32.Vector3{0, -2.92, 0})
	loadTile(scene, "clay", &math32.Vector3{0, 2.92, 0})
	loadTile(scene, "sheep", &math32.Vector3{0.86, 1.46, 0})
	loadTile(scene, "sheep", &math32.Vector3{-0.86, 1.46, 0})
	loadTile(scene, "grain", &math32.Vector3{-0.86, -1.46, 0})
	loadTile(scene, "sheep", &math32.Vector3{0.86, -1.46, 0})
	loadTile(scene, "rock", &math32.Vector3{1.73, 0, 0})
	loadTile(scene, "wood", &math32.Vector3{-1.73, 0, 0})
	loadTile(scene, "grain", &math32.Vector3{1.73, 2.92, 0})
	loadTile(scene, "wood", &math32.Vector3{-1.73, 2.92, 0})
	loadTile(scene, "grain", &math32.Vector3{1.73, -2.92, 0})
	loadTile(scene, "clay", &math32.Vector3{-1.73, -2.92, 0})
	loadTile(scene, "sheep", &math32.Vector3{2.595, 1.46, 0})
	loadTile(scene, "rock", &math32.Vector3{-2.595, 1.46, 0})
	loadTile(scene, "wood", &math32.Vector3{2.595, -1.46, 0})
	loadTile(scene, "clay", &math32.Vector3{-2.595, -1.46, 0})
	loadTile(scene, "rock", &math32.Vector3{-3.46, 0, 0})
	loadTile(scene, "grain", &math32.Vector3{3.46, 0, 0})

	//loadCubeMap(scene, 18)
}

func loadTile(scene *core.Node, name string, pos *math32.Vector3) {

	dec, err := obj.Decode("resources/blender/"+name+".obj", "resources/blender/"+name+".mtl")
	if err != nil {
		panic(err)
	}

	desert, err := dec.NewGroup()
	if err != nil {
		panic(err)
	}
	desert.SetScaleVec(&math32.Vector3{1.0099, 1.0099, 1.0099})

	tile, err := dec.NewGroup()
	if err != nil {
		panic(err)
	}

	tile.TranslateX(pos.X)
	tile.TranslateY(pos.Y)
	tile.TranslateY(pos.Z)

	scene.Add(tile)
}

func loadDessert(scene *core.Node) {

	dec, err := obj.Decode("resources/blender/desert.obj", "resources/blender/desert.mtl")
	if err != nil {
		panic(err)
	}

	tile, err := dec.NewGroup()
	if err != nil {
		panic(err)
	}
	scene.Add(tile)
}

func loadCubeMap(scene *core.Node, posX float32) {

	dec, err := obj.Decode("resources/blender/desert.obj", "resources/blender/desert.mtl")
	if err != nil {
		panic(err)
	}

	desert, err := dec.NewGroup()
	if err != nil {
		panic(err)
	}
	desert.SetScaleVec(&math32.Vector3{1.01, 1.0099, 1.0099})

	dec2, err := obj.Decode("resources/blender/sheep.obj", "resources/blender/sheep.mtl")
	if err != nil {
		panic(err)
	}

	sheep, err := dec2.NewGroup()
	if err != nil {
		panic(err)
	}

	sheep.TranslateX(0.85)
	sheep.TranslateY(1.5)

	scene.Add(desert)
	scene.Add(sheep)

	//cube := geometry.NewCube(6)
	//vbo := cube.GetGeometry().VBOs()
	//uvCoords := vbo[2].Buffer()
	//uvCoords.SetVector2(0, &math32.Vector2{X: 0.25, Y: 0.66})
	//uvCoords.SetVector2(2, &math32.Vector2{X: 0.25, Y: 0.33})
	//uvCoords.SetVector2(4, &math32.Vector2{X: 0, Y: 0.66})
	//uvCoords.SetVector2(6, &math32.Vector2{X: 0, Y: 0.33})
	//
	//uvCoords.SetVector2(8, &math32.Vector2{X: 0.5, Y: 0.66})
	//uvCoords.SetVector2(10, &math32.Vector2{X: 0.5, Y: 0.33})
	//uvCoords.SetVector2(12, &math32.Vector2{X: 0.25, Y: 0.66})
	//uvCoords.SetVector2(14, &math32.Vector2{X: 0.25, Y: 0.33})
	//
	//uvCoords.SetVector2(16, &math32.Vector2{X: 0.75, Y: 0.66})
	//uvCoords.SetVector2(18, &math32.Vector2{X: 0.75, Y: 0.33})
	//uvCoords.SetVector2(20, &math32.Vector2{X: 0.5, Y: 0.66})
	//uvCoords.SetVector2(22, &math32.Vector2{X: 0.5, Y: 0.33})
	//
	//uvCoords.SetVector2(24, &math32.Vector2{X: 1, Y: 0.66})
	//uvCoords.SetVector2(26, &math32.Vector2{X: 1, Y: 0.33})
	//uvCoords.SetVector2(28, &math32.Vector2{X: 0.75, Y: 0.66})
	//uvCoords.SetVector2(30, &math32.Vector2{X: 0.75, Y: 0.33})
	//
	//uvCoords.SetVector2(32, &math32.Vector2{X: 0.5, Y: 1.0})
	//uvCoords.SetVector2(34, &math32.Vector2{X: 0.5, Y: 0.66})
	//uvCoords.SetVector2(36, &math32.Vector2{X: 0.25, Y: 1.0})
	//uvCoords.SetVector2(38, &math32.Vector2{X: 0.25, Y: 0.66})
	//
	//uvCoords.SetVector2(40, &math32.Vector2{X: 0.5, Y: 0.33})
	//uvCoords.SetVector2(42, &math32.Vector2{X: 0.5, Y: 0.0})
	//uvCoords.SetVector2(44, &math32.Vector2{X: 0.25, Y: 0.33})
	//uvCoords.SetVector2(46, &math32.Vector2{X: 0.25, Y: 0.0})

	//mat := material.NewStandard(math32.NewColor("white"))
	//mat.AddTexture(resources.Get["die"])
	//mat.SetBlending(material.BlendNormal)
	//mat.SetSide(material.SideFront)
	//mesh3 := graphic.NewMesh(cube, mat)
	//mesh3.SetPositionX(posX)

}
