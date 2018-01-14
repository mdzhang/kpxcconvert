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
		Notes:    "Custom Fields: null\nNotes: \"\"\nURLs:\n- ex.com\n",
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
		Notes:    "Custom Fields: null\nNotes: \"\"\nURLs: null\n",
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
		Notes:    "Custom Fields: null\nNotes: \"\"\nURLs: null\n",
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
		Notes:    "Custom Fields: null\nNotes: some notes\nURLs: null\n",
	}

	st.Expect(t, ksec, esec)
}

func TestFromSecretHasExtras(t *testing.T) {
	sec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example Website Title",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{},
		Notes:    "some notes",
		Extras: map[string]string{
			"username": "mdzhang@example.com",
			"password": "password1234",
		},
	}

	ksec := fromSecret(sec)

	esec := &Secret{
		Group:    "Primary",
		Title:    "Example Website Title",
		Username: "mdzhang@example.com",
		Password: "password1234",
		URL:      "",
		Notes:    "Custom Fields:\n  password: password1234\n  username: mdzhang@example.com\nNotes: some notes\nURLs: null\n",
	}

	st.Expect(t, ksec, esec)
}
