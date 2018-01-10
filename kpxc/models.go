package gen

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
func fromSecret(sec *secret.Secret) {
	// TODO: write sec.Extras to yaml notes
	var url string
	var notes string

	if len(sec.Urls) == 0 {
		url = nil
		notes = nil
	} else if len(sec.Urls) == 1 {
		url = sec.Urls[0]
		notes = nil
	} else {
		url = sec.Urls[0]
		otherUrls := sec.Urls[1:]
	}

	notes = yaml.Marshal(struct {
		Urls []string
	}{
		Urls: otherUrls,
	})

	ksec := &Secret{
		Title:    sec.Name,
		Username: sec.Username,
		Password: sec.Password,
		URL:      url,
		Notes:    notes,
	}

	return ksec
}
