// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "dummy-secrets": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/btoll/dummy-secrets/design
// --out=$(GOPATH)/src/github.com/btoll/dummy-secrets
// --version=v1.3.0

package app

import (
	"fmt"
	"strings"
)

// SecretsHref returns the resource href.
func SecretsHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/secrets/%v", paramid)
}
