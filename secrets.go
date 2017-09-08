package main

import (
	"github.com/btoll/dummy-secrets/app"
	"github.com/btoll/dummy-secrets/persister"
	"github.com/goadesign/goa"
)

// SecretsController implements the Secrets resource.
type SecretsController struct {
	*goa.Controller
	storage persister.Persister
}

// NewSecretsController creates a Secrets controller.
func NewSecretsController(service *goa.Service, st persister.Persister) *SecretsController {
	return &SecretsController{
		Controller: service.NewController("SecretsController"),
		storage:    st,
	}
}

// Create runs the create action.
func (c *SecretsController) Create(ctx *app.CreateSecretsContext) error {
	secret := ctx.Payload.Secret
	id := c.storage.Store(secret)

	ctx.ResponseData.Header().Set("Location", app.SecretsHref(id.String()))

	return ctx.Created()
}

// Show runs the show action.
func (c *SecretsController) Show(ctx *app.ShowSecretsContext) error {
	id := ctx.ID
	secret, err := c.storage.Retrieve(id)
	if err != nil {
		goa.LogError(ctx, err.Error())
		return ctx.NotFound()
	}

	res := &app.Secret{secret}
	if err := res.Validate(); err != nil {
		goa.LogError(ctx, err.Error())
		return ctx.InternalServerError(goa.ErrInternal(err))

	}

	return ctx.OK(res)
}
