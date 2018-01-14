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

func helperParseSecret(t *testing.T, name string, grp string) (*secret.Secret, error) {
	path := helperFilePath(t, fmt.Sprintf("%s.json", name))
	content, err := ioutil.ReadFile(path)
	check(err)
	return parseSecret(string(content), grp)
}

func TestParseLogin(t *testing.T) {
	sec, _ := helperParseSecret(t, "login", "Primary")

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

	st.Expect(t, esec, sec)
}

func TestParsePassword(t *testing.T) {
	sec, _ := helperParseSecret(t, "password", "Primary")

	st.Expect(t, reflect.ValueOf(sec).IsNil(), true)
}

func TestParseRouter(t *testing.T) {
	sec, _ := helperParseSecret(t, "router", "Primary")

	esec := &secret.Secret{
		Group:    "Primary",
		Name:     "Nickname for network in 1pass",
		Username: "Wifi SSID",
		Password: "password1234",
		Extras: map[string]string{
			"network name":              "Wifi SSID",
			"wireless security":         "wpa2p",
			"wireless network password": "password1234",
		},
	}

	st.Expect(t, esec, sec)
}

func TestParseSecureNote(t *testing.T) {
	sec, _ := helperParseSecret(t, "secure_note", "Primary")

	esec := &secret.Secret{
		Group:  "Primary",
		Name:   "Note title",
		Notes:  "secret note",
		Extras: make(map[string]string),
	}

	st.Expect(t, esec, sec)
}

func TestParseCreditCard(t *testing.T) {
	sec, _ := helperParseSecret(t, "credit_card", "Primary")

	extras := map[string]string{
		"cardholder name":     "Michelle D Zhang",
		"type":                "visa",
		"number":              "4111111111111111",
		"verification number": "999",
		"expiry date":         "203012",
		"balance":             "5.09",
	}

	esec := &secret.Secret{
		Group:    "Primary",
		Name:     "My Example Visa Credit Card",
		Username: "Michelle D Zhang",
		Password: "4111111111111111",
		Extras:   extras,
	}

	st.Expect(t, esec, sec)
}

func TestParse1Pif(t *testing.T) {
	f := helperOpenFile(t, "data.1pif")
	secs := ParseSecrets(f, "Primary")

	st.Expect(t, 3, len(secs))
}

func TestParseUnsupported(t *testing.T) {
	sec, err := helperParseSecret(t, "drivers_license", "Primary")

	st.Expect(t, reflect.ValueOf(sec).IsNil(), true)
	st.Expect(t, reflect.ValueOf(err).IsNil(), false)
}
