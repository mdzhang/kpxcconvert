package opass

import (
	"fmt"
	"github.com/mdzhang/kpxcconvert/secret"
	"github.com/nbio/st"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func helperFilePath(t *testing.T, name string) string {
	wd, err := os.Getwd()
	check(err)
	pardir := filepath.Dir(wd)
	testdir := filepath.Join(pardir, "testdata")
	opassdir := filepath.Join(testdir, "opass")
	path := filepath.Join(opassdir, name)
	return path
}

func helperOpenFile(t *testing.T, name string) *os.File {
	path := helperFilePath(t, name)
	file, err := os.Open(path)
	check(err)
	return file
}

func helperParseSecret(t *testing.T, name string, grp string) *secret.Secret {
	path := helperFilePath(t, fmt.Sprintf("%s.json", name))
	content, err := ioutil.ReadFile(path)
	check(err)
	sec := parseSecret(string(content), grp)
	return sec
}

func TestParseLogin(t *testing.T) {
	sec := helperParseSecret(t, "login", "Primary")

	esec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example Website Title",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{"example.com", "ex.com"},
	}

	st.Expect(t, esec, sec)
}

func TestParsePassword(t *testing.T) {
	sec := helperParseSecret(t, "password", "Primary")

	st.Expect(t, reflect.ValueOf(sec).IsNil(), true)
}

func TestParseRouter(t *testing.T) {
	sec := helperParseSecret(t, "router", "Primary")

	esec := &secret.Secret{
		Group:    "Primary",
		Name:     "Nickname for network in 1pass",
		Username: "Wifi SSID",
		Password: "password1234",
	}

	st.Expect(t, esec, sec)
}

func TestParseSecureNote(t *testing.T) {
	sec := helperParseSecret(t, "secure_note", "Primary")

	esec := &secret.Secret{
		Group: "Primary",
		Name:  "Note title",
		Notes: "secret note",
	}

	st.Expect(t, esec, sec)
}

func TestParse1Pif(t *testing.T) {
	f := helperOpenFile(t, "data.1pif")
	secs := ParseSecrets(f, "Primary")

	st.Expect(t, 3, len(secs))
}
