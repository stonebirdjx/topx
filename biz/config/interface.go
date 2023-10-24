package config

import "os"

type modeClass string

var (
	yamlMode modeClass = "yaml"
)

type Parameter interface {
	Read() error
	Validate() error
}

func NewParameter() Parameter {
	switch modeClass(os.Getenv(ReadMode)) {
	case yamlMode:
		return &YamlPart{}
	default:
		return &EnvPart{}
	}
}
