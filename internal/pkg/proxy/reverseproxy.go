package proxy

type ReverseProxy struct {
	Config Config
}

func NewReverseProxy(config Config) *ReverseProxy {
	return &ReverseProxy{
		Config: config,
	}
}

func (r *ReverseProxy) ServeHttpWithPort() {

}

func (r *ReverseProxy) findHandler() {

}
