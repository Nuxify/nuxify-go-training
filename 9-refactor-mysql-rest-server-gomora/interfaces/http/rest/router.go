/*
|--------------------------------------------------------------------------
| Router
|--------------------------------------------------------------------------
|
| This file contains the routes mapping and groupings of your REST API calls.
| See README.md for the routes UI server.
|
*/
package rest

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"rest-server/interfaces"
	"rest-server/interfaces/http/rest/middlewares/cors"
	"rest-server/interfaces/http/rest/viewmodels"
)

// ChiRouterInterface declares methods for the chi router
type ChiRouterInterface interface {
	InitRouter() *chi.Mux
	Serve(port int)
}

type router struct{}

var (
	m          *router
	routerOnce sync.Once
)

// InitRouter initializes main routes
func (router *router) InitRouter() *chi.Mux {
	// DI assignment
	userCommandController := interfaces.ServiceContainer().RegisterUserRESTCommandController()
	userQueryController := interfaces.ServiceContainer().RegisterUserRESTQueryController()
	postCommandController := interfaces.ServiceContainer().RegisterPostRESTCommandController()
	postQueryController := interfaces.ServiceContainer().RegisterPostRESTQueryController()

	// create router
	r := chi.NewRouter()

	// global and recommended middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(cors.Init().Handler)

	// default route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := viewmodels.HTTPResponseVM{
			Status:  http.StatusOK,
			Success: true,
			Message: "alive",
		}

		response.JSON(w)
	})

	// API routes
	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			// routes for user
			r.Route("/user", func(r chi.Router) {
				r.Post("/", userCommandController.CreateUser)
				r.Delete("/{id}", userCommandController.DeleteUserByID)
				r.Patch("/{id}", userCommandController.UpdateUserByID)
				r.Get("/{id}", userQueryController.GetUserByID)
			})
			r.Get("/users/", userQueryController.GetUsers)

			r.Route("/post", func(r chi.Router) {
				r.Post("/", postCommandController.CreatePost)
				r.Delete("/{id}", postCommandController.DeletePostByID)
				r.Patch("/{id}", postCommandController.UpdatePostByID)
				r.Get("/{id}", postQueryController.GetPostByID)
			})
			r.Get("/posts/", postQueryController.GetPosts)
		})
	})

	return r
}

func (router *router) Serve(port int) {
	log.Printf("[SERVER] REST server running on :%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router.InitRouter())
	if err != nil {
		log.Fatalf("[SERVER] REST server failed %v", err)
	}
}

func registerHandlers() {}

// ChiRouter export instantiated chi router once
func ChiRouter() ChiRouterInterface {
	if m == nil {
		routerOnce.Do(func() {
			// register http handlers
			registerHandlers()

			m = &router{}
		})
	}
	return m
}
