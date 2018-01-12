package kpxc

import (
	"github.com/mdzhang/kpxcconvert/secret"
	"github.com/nbio/st"
	"testing"
)

func TestFromSecretManyUrls(t *testing.T) {
	sec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{"example.com", "ex.com"},
	}

	ksec := fromSecret(sec)

	esec := &Secret{
		Group:    "Primary",
		Title:    "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		URL:      "example.com",
		Notes:    "Notes: \"\"\nUrls:\n- ex.com\n",
	}

	st.Expect(t, ksec, esec)
}

func TestFromSecretSingleUrl(t *testing.T) {
	sec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{"example.com"},
	}

	ksec := fromSecret(sec)

	esec := &Secret{
		Group:    "Primary",
		Title:    "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		URL:      "example.com",
		Notes:    "",
	}

	st.Expect(t, ksec, esec)
}

func TestFromSecretNoUrl(t *testing.T) {
	sec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{},
	}

	ksec := fromSecret(sec)

	esec := &Secret{
		Group:    "Primary",
		Title:    "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		URL:      "",
		Notes:    "",
	}

	st.Expect(t, ksec, esec)
}

func TestFromSecretHasNotes(t *testing.T) {
	sec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Notes:    "some notes",
		Urls:     []string{},
	}

	ksec := fromSecret(sec)

	esec := &Secret{
		Group:    "Primary",
		Title:    "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		URL:      "",
		Notes:    "",
	}

	st.Expect(t, ksec, esec)
}
