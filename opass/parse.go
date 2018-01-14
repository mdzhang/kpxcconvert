package opass

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mdzhang/kpxcconvert/logger"
	"github.com/mdzhang/kpxcconvert/secret"
	"os"
	"reflect"
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

		sec, err := parseSecret(line, grp)

		if err != nil {
			logger.Warn(fmt.Sprintf("Skipping line due to error %s", err))
			continue
		}

		if !reflect.ValueOf(sec).IsNil() {
			ret = append(ret, sec)
		}
	}

	return ret
}

func parseSecret(line string, grp string) (sec *secret.Secret, err error) {
	osec := Secret{}

	if err = json.Unmarshal([]byte(line), &osec); err == nil {
		sec := osec.secret(grp)
		return sec, nil
	}

	return nil, err
}
