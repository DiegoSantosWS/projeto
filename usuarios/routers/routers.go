package routers

import (
	"github.com/DiegoSantosWS/usuarios/controllers"
	"github.com/labstack/echo"
)

// App é uma instacia de echo
var App *echo.Echo

func init() {
	App = echo.New()
	//Pagina inicial da minha aplicação
	App.GET("/", controllers.Home)

	api := App.Group("/v1")
	api.POST("/insert", controllers.Inserir)
	api.DELETE("/delete/:id", controllers.Deletar)
}
