package container

import (
	"user-mapping/domain/services"
	"user-mapping/helper"
	sqlwrapper "user-mapping/infrastructure"
	"user-mapping/infrastructure/repository"
	"user-mapping/internal/config"
)

type ServiceContainer struct {
	LoginService *services.LoginServiceStruct
	UserService  *services.UserServiceStruct
}

// InitializeConfig sets up configuration-dependent helpers
func InitializeJwtAuth(cfg *config.AppConfig) *helper.JWT {
	return &helper.JWT{
		SecretKey:       cfg.JWT.Secret,
		Issuer:          cfg.JWT.Issuer,
		Audience:        cfg.JWT.Audience,
		ExpiresInMinute: cfg.JWT.ExpiresInMinute,
	}
}

func InitializeContainers(cfg *config.AppConfig) (*ServiceContainer, *helper.JWT, error) {
	// JWT helper
	jwtHelper := InitializeJwtAuth(cfg)

	// SQL wrapper
	sqlWrapper, err := sqlwrapper.NewSQLWrapper(cfg)
	if err != nil {
		return nil, nil, err
	}

	// Register Repositories and Service
	loginRepo := repository.NewLoginRepository(sqlWrapper)
	loginService := services.NewLoginService(jwtHelper, loginRepo)

	userRepo := repository.NewUserRepository(sqlWrapper)
	userService := services.NewUserService(jwtHelper, userRepo)

	serviceContainer := &ServiceContainer{
		LoginService: loginService,
		UserService:  userService,
	}

	return serviceContainer, jwtHelper, nil
}
