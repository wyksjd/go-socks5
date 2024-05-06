package socks5

import (
	"net"

	"golang.org/x/net/context"
)

// NameResolver is used to implement custom name resolution
type NameResolver interface {
	Resolve(ctx context.Context, name string) (context.Context, net.IP, error)
}

// DNSResolver uses the system DNS to resolve host names
type DNSResolver struct{}

func (d DNSResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	addr, err := net.ResolveIPAddr("ip6", name)
	if err != nil {
		addr, err = net.ResolveIPAddr("ip4", name)
		if err != nil {
			return ctx, nil, err
		}
	}
	return ctx, addr.IP, nil
}
