package resources

import "github.com/g3n/engine/texture"

type tx = texture.Texture2D

var Get = make(map[string]*tx)

func init() {
	clay, err := texture.NewTexture2DFromImage("resources/pics/clay.jpg")
	if err != nil {
		panic(err)
	}
	Get["clay"] = clay

	grain, err := texture.NewTexture2DFromImage("resources/pics/grain.jpg")
	if err != nil {
		panic(err)
	}
	Get["grain"] = grain

	rock, err := texture.NewTexture2DFromImage("resources/pics/rock.jpg")
	if err != nil {
		panic(err)
	}
	Get["rock"] = rock

	wood, err := texture.NewTexture2DFromImage("resources/pics/wood.jpg")
	if err != nil {
		panic(err)
	}
	Get["wood"] = wood

	clay_field, err := texture.NewTexture2DFromImage("resources/pics/clay_field.png")
	if err != nil {
		panic(err)
	}
	Get["clay_field"] = clay_field

	grain_field, err := texture.NewTexture2DFromImage("resources/pics/grain_field.png")
	if err != nil {
		panic(err)
	}
	Get["grain_field"] = grain_field

	rock_field, err := texture.NewTexture2DFromImage("resources/pics/rock_field.png")
	if err != nil {
		panic(err)
	}
	Get["rock_field"] = rock_field

	sheep_field, err := texture.NewTexture2DFromImage("resources/pics/sheep_field.png")
	if err != nil {
		panic(err)
	}
	Get["sheep_field"] = sheep_field

	wood_field, err := texture.NewTexture2DFromImage("resources/pics/wood_field.png")
	if err != nil {
		panic(err)
	}
	Get["wood_field"] = wood_field

	desert, err := texture.NewTexture2DFromImage("resources/pics/desert.png")
	if err != nil {
		panic(err)
	}
	Get["desert"] = desert

	die, err := texture.NewTexture2DFromImage("resources/pics/dice.png")
	if err != nil {
		panic(err)
	}
	Get["die"] = die
}
