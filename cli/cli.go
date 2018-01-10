package cli

import (
	"github.com/mdzhang/kpxcconvert/logger"
	"github.com/mdzhang/kpxcconvert/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

// CLI flags
var (
	nick  = kingpin.Flag("1pif", "Path to 1Password .1pif export to convert").Short('p').Required().String()
	query = kingpin.Flag("kpxc", "Path to output KeePassXC compatible CSV file").Short('k').Required().String()
)

// Run CLI parser and kpxcconvert program
func Run() {
	logger.Init()
	kingpin.Version("v" + version.Version)

	kingpin.Parse()

	logger.Info("Converting file...")
}
