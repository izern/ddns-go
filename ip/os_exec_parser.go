package ip

import (
	"bytes"
	"errors"
	"github.com/izern/ddns-go/dns"
	"github.com/izern/ddns-go/uitl"
	"os/exec"
	"strings"
)

func init() {

}

type OSExecParser struct {
	ipv4Command string
	ipv6Command string
}

func NewOSExecParser(ipv4Command, ipv6Command string) *OSExecParser {
	return &OSExecParser{
		ipv4Command: ipv4Command,
		ipv6Command: ipv6Command,
	}
}

func (parser *OSExecParser) GetIP(ipType dns.Type) (ip string, e error) {
	switch ipType {
	case dns.IPV6:
		return parser.GetIPV6()
	case dns.IPV4:
		return parser.GetIPV4()
	default:
		return "", errors.New("not support " + ipType.ToString())
	}
}

func (parser *OSExecParser) GetIPV4() (ip string, e error) {

	if parser.ipv4Command == "" {
	}
	return parser.execute(parser.ipv4Command)
}
func (parser *OSExecParser) GetIPV6() (ip string, e error) {

	if parser.ipv6Command == "" {
	}
	return parser.execute(parser.ipv6Command)
}

func (parser *OSExecParser) execute(command string) (ip string, err error) {
	cmd := exec.Command("sh", "-c", command)
	stderr := &bytes.Buffer{} // make sure to import bytes
	stdout := &bytes.Buffer{} // make sure to import bytes
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	err = cmd.Run()

	if err != nil {
		return ip, errors.New(stderr.String())
	}
	result := stdout.String()
	if result == "" {
		return ip, errors.New("parse ip result is empty, please check exec command")
	}
	result = strings.Replace(result, "\n", "", -1)
	if uitl.IsIP(result) {
		return result, nil
	}
	return ip, errors.New("exec result" + result + " invalid IP address")

}
