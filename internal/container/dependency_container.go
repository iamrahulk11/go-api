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

func InitializeContainers(dbConnections *config.DBConnections, jwtConfig *config.JWTConfig) (*ServiceContainer, error) {
	// 3. WRAPPER (public API for repos)
	sqlWrapper := db.NewSQLWrapper(dbConnections)

	// Register Repositories and Service
	loginRepo := repository.NewLoginRepository(sqlWrapper)
	loginService := services.NewLoginService(loginRepo, jwtConfig)

	userRepo := repository.NewUserRepository(sqlWrapper)
	userService := services.NewUserService(userRepo)

	serviceContainer := &ServiceContainer{
		LoginService: loginService,
		UserService:  userService,
	}

	return serviceContainer, nil
}
