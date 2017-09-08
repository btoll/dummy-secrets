//go:generate goagen bootstrap -d github.com/btoll/dummy-secrets/design

package main

import (
	"github.com/btoll/dummy-secrets/app"
	"github.com/btoll/dummy-secrets/persister"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("dummy-secrets")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "Secrets" controller
	mp := persister.NewMemPersister()
	c := NewSecretsController(service, mp)
	app.MountSecretsController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
