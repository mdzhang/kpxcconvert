package opass

import (
	"github.com/mdzhang/kpxcconvert/secret"
	"github.com/nbio/st"
	"reflect"
	"testing"
)

func TestSecret(t *testing.T) {
	urls := []URL{
		URL{
			URL: "example.com",
		},
		URL{
			URL: "ex.com",
		},
	}

	fields := []Field{
		Field{
			Value: "mdzhang@example.com",
			Name:  "username",
		},
		Field{
			Value: "password1234",
			Name:  "password",
		},
	}

	contents := SecureContent{
		URLs:   urls,
		Fields: fields,
	}

	osec := &Secret{
		Title:          "Example.com",
		SecureContents: contents,
	}

	esec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{"example.com", "ex.com"},
	}

	sec := osec.secret("Primary")

	st.Expect(t, sec, esec)
}

func TestSecretFallbackUsername(t *testing.T) {
	urls := []URL{
		URL{
			URL: "example.com",
		},
		URL{
			URL: "ex.com",
		},
	}

	fields := []Field{
		Field{
			Value:       "mdzhang@example.com",
			Name:        "txtUsername",
			Designation: "username",
		},
		Field{
			Value: "password1234",
			Name:  "password",
		},
	}

	contents := SecureContent{
		URLs:   urls,
		Fields: fields,
	}

	osec := &Secret{
		Title:          "Example.com",
		SecureContents: contents,
	}

	esec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{"example.com", "ex.com"},
	}

	sec := osec.secret("Primary")

	st.Expect(t, sec, esec)
}

func TestPasswordTypeSecret(t *testing.T) {
	osec := &Secret{
		Title:    "Example.com",
		TypeName: "passwords.Password",
	}

	sec := osec.secret("Primary")

	st.Expect(t, reflect.ValueOf(sec).IsNil(), true)
}

func TestSecureNote(t *testing.T) {
	contents := SecureContent{
		Notes: "super secret note content",
	}

	osec := &Secret{
		Title:          "Super secret note",
		TypeName:       "securenotes.SecureNote",
		SecureContents: contents,
	}

	esec := &secret.Secret{
		Group: "Primary",
		Name:  "Super secret note",
		Notes: "super secret note content",
	}

	sec := osec.secret("Primary")

	st.Expect(t, sec, esec)
}
