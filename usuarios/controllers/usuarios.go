package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DiegoSantosWS/usuarios/models"
	"github.com/labstack/echo"
)

//Home é a pagina inicial da minha apllication
func Home(c echo.Context) error {

	var usuarios []models.Usuarios
	if err := models.UsuarioModel.Find().All(&usuarios); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"menssgem": "Erro ao tentar recuperar os dados",
		})
	}
	data := map[string]interface{}{
		"titulo":   "Listagem de usuarios",
		"usuarios": usuarios,
	}

	return c.Render(http.StatusOK, "index.html", data)
}

//Inserir salva um registro no banco
func Inserir(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario models.Usuarios
	usuario.Nome = nome
	usuario.Email = email

	if nome != "" && email != "" {
		if _, err := models.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"menssgem": "Não foi possível add o registro no banco, tente novamente",
			})
		}
		return c.JSON(http.StatusCreated, map[string]string{
			"menssgem": "O registro fio aramazenado com sucesso",
		})

	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"menssgem": "Os campos precisa se passados",
	})
}

//Deletar salva um registro no banco
func Deletar(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(usuarioID)
	resultado := models.UsuarioModel.Find("id=?", usuarioID)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"menssagem": "Usuário não deletado!",
		})
	}

	if err := resultado.Delete(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"menssagem": "Não foi possivel delatar o usuário",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"menssagem": "Usuário deletado com sucesso!",
	})
}
