package services

import (
	"fmt"
	"github.com/seancfoley/ipaddress-go/ipaddr"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestInDN11(t *testing.T) {
	testCases := []struct {
		Cidr     string
		Expected bool
	}{
		{"172.16.0.0/16", true},
		{"172.16.0.0/24", true},
		{"172.17.0.0/16", false},
		{"172.16.0.0/15", false},
	}

	for _, testCase := range testCases {
		_, ipNet, _ := net.ParseCIDR(testCase.Cidr)
		if testCase.Expected {
			assert.True(t, InDN11(*ipNet))
		} else {
			assert.False(t, InDN11(*ipNet))
		}
	}
}

func TestIpChecker(t *testing.T) {
	blockStrs := []string{
		"1.1.1.1/24", "1.1.0.2/16", "1.1.1.3/25", "1.2.0.4/16",
	}
	blocks := make([]*ipaddr.IPAddress, 0, len(blockStrs))
	for _, str := range blockStrs {
		blocks = append(blocks,
			ipaddr.NewIPAddressString(str).GetAddress().ToPrefixBlock())
	}
	trie := ipaddr.AddressTrie{}
	for _, block := range blocks {
		trie.Add(block.ToAddressBase())
	}
	fmt.Printf("trie is %v\n", trie)
	for _, block := range blocks {
		intersecting(trie, block)
	}
}
