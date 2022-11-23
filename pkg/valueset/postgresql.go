package pkg

import (
	"io/ioutil"
	"net/url"

	"gopkg.in/yaml.v2"
)

const hostName = "localhost"
const serverPort = 8011

type Config struct {
	Valuesets []Valueset `yaml:"valuesets"`
}

type Valueset struct {
	TableId string  `yaml:"tableId"`
	Queries []Query `yaml:"queries"`
}

type Query struct {
	Query string `yaml:"query"`
}

func parseYAML() ([]string, error) {
	var cfg Config
	data, err := ioutil.ReadFile("testdata/cfg/queries.yaml")
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	queries := make([]string, 0)
	vs_list := cfg.Valuesets
	for _, doc := range vs_list {
		for _, qry := range doc.Queries {
			queries = append(queries, qry.Query)
		}
	}
	return queries, nil
}

func process(addr string) (result map[string]string, err error) {
	urlQuery, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	urlParams := make(map[string]interface{})
	for k, v := range urlQuery.Query() {
		urlParams[k] = v
	}
	return result, nil
}
