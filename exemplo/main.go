/**
Autor: Juan Jose Sanchez Garcia
Fecha: 18-05-2017
Version: 0.1 "Patricia"
Este script monta un servidor web que regresa informacion de usuario en Json a un
front end con jquery, la implementacion de login aqui descrita no es la indicada,
sin embargo sirve como introduccion o ejemplo.
*/

package main

/**
 Importamos las librerias necesarias del script:

 - _ "github.com/go-sql-driver/mysql"
	"database/sql"
	Manejo de base de datos y Mysql

- "github.com/gin-gonic/gin" Manejo de Endpoints (URLS) para entrega de Json
- "gopkg.in/gin-contrib/cors.v1" Manejo de CORS (Comunicacion entre servidores)
- "net/http" Libreria nativa para manejo de HTTP

*/
import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gin-contrib/cors.v1"
)

/**
Metodo que imprime error en caso de que falle
params @err Error variable de error de ejecucion
*/

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/**
Funcion principal
*/
func main() {
	//Conexion a base de datos Mysql
	db, err := sql.Open("mysql", "root:1234@/personas?charset=utf8")
	// Checamos si se pudo conectar a la base de datos, en caso de error, imprimimos el error en consola
	checkErr(err)

	// Creamos nuestro modelo de usuario, en este caso se asemeja a un objeto de usuario
	type User struct {
		Username string
		Password string
		Picture  string
	}
	// Creamos nuestro modelo de persona, para mostrar en la tabla
	type Person struct {
		Firstname string
		Lastname  string
		Picture   string
	}
	//Iniciamos nuestro controlador de rutas, es decir, asigna nuestras urls al servidor HTTP
	router := gin.Default()
	// Habilitamos CORS en nuestro servidor (Permite hacer solicitudes por fuera del servidor)
	router.Use(cors.Default())

	//Le decimos a nuestro servidor que cargue los archivos html de nuestro directorio de templates
	router.LoadHTMLGlob("templates/*.html")
	// Le decimos al servidor que cargue nuestros css en la url css
	router.Static("/css", "templates/css")
	// Le decimos al servidor que cargue nuestros js en la url js
	router.Static("/js", "templates/js")
	// Le decimos al servidor que cargue nuestras fuentes
	router.Static("/fonts", "templates/fonts")

	/**
	Asignamos una ruta ("http://localhost/") con un metodo GET, asignandole el template de login.html
	Es decir , desde nuestro explorador (IE,Firefox,Chrome), mediante el metodo GET, nos carga la pantalla
	de login
	*/
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	/**
	Asignamos una ruta a nuestro servidor, la cual buscara en la base de datos, dado los
	valores enviados mediante Jquery y retornara usuario y contraseña en caso de coincidir.
	Cada parametro se representa en la URL mediante :variable
	Ejemplo de URL: "http://localhost/login/username/Usuario1/password/Pass1"
	Donde:
	/username: Nombre del parametro
	/Usuario1: Valor enviado desde Jquery
	/password: Nombre del parametro
	/Pass1: Nombre enviado del parametro
	*/
	router.GET("/login/username/:username/password/:password", func(c *gin.Context) {
		//Creamos un contenedor para nuestro usuario, "Una instancia del objeto user"
		var (
			user   User
			result gin.H
		)
		//Asignamos valores a nuestros parametros
		username := c.Param("username")
		password := c.Param("password")

		//Declaramos la query que se ejecutara en Mysql
		row := db.QueryRow("select username, pass from persona where username = ? and pass = ?;", username, password)
		// Por cada fila, sacaremos el usuario y la contraseña y la asignaremos a un JSON
		err = row.Scan(&user.Username, &user.Password)
		if err != nil {
			/**
			En caso de no hayar ningun usuario que corresponda con las credenciales, mandara none,
			En caso de si encontrar un usuario que corresponda mandara su usuario y contraseña
			*/
			result = gin.H{
				"Username": "none",
				"Password": "none",
			}
		} else {
			result = gin.H{
				"Username": user.Username,
				"Password": user.Password,
			}
		}
		//Parseamos y devolvemos nuestros resultados en JSON
		c.JSON(http.StatusOK, result)
	})

	/**
	Asignamos una ruta a nuestro servidor, que devolvera una lista de nombres de persona y su foto de nuestra base de datos
	en formato JSON
	*/
	router.GET("/Listado", func(c *gin.Context) {

		//Creamos objetos para nuestros datos
		var (
			persona  Person   // Creamos un objeto de persona, para cada registro de la base de datos
			personas []Person // Creamos un arreglo de personas para guardar los resultados obtenidos de la base de datos
		)
		//Declaramos la query que se ejecutara en Mysql
		rows, err := db.Query("select firstname, lastname, photo from persona limit 200;")

		// Por cada fila, sacaremos el usuario y la contraseña y la asignaremos a un JSON
		for rows.Next() {
			err = rows.Scan(&persona.Firstname, &persona.Lastname, &persona.Picture)
			personas = append(personas, persona)

			// En caso de que  haya algun error lo imprimimos en consola
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		// Devolvemos los resultados en JSON
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"Data": personas,
		})
	})
	// Crea una ruta para el archivo de lista en html
	router.GET("/lista", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lista.html", nil)
	})

	//Corremos nuestro servidor en el Puerto 3000 (http://localhost:3000)
	router.Run(":3000")

}
