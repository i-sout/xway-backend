package config

type Server struct {
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                             // port
}
