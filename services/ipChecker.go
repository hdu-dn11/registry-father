package services

import (
	"github.com/seancfoley/ipaddress-go/ipaddr"
	"net"
	"registry-father/model"
)

var (
	DN11BlockStr = []string{"172.16.0.0/16"}
	DN11Block    = make([]*ipaddr.IPAddress, 0, len(DN11BlockStr))
)

func init() {
	for _, s := range DN11BlockStr {
		DN11Block = append(DN11Block, ipaddr.NewIPAddressString(s).GetAddress().ToPrefixBlock())
	}
}

func InDN11(ip net.IPNet) bool {
	for _, address := range DN11Block {
		if address.Contains(ipaddr.NewIPAddressString(ip.String()).GetAddress()) {
			return true
		}
	}
	return false
}

func CheckCIDRConflict(a, b *model.ASInfo) bool {
	trie := ipaddr.AddressTrie{}
	for _, s := range a.IPv4 {
		trie.Add(ipaddr.NewIPAddressString(s).GetAddress().ToAddressBase())
	}
	for _, s := range a.IPv6 {
		trie.Add(ipaddr.NewIPAddressString(s).GetAddress().ToAddressBase())
	}

	blocks := make([]*ipaddr.IPAddress, 0, len(b.IPv4)+len(b.IPv6))
	for _, s := range b.IPv4 {
		blocks = append(blocks, ipaddr.NewIPAddressString(s).GetAddress())
	}
	for _, s := range b.IPv6 {
		blocks = append(blocks, ipaddr.NewIPAddressString(s).GetAddress())
	}

	for _, block := range blocks {
		if len(intersecting(trie, block)) != 0 {
			return true
		}
	}
	return false
}

func intersecting(trie ipaddr.AddressTrie, cidr *ipaddr.IPAddress) []*ipaddr.IPAddress {
	intersecting := make([]*ipaddr.IPAddress, 0, trie.Size())

	addr := cidr.ToAddressBase() // convert IPAddress to Address
	containingBlocks := trie.ElementsContaining(addr)
	containedBlocks := trie.ElementsContainedBy(addr)

	for block := containingBlocks.ShortestPrefixMatch(); block != nil; block = block.Next() {
		next := block.GetKey().ToIP()
		intersecting = append(intersecting, next)
		//if !next.Equal(cidr) {
		//	intersecting = append(intersecting, next)
		//}
	}
	iter := containedBlocks.Iterator()
	for block := iter.Next(); block != nil; block = iter.Next() {
		next := block.ToIP()
		intersecting = append(intersecting, next)
		//if !next.Equal(cidr) {
		//	intersecting = append(intersecting, next)
		//}
	}
	//fmt.Printf("CIDR %s intersects with %v\n", cidr, intersecting)
	return intersecting
}
