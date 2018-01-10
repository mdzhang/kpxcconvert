package opass

import (
	"github.com/mdzhang/kpxcconvert/secret"
)

// Field stores data about a 1Password entry custom attribute
type Field struct {
	Value string `json:"value"`
	Name  string `json:"name"`
	ID    string `json:"id"`
	Type  string `json:"type"`
}

// URL wraps a url string
type URL struct {
	URL string `json:"url"`
}

// SecureContent stores seemingly arbitrary 1Password entry metadata
type SecureContent struct {
	URLs       URL     `json:"URLs"`
	HTMLMethod string  `json:"htmlMethod"`
	Fields     []Field `json:"fields"`
}

// Secret holds all fields exported for a given 1Password entry
// to a 1pif file
type Secret struct {
	UUID           string              `json:"uuid"`
	UpdatedAt      uint32              `json:"updatedAt"`
	LocationKey    string              `json:"locationKey"`
	SecurityLevel  string              `json:"securityLevel"`
	ContentsHash   string              `json:"contentsHash"`
	Title          string              `json:"title"`
	Location       string              `json:"location"`
	TxTimestamp    uint32              `json:"txTimestamp"`
	CreatedAt      uint32              `json:"createdAt"`
	TypeName       string              `json:"typeName"`
	SecureContents SecretSecureContent `json:"secureContents"`
}

func (sec *Secret) lookupFieldByName(name string) string {
	if sec.SecureContents == nil {
		return nil
	}

	for _, f := range sec.SecureContents.Fields {
		if f.Name == name {
			return f.Value
		}
	}

	return nil
}

func (sec *Secret) username() {
	return sec.lookupFieldByName("username")
}

func (sec *Secret) password() {
	return sec.lookupFieldByName("password")
}

func (sec *Secret) urls() {
	if sec.SecureContents == nil {
		return nil
	}

	var ret []string

	for _, u := range sec.SecureContents.URLs {
		append(ret, u.URL)
	}

	return ret
}

// secret converts an opass.Secret to a secret.Secret
func (sec *Secret) secret(grp string) {
	// TODO: parse other SecureContents.Fields to an extras map
	sec := &secret.Secret{
		Group:    grp,
		Name:     sec.Title,
		Username: sec.username(),
		Password: sec.password(),
		Urls:     sec.urls(),
	}
	return sec
}
