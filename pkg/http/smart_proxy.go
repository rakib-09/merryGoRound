package http

import (
	"context"
	"github.com/imthaghost/merryGoRound/pkg/proxy"
	"net"
	"net/http"
	"sync"
	"time"
)

// SmartProxyClient ...
type SmartProxyClient struct {
	MaxTimeout         time.Duration // Max Timeout
	MaxIdleConnections int           // Max Idle Connections
	once               sync.Once     // sync so we only set up 1 client
	netClient          *http.Client  // http client
}

// New ...
func (s *SmartProxyClient) New() *http.Client {
	s.once.Do(func() {
		dialer := &net.Dialer{
			Timeout:   20 * time.Second, // max dialer timeout
			KeepAlive: 30 * time.Second, // keepalive duration
		}
		// transport configuration
		var netTransport = &http.Transport{
			Proxy:        proxy.SmartProxy(),   // We can use Tor or Smart Proxy - rotating IP addresses - if nil no proxy is used
			MaxIdleConns: s.MaxIdleConnections, // max idle connections
			// Dialer
			DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
				return dialer.DialContext(ctx, network, address)
			},
			TLSHandshakeTimeout: 20 * time.Second, // transport layer security max timeout
		}
		// Client
		s.netClient = &http.Client{
			Timeout:   20 * time.Second, // round stripper timeout
			Transport: netTransport,     // how our HTTP requests are made
		}
	})

	return s.netClient
}

// NewIP ...
func (s *SmartProxyClient) NewIP() {

}
