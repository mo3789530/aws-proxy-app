package proxy

type ReverseProxy struct {
}

func NewReverseProxy() *ReverseProxy {
	return &ReverseProxy{}
}

func (r *ReverseProxy) ServeHttpWithPort() {

}

func (r *ReverseProxy) findHandler() {

}
