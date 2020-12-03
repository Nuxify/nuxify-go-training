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

	"api-term/infrastructures/database/mysql"
	academicYearRepository "api-term/module/academicyear/infrastructure/repository"
	academicYearService "api-term/module/academicyear/infrastructure/service"
	academicYearREST "api-term/module/academicyear/interfaces/http/rest"
	gradingPeriodRepository "api-term/module/gradingperiod/infrastructure/repository"
	gradingPeriodService "api-term/module/gradingperiod/infrastructure/service"
	gradingPeriodREST "api-term/module/gradingperiod/interfaces/http/rest"
	semesterRepository "api-term/module/semester/infrastructure/repository"
	semesterService "api-term/module/semester/infrastructure/service"
	semesterREST "api-term/module/semester/interfaces/http/rest"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	// REST
	RegisterAcademicYearRESTCommandController() academicYearREST.AcademicYearCommandController
	RegisterAcademicYearRESTQueryController() academicYearREST.AcademicYearQueryController
	RegisterGradingPeriodRESTQueryController() gradingPeriodREST.GradingPeriodQueryController
	RegisterSemesterRESTCommandController() semesterREST.SemesterCommandController
	RegisterSemesterRESTQueryController() semesterREST.SemesterQueryController
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
// RegisterAcademicYearRESTCommandController performs dependency injection to the RegisterAcademicYearRESTCommandController
func (k *kernel) RegisterAcademicYearRESTCommandController() academicYearREST.AcademicYearCommandController {
	service := k.academicYearCommandServiceContainer()

	controller := academicYearREST.AcademicYearCommandController{
		AcademicYearCommandServiceInterface: service,
	}

	return controller
}

// RegisterAcademicYearRESTQueryController performs dependency injection to the RegisterAcademicYearRESTQueryController
func (k *kernel) RegisterAcademicYearRESTQueryController() academicYearREST.AcademicYearQueryController {
	service := k.academicYearQueryServiceContainer()

	controller := academicYearREST.AcademicYearQueryController{
		AcademicYearQueryServiceInterface: service,
	}

	return controller
}

// RegisterGradingPeriodRESTQueryController performs dependency injection to the RegisterGradingPeriodRESTQueryController
func (k *kernel) RegisterGradingPeriodRESTQueryController() gradingPeriodREST.GradingPeriodQueryController {
	service := k.gradingPeriodQueryServiceContainer()

	controller := gradingPeriodREST.GradingPeriodQueryController{
		GradingPeriodQueryServiceInterface: service,
	}

	return controller
}

// RegisterSemesterRESTCommandController performs dependency injection to the RegisterSemesterRESTCommandController
func (k *kernel) RegisterSemesterRESTCommandController() semesterREST.SemesterCommandController {
	service := k.semesterCommandServiceContainer()

	controller := semesterREST.SemesterCommandController{
		SemesterCommandServiceInterface: service,
	}

	return controller
}

// RegisterSemesterRESTQueryController performs dependency injection to the RegisterSemesterRESTQueryController
func (k *kernel) RegisterSemesterRESTQueryController() semesterREST.SemesterQueryController {
	service := k.semesterQueryServiceContainer()

	controller := semesterREST.SemesterQueryController{
		SemesterQueryServiceInterface: service,
	}

	return controller
}

//==========================================================================

func (k *kernel) academicYearCommandServiceContainer() *academicYearService.AcademicYearCommandService {
	repository := &academicYearRepository.AcademicYearCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &academicYearService.AcademicYearCommandService{
		AcademicYearCommandRepositoryInterface: &academicYearRepository.AcademicYearCommandRepositoryCircuitBreaker{
			AcademicYearCommandRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) academicYearQueryServiceContainer() *academicYearService.AcademicYearQueryService {
	repository := &academicYearRepository.AcademicYearQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &academicYearService.AcademicYearQueryService{
		AcademicYearQueryRepositoryInterface: &academicYearRepository.AcademicYearQueryRepositoryCircuitBreaker{
			AcademicYearQueryRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) gradingPeriodQueryServiceContainer() *gradingPeriodService.GradingPeriodQueryService {
	repository := &gradingPeriodRepository.GradingPeriodQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &gradingPeriodService.GradingPeriodQueryService{
		GradingPeriodQueryRepositoryInterface: &gradingPeriodRepository.GradingPeriodQueryRepositoryCircuitBreaker{
			GradingPeriodQueryRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) semesterCommandServiceContainer() *semesterService.SemesterCommandService {
	repository := &semesterRepository.SemesterCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &semesterService.SemesterCommandService{
		SemesterCommandRepositoryInterface: &semesterRepository.SemesterCommandRepositoryCircuitBreaker{
			SemesterCommandRepositoryInterface: repository,
		},
	}

	return service
}

func (k *kernel) semesterQueryServiceContainer() *semesterService.SemesterQueryService {
	repository := &semesterRepository.SemesterQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &semesterService.SemesterQueryService{
		SemesterQueryRepositoryInterface: &semesterRepository.SemesterQueryRepositoryCircuitBreaker{
			SemesterQueryRepositoryInterface: repository,
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
