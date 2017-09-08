// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "dummy-secrets": Application User Types
//
// Command:
// $ goagen
// --design=github.com/btoll/dummy-secrets/design
// --out=$(GOPATH)/src/github.com/btoll/dummy-secrets
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// secretPayload user type.
type secretPayload struct {
	// A secret to be shared with someone.
	Secret *string `form:"secret,omitempty" json:"secret,omitempty" xml:"secret,omitempty"`
}

// Validate validates the secretPayload type instance.
func (ut *secretPayload) Validate() (err error) {
	if ut.Secret == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "secret"))
	}
	if ut.Secret != nil {
		if ok := goa.ValidatePattern(`^[[:print:]]+`, *ut.Secret); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.secret`, *ut.Secret, `^[[:print:]]+`))
		}
	}
	if ut.Secret != nil {
		if utf8.RuneCountInString(*ut.Secret) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.secret`, *ut.Secret, utf8.RuneCountInString(*ut.Secret), 1, true))
		}
	}
	if ut.Secret != nil {
		if utf8.RuneCountInString(*ut.Secret) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.secret`, *ut.Secret, utf8.RuneCountInString(*ut.Secret), 255, false))
		}
	}
	return
}

// Publicize creates SecretPayload from secretPayload
func (ut *secretPayload) Publicize() *SecretPayload {
	var pub SecretPayload
	if ut.Secret != nil {
		pub.Secret = *ut.Secret
	}
	return &pub
}

// SecretPayload user type.
type SecretPayload struct {
	// A secret to be shared with someone.
	Secret string `form:"secret" json:"secret" xml:"secret"`
}

// Validate validates the SecretPayload type instance.
func (ut *SecretPayload) Validate() (err error) {
	if ut.Secret == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "secret"))
	}
	if ok := goa.ValidatePattern(`^[[:print:]]+`, ut.Secret); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.secret`, ut.Secret, `^[[:print:]]+`))
	}
	if utf8.RuneCountInString(ut.Secret) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.secret`, ut.Secret, utf8.RuneCountInString(ut.Secret), 1, true))
	}
	if utf8.RuneCountInString(ut.Secret) > 255 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.secret`, ut.Secret, utf8.RuneCountInString(ut.Secret), 255, false))
	}
	return
}
