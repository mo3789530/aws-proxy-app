package proxy

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host    Host    `yaml:"host"`
	Storage Storage `yaml:"storage"`
}

type Host struct {
	Domain string `yaml:"domain"`
	HTTP   []Port `yaml:"http"`
	HTTPS  []Port `yaml:"https"`
}

type Storage struct {
	StorageMode bool   `yaml:"storage_mode"`
	Bucket      Bucket `yaml:"bucket"`
}

type Bucket struct {
	BucketName string `yaml:"bucket_name"`
}

type Port struct {
	ListenPort int `yaml:"listen_port"`
	TargetPort int `yaml:"target_port"`
}

type Parameter struct {
	Name   string `yaml:"name"`
	Config Config `yaml:"config"`
}

type Parameters *[]Parameter

func NewConfig(path string) []*Parameter {
	parameters := make([]*Parameter, 0)
	parameter := &Parameter{
		Name: "test",
		Config: Config{
			Host: Host{
				Domain: "localhost",
				HTTP:   []Port{},
				HTTPS:  []Port{},
			},
			Storage: Storage{
				StorageMode: false,
				Bucket:      Bucket{},
			},
		},
	}
	parameters = append(parameters, parameter)
	data, err := os.ReadFile(path)
	if err != nil {
		slog.Error(fmt.Sprintf("cannot read %v", path))
	}
	if err := yaml.Unmarshal(data, parameters); err != nil {
		slog.Error(fmt.Sprintf("cannot unmarshal %v", err))
	}
	return parameters
}
