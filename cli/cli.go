package cli

import (
	"fmt"
	"github.com/mdzhang/kpxcconvert/kpxc"
	"github.com/mdzhang/kpxcconvert/logger"
	"github.com/mdzhang/kpxcconvert/opass"
	"github.com/mdzhang/kpxcconvert/version"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

// CLI flags
var (
	group = kingpin.Flag("group", "Group to place secrets in e.g. a 1Password vault name").Short('g').Required().String()
	opif  = kingpin.Flag("opif", "Path to 1Password .1pif export to convert").Short('o').Required().String()
	kpxc  = kingpin.Flag("kpxc", "Path to output KeePassXC compatible CSV file").Short('k').Required().String()
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Run CLI parser and kpxcconvert program
func Run() {
	logger.Init()
	kingpin.Version("v" + version.Version)

	kingpin.Parse()

	logger.Info("Reading 1Password file...")
	file, err := os.Open(*opif)
	check(err)
	defer file.Close()

	logger.Info("Parsing 1Password file...")
	secrets := opass.ParseSecrets(file, group)

	logger.Info("Writing KeePassXC CSV file...")
	kpxc.GenerateKpxc(secrets, *kpxc)

	logger.Info("Ran successfully")
}
