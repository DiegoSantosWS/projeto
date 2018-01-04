package manipulador

import (
	"html/template"
)

//Modelos de tamplates
var Modelos = template.Must(template.ParseFiles("html/tmplates/home.html"))
