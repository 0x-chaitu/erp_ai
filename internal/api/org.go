package api

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/0x-chaitu/rag_erp/internal/domain"
	"github.com/0x-chaitu/rag_erp/pkg/utils"
	"github.com/go-chi/chi"
)

func (a *api) initAdminOrgsRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", a.adminCreateOrg)

	return r

}

func (a *api) adminCreateOrg(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := utils.ReadJSON[domain.Organization](r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	tenant := auth.TenantToCreate{}
	tenant.AllowPasswordSignUp(true)
	tenant.DisplayName(body.Name)

	res, err := a.authClient.TenantManager.CreateTenant(ctx, &tenant)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	body.OrgID = res.ID
	err = a.orgRepo.Create(ctx, &body)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, body)
}
