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
		Notes:    "Notes: \"\"\nURLs:\n- ex.com\n",
	}

	st.Expect(t, ksec, esec)
}

func TestFromSecretSingleUrl(t *testing.T) {
	sec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{"ex.com"},
	}

	ksec := fromSecret(sec)

	esec := &Secret{
		Group:    "Primary",
		Title:    "Example.com",
		Username: "mdzhang@example.com",
		Password: "password1234",
		URL:      "ex.com",
		Notes:    "Notes: \"\"\nURLs: null\n",
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
		Notes:    "Notes: \"\"\nURLs: null\n",
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
		Notes:    "Notes: some notes\nURLs: null\n",
	}

	st.Expect(t, ksec, esec)
}
