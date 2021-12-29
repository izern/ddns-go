package ip

import (
	"errors"
	"github.com/spf13/viper"
	"sync"
)

var ipParserRegistry sync.Map

func init() {
	err := Register("unixIpParser", NewUnixIpParserFromViper)
	if err != nil {
		panic(err)
	}
	err = Register("osExecParser", NewOsExecParserFromViper)
	if err != nil {
		panic(err)
	}
}

func Register(name string, fun func() Parser) error {
	_, loaded := ipParserRegistry.LoadOrStore(name, fun)
	if loaded {
		return errors.New(name + " is exist")
	}
	return nil
}

func LoadParser(name string) (Parser, error) {
	load, ok := ipParserRegistry.Load(name)
	if !ok {
		return nil, errors.New(name + " is not exist")
	}
	f := load.(func() Parser)
	return f(), nil

}

func NewUnixIpParserFromViper() Parser {
	viper.SetDefault("ip.ext.unixIpParser.device", "eth0")
	device := viper.GetString("ip.ext.unixIpParser.device")
	return NewUnixIPParser(device)
}

func NewOsExecParserFromViper() Parser {
	viper.SetDefault("ip.ext.osExecParser.ipv4Cmd", "echo 192.168.0.1")
	viper.SetDefault("ip.ext.osExecParser.ipv6Cmd", "echo fe80:::::0001")
	ipv4Cmd := viper.GetString("ip.ext.osExecParser.ipv4Cmd")
	ipv6Cmd := viper.GetString("ip.ext.osExecParser.ipv6Cmd")
	return NewOSExecParser(ipv4Cmd, ipv6Cmd)
}
