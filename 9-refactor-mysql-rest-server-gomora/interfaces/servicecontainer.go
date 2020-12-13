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

	discussionRepository "rest-server/module/discussion/infrastructure/repository"
	discussionService "rest-server/module/discussion/infrastructure/service"
	discussionREST "rest-server/module/discussion/interfaces/http/rest"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	// REST
	RegisterUserRESTCommandController() userREST.UserCommandController
	RegisterUserRESTQueryController() userREST.UserQueryController
	RegisterPostRESTCommandController() discussionREST.DiscussionCommandController
	RegisterPostRESTQueryController() discussionREST.PostQueryController
	RegisterCommentRESTCommandController() discussionREST.DiscussionCommandController
	RegisterCommentRESTQueryController() discussionREST.CommentQueryController
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
// ===============================================USER===============================================
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

// ===============================================POST===============================================
// RegisterPostRESTCommandController performs dependency injection to the RegisterPostRESTCommandController
func (k *kernel) RegisterPostRESTCommandController() discussionREST.DiscussionCommandController {
	service := k.postCommandServiceContainer()

	controller := discussionREST.DiscussionCommandController{
		DiscussionCommandServiceInterface: service,
	}
	return controller
}

// RegisterPostRESTQueryController performs dependency injection to the RegisterPostRESTQueryController
func (k *kernel) RegisterPostRESTQueryController() discussionREST.PostQueryController {
	service := k.postQueryServiceContainer()

	controller := discussionREST.PostQueryController{
		DiscussionQueryServiceInterface: service,
	}

	return controller
}

// ===============================================COMMENT===============================================
// RegisterCommentRESTCommandController performs dependency injection to the RegisterCommentRESTCommandController
func (k *kernel) RegisterCommentRESTCommandController() discussionREST.DiscussionCommandController {
	service := k.commentCommandServiceContainer()

	controller := discussionREST.DiscussionCommandController{
		DiscussionCommandServiceInterface: service,
	}

	return controller
}

// RegisterCommentRESTQueryController performs dependency injection to the RegisterCommentRESTQueryController
func (k *kernel) RegisterCommentRESTQueryController() discussionREST.CommentQueryController {
	service := k.commentQueryServiceContainer()

	controller := discussionREST.CommentQueryController{
		DiscussionQueryServiceInterface: service,
	}

	return controller
}

// ===============================================USER===============================================
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

// ===============================================POST===============================================
func (k *kernel) postCommandServiceContainer() *discussionService.DiscussionCommandService {
	repository := &discussionRepository.DiscussionCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &discussionService.DiscussionCommandService{
		DiscussionCommandRepositoryInterface: &discussionRepository.DiscussionCommandRepositoryCircuitBreaker{
			DiscussionCommandRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) postQueryServiceContainer() *discussionService.DiscussionQueryService {
	repository := &discussionRepository.DiscussionQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &discussionService.DiscussionQueryService{
		DiscussionQueryRepositoryInterface: &discussionRepository.DiscussionQueryRepositoryCircuitBreaker{
			DiscussionQueryRepositoryInterface: repository,
		},
	}

	return service
}

// ===============================================COMMENT===============================================
func (k *kernel) commentCommandServiceContainer() *discussionService.DiscussionCommandService {
	repository := &discussionRepository.DiscussionCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &discussionService.DiscussionCommandService{
		DiscussionCommandRepositoryInterface: &discussionRepository.DiscussionCommandRepositoryCircuitBreaker{
			DiscussionCommandRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) commentQueryServiceContainer() *discussionService.DiscussionQueryService {
	repository := &discussionRepository.DiscussionQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &discussionService.DiscussionQueryService{
		DiscussionQueryRepositoryInterface: &discussionRepository.DiscussionQueryRepositoryCircuitBreaker{
			DiscussionQueryRepositoryInterface: repository,
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
