package main_game

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
	"math/rand"
	"time"
)

func LoadScene(scene *core.Node) {
	//loadDessert(scene)
	loadTile(scene, "desert", &math32.Vector3{0, 0, 0})

	arr := []uint{0, 1, 2, 3, 4, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3}
	arr = shuffleBoard(arr)
	loadTile(scene, getTileFromNum(arr[0]), &math32.Vector3{0, -2.92, 0})
	loadTile(scene, getTileFromNum(arr[1]), &math32.Vector3{0, 2.92, 0})
	loadTile(scene, getTileFromNum(arr[2]), &math32.Vector3{0.86, 1.46, 0})
	loadTile(scene, getTileFromNum(arr[3]), &math32.Vector3{-0.86, 1.46, 0})
	loadTile(scene, getTileFromNum(arr[4]), &math32.Vector3{-0.86, -1.46, 0})
	loadTile(scene, getTileFromNum(arr[5]), &math32.Vector3{0.86, -1.46, 0})
	loadTile(scene, getTileFromNum(arr[6]), &math32.Vector3{1.73, 0, 0})
	loadTile(scene, getTileFromNum(arr[7]), &math32.Vector3{-1.73, 0, 0})
	loadTile(scene, getTileFromNum(arr[8]), &math32.Vector3{1.73, 2.92, 0})
	loadTile(scene, getTileFromNum(arr[9]), &math32.Vector3{-1.73, 2.92, 0})
	loadTile(scene, getTileFromNum(arr[10]), &math32.Vector3{1.73, -2.92, 0})
	loadTile(scene, getTileFromNum(arr[11]), &math32.Vector3{-1.73, -2.92, 0})
	loadTile(scene, getTileFromNum(arr[12]), &math32.Vector3{2.595, 1.46, 0})
	loadTile(scene, getTileFromNum(arr[13]), &math32.Vector3{-2.595, 1.46, 0})
	loadTile(scene, getTileFromNum(arr[14]), &math32.Vector3{2.595, -1.46, 0})
	loadTile(scene, getTileFromNum(arr[15]), &math32.Vector3{-2.595, -1.46, 0})
	loadTile(scene, getTileFromNum(arr[16]), &math32.Vector3{-3.46, 0, 0})
	loadTile(scene, getTileFromNum(arr[17]), &math32.Vector3{3.46, 0, 0})

	//loadCubeMap(scene, 18)
}

func loadTile(scene *core.Node, name string, pos *math32.Vector3) {

	dec, err := obj.Decode("resources/blender/"+name+".obj", "resources/blender/"+name+".mtl")
	if err != nil {
		panic(err)
	}

	tile, err := dec.NewGroup()
	if err != nil {
		panic(err)
	}

	tile.TranslateX(pos.X)
	tile.TranslateY(pos.Y)
	tile.TranslateY(pos.Z)

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
}

func shuffleBoard(arr []uint) []uint {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func getTileFromNum(num uint) string {
	switch num {
	case 0:
		return "clay"
	case 1:
		return "grain"
	case 2:
		return "rock"
	case 3:
		return "wood"
	case 4:
		return "sheep"
	default:
		return ""
	}
}
