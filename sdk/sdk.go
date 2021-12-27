package sdk

import (
	"github.com/izern/ddns-go/dns"
	"github.com/izern/ddns-go/ip"
	"github.com/spf13/viper"
)

func init() {

}

type YunSdk interface {
	Init(value *viper.Viper) error
	UpdateRecord(parse ip.Parser, record *dns.UpdateRecord) error
}
