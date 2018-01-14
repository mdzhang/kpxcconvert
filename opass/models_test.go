package opass

import (
	"encoding/json"
	"fmt"
	"github.com/mdzhang/kpxcconvert/secret"
	"github.com/nbio/st"
	"io/ioutil"
	"testing"
)

func helperParseOnePassSecret(t *testing.T, name string) *Secret {
	path := helperFilePath(t, fmt.Sprintf("%s.json", name))
	content, err := ioutil.ReadFile(path)
	check(err)

	osec := &Secret{}
	if err := json.Unmarshal(content, osec); err != nil {
		panic(err)
	}
	return osec
}

func TestSecret(t *testing.T) {
	osec := helperParseOnePassSecret(t, "login")

	esec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example Website Title",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{"example.com", "ex.com"},
		Extras: map[string]string{
			"username": "mdzhang@example.com",
			"password": "password1234",
		},
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
		Extras: map[string]string{
			"txtUsername": "mdzhang@example.com",
			"password":    "password1234",
		},
	}

	sec := osec.secret("Primary")

	st.Expect(t, sec, esec)
}
