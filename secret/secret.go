package secret

// Secret represents any entry in a password manager
type Secret struct {
	Group    string
	Name     string
	Username string
	Password string
	Notes    string
	Extras   map[string]string
	Urls     []string
}
