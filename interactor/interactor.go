package interactor

import (
	"authenticator-backend/config"
	gocloak_client "authenticator-backend/infrastructure/gocloak"
	"authenticator-backend/infrastructure/gocloak/repository"
	"authenticator-backend/infrastructure/persistence/datastore"
	"authenticator-backend/presentation/http/echo/handler"
	"authenticator-backend/presentation/http/echo/middleware"
	"authenticator-backend/usecase"

	"gorm.io/gorm"
)

// Interactor
// Summary: This is interface which defines the methods that the interactor struct should implement.
type Interactor interface {
	NewAppHandler() handler.AppHandler
	NewAuthMiddleware() middleware.AuthMiddleware
}

type interactor struct {
	cfg           *config.Config
	db            *gorm.DB
	gocloakConfig *config.GocloakConfig
}

// NewInteractor
// Summary: This is function to create a new interactor struct.
// input: cfg(*config.Config) configuration
// input: db(*gorm.DB) database connection
// input: fc(*gocloak.Config) gocloak configuration
// output: Interactor
func NewInteractor(
	cfg *config.Config,
	db *gorm.DB,
	gocloakConfig *config.GocloakConfig,
) Interactor {
	return &interactor{
		cfg,
		db,
		gocloakConfig,
	}
}

// appHandler
// Summary: This is structure which defines the fields that the appHandler struct should have.
type appHandler struct {
	handler.AuthHandler
	handler.OuranosHandler
}

// NewAppHandler
// Summary: This is function to create a new appHandler struct.
// output: handler.AppHandler
func (i *interactor) NewAppHandler() handler.AppHandler {
	gocloakClient, _ := gocloak_client.NewClient(i.gocloakConfig.BaseURL, i.gocloakConfig.Realm, i.gocloakConfig.ClientID, i.gocloakConfig.ClientSecret)
	ouranosRepository := datastore.NewOuranosRepository(i.db)
	authRepository := datastore.NewAuthRepository(i.db)
	gocloakRepository := repository.NewGocloak(gocloakClient, i.gocloakConfig)
	authUsecase := usecase.NewAuthUsecase(gocloakRepository)
	verifyUsecase := usecase.NewVerifyUsecase(gocloakRepository, authRepository)
	operatorUsecase := usecase.NewOperatorUsecase(ouranosRepository)
	plantUsecase := usecase.NewPlantUsecase(ouranosRepository)
	resetUsecase := usecase.NewResetUsecase(ouranosRepository, authRepository)

	operatorHandler := handler.NewOperatorHandler(operatorUsecase)
	plantHandler := handler.NewPlantHandler(plantUsecase)
	resetHandler := handler.NewResetHandler(resetUsecase)

	authHandler := handler.NewAuthHandler(
		authUsecase,
		verifyUsecase,
	)
	ouranosHandler := handler.NewOuranosHandler(
		operatorHandler,
		plantHandler,
		resetHandler,
	)
	appHandler := &appHandler{
		AuthHandler:    authHandler,
		OuranosHandler: ouranosHandler,
	}
	return appHandler
}

// NewAuthMiddleware
// Summary: This is function to create a new authMiddleware struct.
// output: middleware.AuthMiddleware
func (i *interactor) NewAuthMiddleware() middleware.AuthMiddleware {
	gocloakClient, _ := gocloak_client.NewClient(i.gocloakConfig.BaseURL, i.gocloakConfig.Realm, i.gocloakConfig.ClientID, i.gocloakConfig.ClientSecret)
	authRepository := datastore.NewAuthRepository(i.db)
	gocloakRepository := repository.NewGocloak(gocloakClient, i.gocloakConfig)
	verifyUsecase := usecase.NewVerifyUsecase(gocloakRepository, authRepository)

	return middleware.NewAuthMiddleware(verifyUsecase)
}
