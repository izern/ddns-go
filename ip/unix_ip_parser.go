package ip

import (
	"errors"
	"fmt"
	"github.com/izern/ddns-go/dns"
)

func init() {

}

type UnixIpParser struct {
	device     string
	ExecParser *OSExecParser
}

func NewUnixIPParser(device string) *UnixIpParser {
	return &UnixIpParser{
		device: device,
		ExecParser: NewOSExecParser(
			fmt.Sprintf("ip -o -4 addr show %v scope global | awk '{print $4}' | cut -d/ -f1", device),
			fmt.Sprintf("ip -o -6 addr show %v scope global | awk '{print $4}' | cut -d/ -f1", device),
		),
	}
}

func (parser *UnixIpParser) GetIP(ipType dns.Type) (ip string, e error) {
	switch ipType {
	case dns.IPV6:
		return parser.GetIPV6()
	case dns.IPV4:
		return parser.GetIPV4()
	default:
		return "", errors.New("not support " + ipType.ToString())
	}
}

func (parser *UnixIpParser) GetIPV4() (ip string, e error) {
	return parser.ExecParser.GetIPV4()
}

func (parser *UnixIpParser) GetIPV6() (ip string, e error) {
	return parser.ExecParser.GetIPV6()
}
