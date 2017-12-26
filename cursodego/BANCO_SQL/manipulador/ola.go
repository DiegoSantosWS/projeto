package manipulador

import (
	"github.com/DiegoSantosWS/cursodego/SERVIDOR/model"
	"fmt"
	"net/http"
	"time"
)

func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora

	if err := ModeloOla.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro na renderização da pagina", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execução do modelo: ", err.Error())
	}
}