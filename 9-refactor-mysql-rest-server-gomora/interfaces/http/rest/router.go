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
	"os"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"api-term/interfaces"
	"api-term/interfaces/http/rest/middlewares/cors"
	"api-term/interfaces/http/rest/viewmodels"
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
	academicYearCommandController := interfaces.ServiceContainer().RegisterAcademicYearRESTCommandController()
	academicYearQueryController := interfaces.ServiceContainer().RegisterAcademicYearRESTQueryController()
	gradingPeriodQueryController := interfaces.ServiceContainer().RegisterGradingPeriodRESTQueryController()
	semesterCommandController := interfaces.ServiceContainer().RegisterSemesterRESTCommandController()
	semesterQueryController := interfaces.ServiceContainer().RegisterSemesterRESTQueryController()
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
		// set jwt verifier
		tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Route("/api", func(r chi.Router) {
			r.Route("/v1", func(r chi.Router) {
				r.Route("/academic-year", func(r chi.Router) {
					r.Post("/", academicYearCommandController.CreateAcademicYear)
					r.Delete("/{id}", academicYearCommandController.DeleteAcademicYearByID)
					r.Patch("/{id}", academicYearCommandController.UpdateAcademicYearByID)
				})
				r.Get("/academic-years", academicYearQueryController.GetAcademicYears)

				r.Get("/grading-periods", gradingPeriodQueryController.GetGradingPeriods)

				r.Route("/semester", func(r chi.Router) {
					r.Post("/", semesterCommandController.CreateSemester)
					r.Delete("/{id}", semesterCommandController.DeleteSemesterByID)
					r.Patch("/{id}", semesterCommandController.UpdateSemesterByID)
				})
				r.Get("/semesters", semesterQueryController.GetSemesters)
			})
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
