package main

import (
	"Catan/resources"
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
	"time"
)

func main() {

	// Create application and scene
	a := app.App()
	scene := core.NewNode()

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Create perspective camera
	cam := camera.New(1)
	cam.SetPosition(6, 0, 12)
	scene.Add(cam)

	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	geom := geometry.NewPlane(6, 6)
	mat := material.NewStandard(math32.NewColor("white"))
	mat.AddTexture(resources.Get["rock_field"])
	mat.SetBlending(material.BlendNormal)
	mat.SetSide(material.SideDouble)
	mesh := graphic.NewMesh(geom, mat)
	scene.Add(mesh)

	plane := geometry.NewPlane(4, 6)
	mat2 := material.NewStandard(math32.NewColor("white"))
	mat2.AddTexture(resources.Get["clay"])
	mat2.SetBlending(material.BlendNormal)
	mat2.SetSide(material.SideDouble)
	mesh2 := graphic.NewMesh(plane, mat2)
	mesh2.SetPositionX(6)
	scene.Add(mesh2)

	//skyboxData := graphic.SkyboxData{DirAndPrefix: "resources/pics/", Suffixes: [6]string{"left", "right", "top", "bottom", "back", "front"}, Extension: "jpg"}
	//skybox, er := graphic.NewSkybox(skyboxData)
	//if er != nil {
	//	panic(er)
	//}
	cube := geometry.NewCube(6)
	vbo := cube.GetGeometry().VBOs()
	test := vbo[2].Buffer()
	test.SetVector2(0, &math32.Vector2{X: 0.25, Y: 0.66})
	test.SetVector2(2, &math32.Vector2{X: 0.25, Y: 0.33})
	test.SetVector2(4, &math32.Vector2{X: 0, Y: 0.66})
	test.SetVector2(6, &math32.Vector2{X: 0, Y: 0.33})

	test.SetVector2(8, &math32.Vector2{X: 0.5, Y: 0.66})
	test.SetVector2(10, &math32.Vector2{X: 0.5, Y: 0.33})
	test.SetVector2(12, &math32.Vector2{X: 0.25, Y: 0.66})
	test.SetVector2(14, &math32.Vector2{X: 0.25, Y: 0.33})

	test.SetVector2(16, &math32.Vector2{X: 0.75, Y: 0.66})
	test.SetVector2(18, &math32.Vector2{X: 0.75, Y: 0.33})
	test.SetVector2(20, &math32.Vector2{X: 0.5, Y: 0.66})
	test.SetVector2(22, &math32.Vector2{X: 0.5, Y: 0.33})

	test.SetVector2(24, &math32.Vector2{X: 1, Y: 0.66})
	test.SetVector2(26, &math32.Vector2{X: 1, Y: 0.33})
	test.SetVector2(28, &math32.Vector2{X: 0.75, Y: 0.66})
	test.SetVector2(30, &math32.Vector2{X: 0.75, Y: 0.33})

	test.SetVector2(32, &math32.Vector2{X: 0.5, Y: 1.0})
	test.SetVector2(34, &math32.Vector2{X: 0.5, Y: 0.66})
	test.SetVector2(36, &math32.Vector2{X: 0.25, Y: 1.0})
	test.SetVector2(38, &math32.Vector2{X: 0.25, Y: 0.66})

	test.SetVector2(40, &math32.Vector2{X: 0.5, Y: 0.33})
	test.SetVector2(42, &math32.Vector2{X: 0.5, Y: 0.0})
	test.SetVector2(44, &math32.Vector2{X: 0.25, Y: 0.33})
	test.SetVector2(46, &math32.Vector2{X: 0.25, Y: 0.0})

	mat3 := material.NewStandard(math32.NewColor("white"))
	mat3.AddTexture(resources.Get["die"])
	mat3.SetBlending(material.BlendNormal)
	mat3.SetSide(material.SideFront)
	mesh3 := graphic.NewMesh(cube, mat3)
	mesh3.SetPositionX(12)

	scene.Add(mesh3)

	// Create and add lights to the scene
	scene.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8))
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	scene.Add(pointLight)

	// Create and add an axis helper to the scene
	node := helper.NewAxes(0.5)
	scene.Add(node)

	// Set background color to gray
	a.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}
