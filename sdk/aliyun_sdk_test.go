package sdk

import (
	"github.com/izern/ddns-go/dns"
	"github.com/izern/ddns-go/ip"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func init() {

}

func TestInit(t *testing.T) {
	viper.SetEnvPrefix("ddns")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetDefault("aliyun.accessKeySecret", "keySecret")
	viper.SetDefault("aliyun.accesskey", "key")
	sdk := &AliyunSDK{}
	err := sdk.Init(viper.GetViper())
	assert.NoError(t, err)
}

func TestUpdateRecord(t *testing.T) {
	viper.SetEnvPrefix("ddns")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	sdk := &AliyunSDK{}
	err := sdk.Init(viper.GetViper())
	assert.NoError(t, err)

	err = sdk.UpdateRecord(ip.NewUnixIPParser("wlan0"), &dns.UpdateRecord{
		Type:       dns.IPV6,
		DomainName: "izern.cn",
		RR:         "blog",
	})
	assert.NoError(t, err)
}
