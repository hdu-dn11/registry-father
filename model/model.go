package model

import (
	"gopkg.in/yaml.v2"
	"time"
)

type ASInfo struct {
	ASN   uint32   `yaml:"asn"`
	Owner string   `yaml:"owner"`
	IPv4  []string `yaml:"ipv4"`
	IPv6  []string `yaml:"ipv6"`

	Path      string    `yaml:"-"`
	UpdatedAt time.Time `yaml:"-"`
}

func (i ASInfo) YAML() []byte {
	out, _ := yaml.Marshal(i)
	return out
}

func (i ASInfo) Clone() ASInfo {
	clone := ASInfo{
		ASN:       i.ASN,
		Owner:     i.Owner,
		Path:      i.Path,
		UpdatedAt: i.UpdatedAt,
	}
	clone.IPv4 = append(clone.IPv4, i.IPv4...)
	clone.IPv6 = append(clone.IPv6, i.IPv6...)
	return clone
}
