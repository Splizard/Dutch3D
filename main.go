package main

import (
	"github.com/g3n/engine/util/application"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	. "github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/window"

	"fmt"
	"os"
	"strconv"
)

func ToVector(lat, lon float32) *Vector3 {
	lat *= Pi / 180
	lon *= Pi / 180

	lat = -lat
	return &Vector3{
		Cos(lon) * Cos(lat),

		Sin(lat),
		Sin(lon) * Cos(lat),
	}
}

func main() {
	if len(os.Args) <= 3 {
		fmt.Println("Usage: Dutch3d [display] lat lon ")
		return
	}

	app, _ := application.Create(application.Options{
		Title:  "Dutch 3D",
		Width:  800,
		Height: 600,
	})

	//This is the texture for the globe.
	texDay, err := texture.NewTexture2DFromImage("data/textures/Surface.mat_baseColor.jpeg")
	if err != nil {
		app.Log().Fatal("Error loading texture: %s", err)
	}
	texDay.SetFlipY(true)

	// Create sphere
	matEarth := material.NewPhong(&Color{1, 1, 1})
	matEarth.AddTexture(texDay)
	matEarth.SetShininess(20)

	geom := geometry.NewSphere(1, 32, 32, 0, Pi*2, 0, Pi)
	sphere := graphic.NewMesh(geom, matEarth)
	app.Scene().Add(sphere)

	// Add light to the scene
	ambientLight := light.NewAmbient(&Color{1.0, 1.0, 1.0}, 0.8)
	app.Scene().Add(ambientLight)

	//Marker, this is a small sphere that will be placed at the lat & lon
	red := material.NewPhong(&Color{1, 0, 0})
	geom = geometry.NewSphere(0.01, 32, 32, 0, Pi*2, 0, Pi)
	marker := graphic.NewMesh(geom, red)

	//Circle
	geom2 := geometry.NewCircle(0.1, 50)
	mat2 := material.NewStandard(&Color{0, 0, 1})
	circle := graphic.NewMesh(geom2, mat2)

	//Parse the Latitude & longitude.
	if len(os.Args) > 3 {

		lat, err := strconv.ParseFloat(os.Args[2], 32)
		if err != nil {
		}
		lon, err := strconv.ParseFloat(os.Args[3], 32)
		if err != nil {
		}

		println(lat, lon)

		//Convert lat & lon to a position vector.
		vec := ToVector(float32(lat), float32(lon))

		marker.SetPositionVec(vec)

		var rotation Quaternion
		rotation.SetFromUnitVectors(&Vector3{0, 0, -1}, vec.Negate())

		circle.SetRotationQuat(&rotation)

		marker.Add(circle)
	}

	// Subscribe to before render events to call current test Render method
	app.Subscribe(application.OnBeforeRender, func(evname string, ev interface{}) {

		//Updates.

	})

	app.Window().Subscribe(window.OnKeyDown, func(evname string, ev interface{}) {
		key := ev.(*window.KeyEvent)

		if key.Keycode == window.KeySpace {
			//Print something on space key.
		}
	})

	app.Scene().Add(marker)
	app.CameraPersp().SetPosition(0, 0, 10)
	app.Run()
}
