package container

import (
	"user-mapping/domain/services"
	db "user-mapping/infrastructure"
	"user-mapping/infrastructure/repository"
	"user-mapping/internal/config"
)

type ServiceContainer struct {
	LoginService *services.LoginServiceStruct
	UserService  *services.UserServiceStruct
}

func InitializeContainers(cfg *config.AppConfig) (*ServiceContainer, error) {

	// 2. INFRASTRUCTURE CORE (SINGLETON)
	dbManager := db.GetDBManager(&cfg.DBConfiguration)

	// 3. WRAPPER (public API for repos)
	sqlWrapper := db.NewSQLWrapper(dbManager)

	// Register Repositories and Service
	loginRepo := repository.NewLoginRepository(sqlWrapper)
	loginService := services.NewLoginService(loginRepo)

	userRepo := repository.NewUserRepository(sqlWrapper)
	userService := services.NewUserService(userRepo)

	serviceContainer := &ServiceContainer{
		LoginService: loginService,
		UserService:  userService,
	}

	return serviceContainer, nil
}
