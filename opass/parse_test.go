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
	f := helperOpenFile(t, "opass.1pif")
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
