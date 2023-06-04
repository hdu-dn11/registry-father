package services

import (
	"github.com/stretchr/testify/assert"
	netUtil "github.com/yl2chen/cidranger/net"
	"net"
	"testing"
)

func TestIpChecker(t *testing.T) {
	_, inner, _ := net.ParseCIDR("172.16.0.0/16")
	//_ = DN11Range.Insert(cidranger.NewBasicRangerEntry(*inner))
	_, outer, _ := net.ParseCIDR("172.12.2.0/24")
	//
	//coveredNetworks, err := DN11Range.CoveredNetworks(*cidranger.AllIPv4)
	//if err != nil {
	//	return
	//}
	//t.Log(coveredNetworks)
	result := netUtil.NewNetwork(*inner).Covers(netUtil.NewNetwork(*outer))
	assert.True(t, result)

}
