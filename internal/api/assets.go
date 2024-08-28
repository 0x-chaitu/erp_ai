package api

import (
	"net/http"
	"strconv"

	"github.com/0x-chaitu/rag_erp/pkg/utils"
	"github.com/go-chi/chi"
)

func (a *api) initAssetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", a.getAssets)
	r.Get("/count", a.getAssetCount)

	return r

}

func (a *api) getAssets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	offsetP := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offsetP)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	limitP := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitP)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	body, err := a.assetRepo.GetAssets(ctx, offset, limit)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, body)
}

func (a *api) getAssetCount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := a.assetRepo.GetAssetCount(ctx)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, body)
}
