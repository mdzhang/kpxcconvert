package opass

import (
	"encoding/json"
	"fmt"
	"github.com/mdzhang/kpxcconvert/secret"
	"strconv"
)

// Field stores data about a 1Password entry custom attribute
type Field struct {
	Value       string `json:"value"`
	Designation string `json:"designation"`
	Name        string `json:"name"`
}

// SectionField stores data about a 1Password entry custom attribute
// that is grouped under a section
type SectionField struct {
	Value string `json:"v"`
	Name  string `json:"t"`
}

// Section stores data about a set of related custom attributes
type Section struct {
	Name   string         `json:"title"`
	Fields []SectionField `json:"fields"`
}

// URL wraps a url string
type URL struct {
	URL string `json:"url"`
}

// SecureContent stores seemingly arbitrary 1Password entry metadata
type SecureContent struct {
	// all
	Fields   []Field   `json:"fields"`
	Sections []Section `json:"sections"`
	Notes    string    `json:"notesPlain"`

	// web logins
	URLs       []URL  `json:"URLs"`
	HTMLMethod string `json:"htmlMethod"`

	// routers
	NetworkName      string `json:"network_name"`
	WirelessPassword string `json:"wireless_password"`

	// credit cards
	Cardholder    string `json:"cardholder"`
	Cvv           string `json:"cvv"`
	CcType        string `json:"type"`
	CcNumber      string `json:"ccnum"`
	CcExpiryYear  string `json:"expiry_mm"`
	CcExpiryMonth string `json:"expiry_yy"`
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

func (osec *Secret) lookupField(field string, value string) string {
	for _, f := range osec.SecureContents.Fields {
		if field == "name" && f.Name == value {
			return f.Value
		} else if field == "designation" && f.Designation == value {
			return f.Value
		}
	}

	return ""
}

func (osec *Secret) username() string {
	if osec.TypeName == "wallet.computer.Router" {
		return osec.SecureContents.NetworkName
	}

	if osec.TypeName == "wallet.financial.CreditCard" {
		return osec.SecureContents.Cardholder
	}

	username := osec.lookupField("name", "username")

	if username == "" {
		username = osec.lookupField("designation", "username")
	}
	return username
}

func (osec *Secret) password() string {
	if osec.TypeName == "wallet.computer.Router" {
		return osec.SecureContents.WirelessPassword
	}

	if osec.TypeName == "wallet.financial.CreditCard" {
		return osec.SecureContents.CcNumber
	}

	return osec.lookupField("name", "password")
}

func (osec *Secret) urls() []string {
	var ret []string

	for _, u := range osec.SecureContents.URLs {
		ret = append(ret, u.URL)
	}

	return ret
}

func (osec *Secret) extras() map[string]string {
	extras := make(map[string]string)

	for _, f := range osec.SecureContents.Fields {
		if f.Value != "" {
			extras[f.Name] = f.Value
		}
	}

	for _, s := range osec.SecureContents.Sections {
		for _, f := range s.Fields {
			if f.Value != "" {
				fmt.Printf("name %s", f.Name)
				fmt.Printf("value %s", f.Value)
				extras[f.Name] = f.Value
			}
		}
	}

	return extras
}

// secret converts an opass.Secret to a secret.Secret
func (osec *Secret) secret(grp string) *secret.Secret {
	if osec.TypeName == "passwords.Password" {
		return nil
	}

	sec := &secret.Secret{
		Group:    grp,
		Name:     osec.Title,
		Username: osec.username(),
		Password: osec.password(),
		Urls:     osec.urls(),
		Notes:    osec.SecureContents.Notes,
		Extras:   osec.extras(),
	}

	return sec
}

type sectionFieldV struct {
	Value json.RawMessage `json:"v"`
	Name  string          `json:"t"`
}

// UnmarshalJSON is a custom implementation to account for values of
// different types
func (sec *SectionField) UnmarshalJSON(b []byte) (err error) {
	sf, val, s, n := sectionFieldV{}, "", "", uint64(0)

	if err = json.Unmarshal(b, &sf); err != nil {
		return
	}

	if sf.Value != nil {
		if err = json.Unmarshal(sf.Value, &s); err == nil {
			val = s
		} else if err = json.Unmarshal(sf.Value, &n); err == nil {
			val = strconv.FormatUint(n, 10)
		}
	}

	*sec = SectionField{
		Value: val,
		Name:  sf.Name,
	}

	return
}
