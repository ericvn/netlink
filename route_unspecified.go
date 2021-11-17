//go:build !linux
// +build !linux

package netlink

import "strconv"

const (
	SCOPE_UNIVERSE Scope = 0
	SCOPE_SITE     Scope = 1
	SCOPE_LINK     Scope = 2
	SCOPE_HOST     Scope = 3
	SCOPE_NOWHERE  Scope = 4
)

func (r *Route) ListFlags() []string {
	return []string{}
}

func (n *NexthopInfo) ListFlags() []string {
	return []string{}
}

func (s Scope) String() string {
	return "unknown"
}

func (p RouteProtocol) String() string {
	return strconv.Itoa(int(p))
}
