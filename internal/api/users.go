package api

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/0x-chaitu/rag_erp/internal/domain"
	"github.com/0x-chaitu/rag_erp/pkg/utils"
	"github.com/go-chi/chi"
)

func (a *api) initAdminUsersRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", a.adminCreateUser)
	// r.Post("/signin", h.userSignIn)
	// r.Get("/user/{userId}", h.getUser)
	// r.Route("/", func(r chi.Router) {
	// 	r.Use(h.parseUser)
	// 	r.Put("/update", h.userUpdate)
	// 	r.Delete("/delete", h.deleteUser)
	// })
	return r

}

type UserParams struct {
	// *domain.User
	IdToken   string `json:"idToken"`
	Tenant_id string `json:"tenant_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (a *api) adminCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := utils.ReadJSON[UserParams](r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	tenantClient, err := a.authClient.TenantManager.AuthForTenant(body.Tenant_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	u := auth.UserToCreate{}

	u.Email(body.Email)
	u.Password(body.Password)

	res, err := tenantClient.CreateUser(ctx, &u)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	id, err := a.userRepo.Create(ctx, &domain.User{
		GoogleUserID: res.UID,
		OrgId:        body.Tenant_id,
	})
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, domain.User{
		ID:           id,
		GoogleUserID: res.UID,
	})
}
