/*
|--------------------------------------------------------------------------
| Service Container
|--------------------------------------------------------------------------
|
| This file performs the compiled dependency injection for your middlewares,
| controllers, services, providers, repositories, etc..
|
*/
package interfaces

import (
	"log"
	"os"
	"sync"

	"rest-server/infrastructures/database/mysql"

	userRepository "rest-server/module/user/infrastructure/repository"
	userService "rest-server/module/user/infrastructure/service"
	userREST "rest-server/module/user/interfaces/http/rest"

	postRepository "rest-server/module/discussion/infrastructure/repository"
	postService "rest-server/module/discussion/infrastructure/service"
	postREST "rest-server/module/discussion/interfaces/http/rest"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	// REST
	RegisterUserRESTCommandController() userREST.UserCommandController
	RegisterUserRESTQueryController() userREST.UserQueryController
	RegisterPostRESTCommandController() postREST.PostCommandController
	RegisterPostRESTQueryController() postREST.PostQueryController
}

type kernel struct{}

var (
	m              sync.Mutex
	k              *kernel
	containerOnce  sync.Once
	mysqlDBHandler *mysql.MySQLDBHandler
)

//==========================================================================

//================================= REST ===================================
// RegisterUserRESTCommandController performs dependency injection to the RegisterUserRESTCommandController
func (k *kernel) RegisterUserRESTCommandController() userREST.UserCommandController {
	service := k.userCommandServiceContainer()

	controller := userREST.UserCommandController{
		UserCommandServiceInterface: service,
	}

	return controller
}

// RegisterUserRESTQueryController performs dependency injection to the RegisterUserRESTQueryController
func (k *kernel) RegisterUserRESTQueryController() userREST.UserQueryController {
	service := k.userQueryServiceContainer()

	controller := userREST.UserQueryController{
		UserQueryServiceInterface: service,
	}

	return controller
}

// RegisterPostRESTCommandController performs dependency injection to the RegisterPostRESTCommandController
func (k *kernel) RegisterPostRESTCommandController() postREST.PostCommandController {
	service := k.postCommandServiceContainer()

	controller := postREST.PostCommandController{
		PostCommandServiceInterface: service,
	}

	return controller
}

// RegisterPostRESTQueryController performs dependency injection to the RegisterPostRESTQueryController
func (k *kernel) RegisterPostRESTQueryController() postREST.PostQueryController {
	service := k.postQueryServiceContainer()

	controller := postREST.PostQueryController{
		PostQueryServiceInterface: service,
	}

	return controller
}

//==========================================================================

func (k *kernel) userCommandServiceContainer() *userService.UserCommandService {
	repository := &userRepository.UserCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &userService.UserCommandService{
		UserCommandRepositoryInterface: &userRepository.UserCommandRepositoryCircuitBreaker{
			UserCommandRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) userQueryServiceContainer() *userService.UserQueryService {
	repository := &userRepository.UserQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &userService.UserQueryService{
		UserQueryRepositoryInterface: &userRepository.UserQueryRepositoryCircuitBreaker{
			UserQueryRepositoryInterface: repository,
		},
	}

	return service
}

// ===================================================================================
func (k *kernel) postCommandServiceContainer() *postService.PostCommandService {
	repository := &postRepository.PostCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &postService.PostCommandService{
		PostCommandRepositoryInterface: &postRepository.PostCommandRepositoryCircuitBreaker{
			PostCommandRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) postQueryServiceContainer() *postService.PostQueryService {
	repository := &postRepository.PostQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &postService.PostQueryService{
		PostQueryRepositoryInterface: &postRepository.PostQueryRepositoryCircuitBreaker{
			PostQueryRepositoryInterface: repository,
		},
	}

	return service
}

func registerHandlers() {
	// create new mysql database connection
	mysqlDBHandler = &mysql.MySQLDBHandler{}
	err := mysqlDBHandler.Connect(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"))
	if err != nil {
		log.Fatalf("[SERVER] mysql database is not responding %v", err)
	}
}

// ServiceContainer export instantiated service container once
func ServiceContainer() ServiceContainerInterface {
	m.Lock()
	defer m.Unlock()

	if k == nil {
		containerOnce.Do(func() {
			// register container handlers
			registerHandlers()

			k = &kernel{}
		})
	}
	return k
}
