package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"github.com/0x-chaitu/rag_erp/internal/domain"
	"github.com/0x-chaitu/rag_erp/internal/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type api struct {
	httpClient *http.Client
	userRepo   domain.UserRepo
	orgRepo    domain.OrgRepo
	authClient *auth.Client
}

func NewAPI(
	ctx context.Context, pool *pgxpool.Pool,
	authClient *auth.Client) *api {
	client := &http.Client{}

	userRepo := repository.NewPostgresUser(pool)
	orgRepo := repository.NewPostgresOrg(pool)

	return &api{
		httpClient: client,

		// repos
		userRepo: userRepo,
		orgRepo:  orgRepo,

		authClient: authClient,
	}
}

func (a *api) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Route("/v1", func(r chi.Router) {
		r.Route("/admin", func(r chi.Router) {
			r.Mount("/orgs", a.initAdminOrgsRoutes())
			r.Mount("/users", a.initAdminUsersRoutes())

		})
	})

	return r
}

func (a *api) Server(port int) *http.Server {
	return &http.Server{
		Addr:    ":" + fmt.Sprintf("%v", port),
		Handler: a.Routes(),
		// 	ReadTimeout:    cfg.HTTP.ReadTimeout,
		// 	WriteTimeout:   cfg.HTTP.WriteTimeout,
		// 	MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
	}
}
