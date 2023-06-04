package model

import (
	"fmt"
	"testing"
)

func TestASInfo_YAML(t *testing.T) {
	info := &ASInfo{
		ASN:   1111111111,
		Owner: "DN11",
		IPv4:  []string{"11.11.11.11/24", "11.11.11.11/32"},
		IPv6:  []string{"11:11:11:11::11/128"},
	}

	fmt.Println(string(info.YAML()))
}
