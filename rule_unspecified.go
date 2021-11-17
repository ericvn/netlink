//go:build !linux
// +build !linux

package netlink

func RuleAdd(rule *Rule) error {
	return ErrNotImplemented
}
