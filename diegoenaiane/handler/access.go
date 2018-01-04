package handler

import (
	"log"
	"net/http"

	"github.com/DiegoSantosWS/diegoenaiane/lib/context"
	"github.com/DiegoSantosWS/diegoenaiane/lib/query"
	"github.com/DiegoSantosWS/diegoenaiane/model"
	"github.com/DiegoSantosWS/diegoenaiane/repo"
)

func ListAccessPageVeic(ctx *context.Context) {
	ctx.Data["rows"] = 0
	ctx.HTML(http.StatusOK, "list")
}

func ListAccessPage(ctx *context.Context) {
	ctx.Data["rows"] = 0
	ctx.HTML(http.StatusOK, "list")
}

func ListAcessBy(ctx *context.Context) {
	access, err := list(ctx)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ctx.Data["access"] = access
	ctx.Data["rows"] = len(access)
	ctx.HTML(http.StatusOK, "list")
}

func ListAccessForApi(ctx *context.Context) {
	access, err := list(ctx)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
	}
	ctx.JSON(http.StatusOK, access)
}

func list(ctx *context.Context) ([]model.Access, error) {
	fields := make(map[string]string)
	if marca := ctx.Query("marca"); marca != "" {
		fields["marca"] = marca
	}

	if modelo := ctx.Query("modelo"); modelo != "" {
		fields["modelo"] = modelo
	}

	if combustivel := ctx.Query("combustivel"); combustivel != "" {
		fields["combustivel"] = combustivel
	}

	if propriedade := ctx.Query("propriedade"); propriedade != "" {
		fields["propriedade"] = propriedade
	}

	if ano := ctx.Query("ano"); ano != "" {
		fields["ano"] = ano
	}

	if dataInicio := ctx.Query("dataInicio"); dataInicio != "" {
		fields["dataInicio"] = dataInicio
	}

	if dataFim := ctx.Query("dataFim"); dataFim != "" {
		fields["dataFim"] = dataFim
	}

	repo, err := repository.NewAccessRepository()
	if err != nil {
		return nil, err
	}
	access, err := repo.FindAllBy(query.Build(fields), ctx.Cache)
	if err != nil {
		return nil, err
	}
	return access, nil
}
