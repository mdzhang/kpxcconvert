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
		sec := Secret{}

		// ignore lines starting with ***
		if strings.HasPrefix(line, "***") {
			continue
		}

		if err := json.Unmarshal([]byte(line), &sec); err != nil {
			panic(err)
		}

		ret = append(ret, sec.secret(grp))
	}

	return ret
}
