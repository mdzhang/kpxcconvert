package kpxc

import (
	"encoding/csv"
	"github.com/mdzhang/kpxcconvert/secret"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GenerateKpxc writes a KeePassXC csv to the file provided
// using the secrets provided
func GenerateKpxc(secs []*secret.Secret, f *os.File) {
	var ksecs []*Secret

	for _, s := range secs {
		k := fromSecret(s)
		ksecs = append(ksecs, k)
	}

	writer := csv.NewWriter(f)
	defer writer.Flush()

	headers := []string{"Group", "Title", "Username", "Password", "URL", "Notes"}

	err := writer.Write(headers)
	check(err)

	for _, k := range ksecs {
		row := []string{k.Group, k.Title, k.Username, k.Password, k.URL, k.Notes}
		err := writer.Write(row)
		check(err)
	}
}
