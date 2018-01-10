package opass

import (
	"bufio"
	"json"
	"os"
	"strings"
)

// ParseSecrets parses secret.Secrets from a file
func ParseSecrets(f *os.File, grp string) *Secret {
	var ret []Secret

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sec := Secret{}

		// ignore lines starting with ***
		if strings.HasPrefix(line, "***") {
			continue
		}

		if err := json.Unmarshal([]byte(line), &sec); err != nil {
			panic(err)
		}

		append(ret, sec.secret(grp))
	}

	return ret
}
