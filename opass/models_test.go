package opass

import (
	"github.com/mdzhang/kpxcconvert/secret"
	"github.com/nbio/st"
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
