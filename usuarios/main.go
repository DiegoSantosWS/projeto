package main

import (
	r "github.com/DiegoSantosWS/usuarios/routers"
	"github.com/labstack/echo/middleware"

	pongor "github.com/MarcusMann/pongor-echo"
)

func main() {
	e := r.App

	p := pongor.GetRenderer(pongor.PongorOption{Reload: true})
	p.Directory = "views"
	e.Renderer = p

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":3000"))
}
