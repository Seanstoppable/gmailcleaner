package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

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
	if term.OlderThanDays > 0 {
		t := time.Now()
		t = t.AddDate(0, 0, -term.OlderThanDays)
		resultArr = append(resultArr, fmt.Sprintf("before:%d/%d/%d", t.Year(), t.Month(), t.Day()))
	}

	return strings.Join(resultArr, " ")
}

type config struct {
	Rules []Rule
}

func LoadRules() ([]Rule, error) {

	config := config{}

	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/.config/gmailcleaner/rules.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return config.Rules, nil

}
