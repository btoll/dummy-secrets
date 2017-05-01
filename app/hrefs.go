// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "dummy-secrets": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/julianvilas/dummy-secrets/design
// --out=$(GOPATH)/src/github.com/julianvilas/dummy-secrets
// --version=v1.1.0-dirty

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