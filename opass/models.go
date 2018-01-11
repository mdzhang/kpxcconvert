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
	URLs       []URL   `json:"URLs"`
	HTMLMethod string  `json:"htmlMethod"`
	Fields     []Field `json:"fields"`
}

// Secret holds all fields exported for a given 1Password entry
// to a 1pif file
type Secret struct {
	UUID           string        `json:"uuid"`
	UpdatedAt      uint32        `json:"updatedAt"`
	LocationKey    string        `json:"locationKey"`
	SecurityLevel  string        `json:"securityLevel"`
	ContentsHash   string        `json:"contentsHash"`
	Title          string        `json:"title"`
	Location       string        `json:"location"`
	TxTimestamp    uint32        `json:"txTimestamp"`
	CreatedAt      uint32        `json:"createdAt"`
	TypeName       string        `json:"typeName"`
	SecureContents SecureContent `json:"secureContents"`
}

func (osec *Secret) lookupFieldByName(name string) string {
	for _, f := range osec.SecureContents.Fields {
		if f.Name == name {
			return f.Value
		}
	}

	return ""
}

func (osec *Secret) username() string {
	return osec.lookupFieldByName("username")
}

func (osec *Secret) password() string {
	return osec.lookupFieldByName("password")
}

func (osec *Secret) urls() []string {
	var ret []string

	for _, u := range osec.SecureContents.URLs {
		ret = append(ret, u.URL)
	}

	return ret
}

// secret converts an opass.Secret to a secret.Secret
func (osec *Secret) secret(grp string) *secret.Secret {
	// TODO: parse other SecureContents.Fields to an extras map
	sec := &secret.Secret{
		Group:    grp,
		Name:     osec.Title,
		Username: osec.username(),
		Password: osec.password(),
		Urls:     osec.urls(),
	}
	return sec
}
