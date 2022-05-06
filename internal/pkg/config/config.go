package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	EnableSites EnableSites `yaml:"enable_sites"`
}

type EnableSites map[string]*EnableSiteConfig

type EnableSiteConfig struct {
	Vendor   string `yaml:"vendor"`
	Endpoint string `yaml:"endpoint"`
	Bucket   string `yaml:"bucket"`
}

func Load(path string) (*Config, error) {
	if path == "" {
		path = "./configs/config.yml"
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	return cfg, yaml.Unmarshal(buf, cfg)
}
