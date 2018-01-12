package opass

import (
	"github.com/mdzhang/kpxcconvert/secret"
	"github.com/nbio/st"
	"os"
	"path/filepath"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func helperOpenFile(t *testing.T, name string) *os.File {
	wd, err := os.Getwd()
	check(err)
	pardir := filepath.Dir(wd)
	testdir := filepath.Join(pardir, "testdata")
	path := filepath.Join(testdir, name)
	file, err := os.Open(path)
	check(err)
	return file
}

func TestParseSecrets(t *testing.T) {
	f := helperOpenFile(t, "opass_web_login.1pif")
	secs := ParseSecrets(f, "Primary")

	esec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example Website Title",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{"example.com", "ex.com"},
	}

	st.Expect(t, 1, len(secs))
	st.Expect(t, esec, secs[0])
}

func TestParseSecretsSkipPassword(t *testing.T) {
	f := helperOpenFile(t, "opass_with_pw.1pif")
	secs := ParseSecrets(f, "Primary")

	st.Expect(t, 0, len(secs))
}

func TestParseRouter(t *testing.T) {
	f := helperOpenFile(t, "opass_router.1pif")
	secs := ParseSecrets(f, "Primary")

	esec := &secret.Secret{
		Group:    "Primary",
		Name:     "Nickname for network in 1pass",
		Username: "Wifi SSID",
		Password: "password1234",
	}

	st.Expect(t, 1, len(secs))
	st.Expect(t, esec, secs[0])
}

func TestParseSecureNote(t *testing.T) {
	f := helperOpenFile(t, "opass_secure_note.1pif")
	secs := ParseSecrets(f, "Primary")

	esec := &secret.Secret{
		Group: "Primary",
		Name:  "Note title",
		Notes: "secret note",
	}

	st.Expect(t, 1, len(secs))
	st.Expect(t, esec, secs[0])
}
