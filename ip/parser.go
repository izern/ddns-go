package ip

import "github.com/izern/ddns-go/dns"

func init() {

}

// Parser IP Parser
type Parser interface {
	// GetIP getIP
	GetIPV4() (ip string, e error)
	GetIPV6() (ip string, e error)
	GetIP(ipType dns.Type) (ip string, e error)
}
