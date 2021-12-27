package dns

import (
	"errors"
	"fmt"
	"strings"
)

func init() {

}

type Type string

const (
	IPV4 = "A"
	IPV6 = "AAAA"
)

func (t Type) ToString() string {
	return string(t)
}

func ParseIPType(ipType string) (Type, error) {
	switch strings.ToUpper(ipType) {
	case "A", "IPV4":
		return IPV4, nil
	case "AAAA", "IPV6":
		return IPV6, nil
	default:
		return IPV4, errors.New("invalid ipType")
	}
}

type UpdateRecord struct {
	Type       Type   `json:"Type,omitempty" xml:"Type,omitempty"`
	RR         string `json:"RR,omitempty" xml:"RR,omitempty"`
	DomainName string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (record *UpdateRecord) ToString() string {
	return fmt.Sprintf("{type:%v,rr:%v,domainName:%v}", record.Type, record.RR, record.DomainName)
}
