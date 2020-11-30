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

	"gomora/infrastructures/database/mysql"

	waitlistRepository "gomora/module/waitlist/infrastructure/repository"
	waitlistService "gomora/module/waitlist/infrastructure/service"
	waitlistREST "gomora/module/waitlist/interfaces/http/rest"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	// REST
	RegisterWaitlistRESTCommandController() waitlistREST.WaitlistCommandController
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
// RegisterWaitlistRESTCommandController performs dependency injection to the RegisterWaitlistRESTCommandController
func (k *kernel) RegisterWaitlistRESTCommandController() waitlistREST.WaitlistCommandController {
	service := k.waitlistCommandServiceContainer()

	controller := waitlistREST.WaitlistCommandController{
		WaitlistCommandServiceInterface: service,
	}

	return controller
}

//==========================================================================

func (k *kernel) waitlistCommandServiceContainer() *waitlistService.WaitlistCommandService {
	repository := &waitlistRepository.WaitlistCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &waitlistService.WaitlistCommandService{
		WaitlistCommandRepositoryInterface: &waitlistRepository.WaitlistCommandRepositoryCircuitBreaker{
			WaitlistCommandRepositoryInterface: repository,
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
