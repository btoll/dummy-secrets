package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("dummy-secrets", func() {
	Title("Dummy secrets")
	Description("Share your secrets using a dummy API")
	Version("1.0")
	BasePath("/v1")
	Scheme("http")
	Host("localhost:8080")
	Consumes("application/json")
})

var _ = Resource("Secrets", func() {
	BasePath("/secrets")

	Response(Created, func() {
		Status(201)
		Headers(func() {
			Header("Location", String, "Resource location", func() {
				Pattern("/secrets/[-a-zA-Z0-9]+")
			})
		})
	})

	Action("create", func() {
		Routing(POST("/"))
		Description("Store a new Secret")
		Payload(SecretPayload)
		Response(Created)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", UUID)
		})
		Description("Get a Secret by its ID")
		Response(OK, SecretMedia)
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
	})
})

var SecretPayload = Type("SecretPayload", func() {
	Attribute("secret", String, func() {
		Description("A secret to be shared with someone.")
		Example(`I'm a secret, share me with someone safely please.`)
		MinLength(1)
		MaxLength(255)
		Pattern("^[[:print:]]+")
	})
	Required("secret")
})

var SecretMedia = MediaType("application/vnd.secret+json", func() {
	Reference(SecretPayload)

	Attributes(func() {
		Attribute("secret")
		Required("secret")
	})

	View("default", func() {
		Attribute("secret")
	})
})
