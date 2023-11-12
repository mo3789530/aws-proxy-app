package proxy

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

type Root struct {
	Keys []Key `yaml:"keys"`
}

type Key struct {
	SID SID `yaml:"sid"`
}

type SID struct {
	Name   string `yaml:"name"`
	Config Config `yaml:"config"`
}

type Config struct {
	Host    Host    `yaml:"host"`
	Storage Storage `yaml:"storage"`
}

type Host struct {
	Domain string `yaml:"domain"`
	HTTP   []HTTP `yaml:"http"`
}

type HTTP struct {
	ListenPort int `yaml:"listen_port"`
	TargetPort int `yaml:"target_port"`
}

type Storage struct {
	StorageMode bool   `yaml:"storage_mode"`
	Bucket      Bucket `yaml:"bucket"`
}

type Bucket struct {
	BucketName string `yaml:"bucket_name"`
}

func NewConfig(path string) Root {
	var root Root
	data, err := os.ReadFile(path)
	if err != nil {
		slog.Error(fmt.Sprintf("cannot read %v", path))
	}
	if err := yaml.Unmarshal(data, &root); err != nil {
		slog.Error(fmt.Sprintf("cannot unmarshal %v", err))
	}
	return root
}
