package manipulador

import "html/template"

//ModeloOla executa template
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal  executa template
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
