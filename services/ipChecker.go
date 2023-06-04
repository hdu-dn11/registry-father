package services

import (
	netUtil "github.com/yl2chen/cidranger/net"
	"net"
)

var DN11Range netUtil.Network

func init() {
	_, inner, _ := net.ParseCIDR("172.16.0.0/16")
	DN11Range = netUtil.NewNetwork(*inner)
}

func InDN11(ip net.IPNet) bool {
	return DN11Range.Covers(netUtil.NewNetwork(ip))
}
