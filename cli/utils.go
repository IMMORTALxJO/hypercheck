package cli

import (
	probe "probe/probe"
	re "regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

type CheckInputs []*probe.Input

// <key>[:aggregator]<operation><target>,...
func ParseArguments(input string) CheckInputs {
	inputs := []*probe.Input{}
	firstPartReg := re.MustCompile("^\\w{1,}(:\\w{1,})?")
	lastPartReg := re.MustCompile("[\\d\\w]+$")
	for id, argument := range strings.Split(input, ",") {
		log.Debugf("parse argument %d '%s'", id, argument)
		firstPart := string(firstPartReg.Find([]byte(argument)))
		log.Debugf("detected first part '%s'", firstPart)
		splittedFirstPart := strings.Split(firstPart, ":")
		key := splittedFirstPart[0]
		aggregator := ""
		if len(splittedFirstPart) > 1 {
			aggregator = splittedFirstPart[1]
		}
		lastPart := string(lastPartReg.Find([]byte(argument)))
		log.Debugf("detected list part '%s'", lastPart)
		value := ""
		operator := ""
		if len(lastPart)+len(firstPart) >= len(argument) {
			log.Debug("looks like this argument has no value or operator")
		} else {
			value = argument[len(argument)-len(lastPart):]
			operator = argument[len(firstPart) : len(argument)-len(lastPart)]
		}
		log.Debugf("key:'%s' aggregator:'%s' operator:'%s' value:'%s'", key, aggregator, operator, value)
		inputs = append(inputs, &probe.Input{
			Key:        key,
			Operator:   operator,
			Aggregator: aggregator,
			Value:      value,
		})
	}
	return inputs
}
