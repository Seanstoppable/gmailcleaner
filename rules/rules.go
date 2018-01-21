package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type Rule struct {
	Query         SearchTerms
	Modifications Modifications
}

type Modifications struct {
	AddLabels    []string
	RemoveLabels []string
}

type SearchTerms struct {
	Labels        []string
	To            string
	From          string
	OlderThanDays int
}

func (term SearchTerms) CreateQuery() string {
	resultArr := []string{}
	for _, label := range term.Labels {
		resultArr = append(resultArr, fmt.Sprintf("in:%s", label))
	}
	if term.To != "" {
		resultArr = append(resultArr, fmt.Sprintf("to:%s", term.To))
	}
	if term.From != "" {
		resultArr = append(resultArr, fmt.Sprintf("from:%s", term.From))
	}
	//TODO date stuff

	return strings.Join(resultArr, " ")
}

type config struct {
	Rules []Rule
}

func LoadRules() ([]Rule, error) {

	config := config{}

	data, err := ioutil.ReadFile("rules.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return config.Rules, nil

}
