package proxy

type Config struct {
	Host    Host    `yaml:"host"`
	Port    Port    `yaml:"port"`
	Storage Storage `yaml:"storage"`
}

type Host struct {
	WebApi string `yaml:"webapi"`
}

type Storage struct {
	StorageMode bool `yaml:"storage_mode"`
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

func NewConfig(path string) *Parameter {
}
