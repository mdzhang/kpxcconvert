package kpxc

import (
	"flag"
	"github.com/mdzhang/kpxcconvert/secret"
	"github.com/nbio/st"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

func absPath(name string) string {
	wd, err := os.Getwd()
	check(err)
	pardir := filepath.Dir(wd)
	testdir := filepath.Join(pardir, "testdata")
	path := filepath.Join(testdir, name)
	return path
}

func helperOpenFile(t *testing.T, name string) *os.File {
	path := absPath(name)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	file.Truncate(0)
	check(err)
	return file
}

func TestGenerateKpxc(t *testing.T) {
	sec := &secret.Secret{
		Group:    "Primary",
		Name:     "Example Website Title",
		Username: "mdzhang@example.com",
		Password: "password1234",
		Urls:     []string{"example.com", "ex.com"},
	}

	secs := []*secret.Secret{sec}

	f := helperOpenFile(t, "kpxc.out.csv")
	GenerateKpxc(secs, f)

	golden := absPath(t.Name() + ".golden")
	silver := absPath("kpxc.out.csv")
	actual, _ := ioutil.ReadFile(silver)

	if *update {
		ioutil.WriteFile(golden, actual, 0644)
	}
	expected, _ := ioutil.ReadFile(golden)

	st.Expect(t, string(actual), string(expected))

	os.Remove(silver)
}
