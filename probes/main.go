package probes

import (
	"probe/options"
	"regexp"
	"strings"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type Check struct {
	Name   string
	Option options.Option
	Tests  []string
}

func OptionsToChecks(opts string, checks map[string]Check) []Check {
	var parsedProbes []Check
	chars := regexp.MustCompile("[!<>=]+")
	for _, arg := range strings.Split(opts, ",") {
		log.Debugf("Probe.ParseOptions: arg: %s", arg)
		// mode=077 -> mode
		checkName := chars.Split(arg, -1)[0]
		log.Debugf("Probe.ParseOptions: checkName: %s", checkName)
		// get probe by name
		probeCheck, ok := checks[checkName]
		if !ok {
			log.Fatalf("Probe.ParseOptions: probe has no '%s' check", checkName)
		}

		check := probeCheck.GetCopy()
		check.Option.ParseOptions(arg)

		log.Debugf("Probe.ParseOptions: registered %s check '%v', options %s", check.Option.GetType(), check.Name, check.Option.GetOptions())
		parsedProbes = append(parsedProbes, check)
	}
	return parsedProbes
}

func (c *Check) GetCopy() Check {
	copy := Check{}
	option := c.Option.GetCopy()
	copier.Copy(&copy, &c)
	copy.Option = option
	return copy
}
