// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "dummy-secrets": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/btoll/dummy-secrets/design
// --out=$(GOPATH)/src/github.com/btoll/dummy-secrets
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
)

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Secret media type (default view)
//
// Identifier: application/vnd.secret+json; view=default
type Secret struct {
	// A secret to be shared with someone.
	Secret string `form:"secret" json:"secret" xml:"secret"`
}

// Validate validates the Secret media type instance.
func (mt *Secret) Validate() (err error) {
	if mt.Secret == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secret"))
	}
	if ok := goa.ValidatePattern(`^[[:print:]]+`, mt.Secret); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.secret`, mt.Secret, `^[[:print:]]+`))
	}
	if utf8.RuneCountInString(mt.Secret) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.secret`, mt.Secret, utf8.RuneCountInString(mt.Secret), 1, true))
	}
	if utf8.RuneCountInString(mt.Secret) > 255 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.secret`, mt.Secret, utf8.RuneCountInString(mt.Secret), 255, false))
	}
	return
}

// DecodeSecret decodes the Secret instance encoded in resp body.
func (c *Client) DecodeSecret(resp *http.Response) (*Secret, error) {
	var decoded Secret
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
