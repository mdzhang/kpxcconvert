package kpxc

import (
	"github.com/ghodss/yaml"
	"github.com/mdzhang/kpxcconvert/secret"
)

// Secret holds fields as exported into a KeePassXC CSV
type Secret struct {
	Group    string
	Title    string
	Username string
	Password string
	URL      string
	Notes    string
}

// fromSecret converts between a kpxc.Secret and the more
// generic secret.Secret
func fromSecret(sec *secret.Secret) *Secret {
	// TODO: write sec.Extras to yaml notes
	url := ""
	var otherUrls []string
	var notes string

	if len(sec.Urls) == 1 {
		url = sec.Urls[0]
	} else if len(sec.Urls) > 1 {
		url = sec.Urls[0]
		otherUrls = sec.Urls[1:]
	}

	if len(otherUrls) > 0 {
		notesBytes, err := yaml.Marshal(struct {
			Urls  []string
			Notes string
		}{
			Urls:  otherUrls,
			Notes: sec.Notes,
		})

		if err != nil {
			panic(err)
		}

		notes = string(notesBytes)
	} else {
		notes = ""
	}

	ksec := &Secret{
		Group:    sec.Group,
		Title:    sec.Name,
		Username: sec.Username,
		Password: sec.Password,
		URL:      url,
		Notes:    notes,
	}

	return ksec
}
