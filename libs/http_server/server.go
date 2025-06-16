package http_server

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

func New(router fasthttp.RequestHandler, cert *tls.Certificate, listenerName string) *listener {
	contract := &listener{
		name: listenerName,
	}

	var certificates []tls.Certificate
	if cert != nil {
		certificates = []tls.Certificate{*cert}
	}
	contract.router = &fasthttp.Server{
		TLSConfig: &tls.Config{
			Certificates: certificates,
		},
		Handler: router,
	}

	return contract
}

type listener struct {
	name   string
	router *fasthttp.Server
}

func (c *listener) Info() string {
	return fmt.Sprintf("%s, http(s) listener", c.name)
}

func (c *listener) Run(port int) (err error) {
	if len(c.router.TLSConfig.Certificates) > 0 {
		err = c.router.ListenAndServeTLS(fmt.Sprintf(":%d", port), "", "")
	} else {
		err = c.router.ListenAndServe(fmt.Sprintf(":%d", port))
	}

	if errors.Is(err, http.ErrServerClosed) {
		err = nil
	}

	return
}

func (c *listener) Stop() error {
	return c.router.Shutdown()
}
