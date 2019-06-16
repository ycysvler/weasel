package sections

import (
	"github.com/weasel/config"
)

type app struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
}

func (s *app) SectionName() string {
	return "app"
}

var App = &app{}

func init() {
	config.Load(App)
}
