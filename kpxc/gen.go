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
	var ksecs []Secret

	for _, s := range secs {
		k := fromSecret(s)
		append(ksecs, k)
	}

	writer := csv.NewWriter(f)
	defer writer.Flush()

	headers := []string{"Group", "Title", "Username", "Password", "URL", "Notes"}

	err := writer.write(headers)
	check(err)

	for _, k := range ksecs {
		row := []string{ksec.Group, ksec.Title, ksec.Username, ksec.Password, ksec.URL, ksec.Notes}
		err := writer.write(row)
		check(err)
	}
}
