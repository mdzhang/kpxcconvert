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

// Extras is for holding other field data other password
// managers may keep track of but KeePassXC does not
type Extras struct {
	URLs         []string          `json:"URLs"`
	Notes        string            `json:"Notes"`
	CustomFields map[string]string `json:"Custom Fields"`
}

// fromSecret converts between a kpxc.Secret and the more
// generic secret.Secret
func fromSecret(sec *secret.Secret) *Secret {
	url := ""
	var otherUrls []string
	var notes string

	if len(sec.Urls) == 1 {
		url = sec.Urls[0]
	} else if len(sec.Urls) > 1 {
		url = sec.Urls[0]
		otherUrls = sec.Urls[1:]
	}

	extras := Extras{
		URLs:         otherUrls,
		Notes:        sec.Notes,
		CustomFields: sec.Extras,
	}

	notesBytes, err := yaml.Marshal(extras)

	if err != nil {
		panic(err)
	}

	notes = string(notesBytes)

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
