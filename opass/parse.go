package opass

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mdzhang/kpxcconvert/logger"
	"github.com/mdzhang/kpxcconvert/secret"
	"os"
	"strings"
)

// ParseSecrets parses secret.Secrets from a file
func ParseSecrets(f *os.File, grp string) []*secret.Secret {
	var ret []*secret.Secret

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		logger.Info(fmt.Sprintf("Reading line %s", line))

		// ignore lines starting with ***
		if strings.HasPrefix(line, "***") {
			continue
		}

		sec := parseSecret(line, grp)

		if sec != nil {
			ret = append(ret, sec)
		}
	}

	return ret
}

func parseSecret(line string, grp string) *secret.Secret {
	osec := Secret{}

	if err := json.Unmarshal([]byte(line), &osec); err != nil {
		panic(err)
	}

	sec := osec.secret(grp)
	return sec
}
