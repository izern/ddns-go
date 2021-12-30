package main

import (
	"github.com/izern/ddns-go/dns"
	"github.com/izern/ddns-go/ip"
	"github.com/izern/ddns-go/sdk"
	"github.com/izern/logging"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

var log = logging.GetLogger("ddns")

func main() {

	viper.SetEnvPrefix("ddns")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	configPath := pflag.String("config", "./config.yaml", "set config file path")
	err := viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()
	if err != nil {
		log.Fatal(err.Error())
	}
	viper.SetConfigFile(*configPath)
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	dnsV := viper.Get("dns")
	if dnsV == nil {
		log.Fatal("config error, lost dns config")
	}

	ipParserStr := viper.GetString("ip.parser")
	parser, err := ip.LoadParser(ipParserStr)
	if err != nil {
		log.Fatal("load ip parser failed. please set ip.parser")
	}

	logging.InitZapLoggerFromViper(viper.GetViper())
	log = logging.GetLogger("ddns")

	switch dnsV.(type) {
	case []map[interface{}]interface{}:
		dnsConfigs := dnsV.([]map[interface{}]interface{})
		for _, dnsConfig := range dnsConfigs {
			doUpdate(dnsConfig, parser)
		}
	case []interface{}:
		dnsArray := dnsV.([]interface{})
		for _, obj := range dnsArray {
			doUpdate(obj.(map[interface{}]interface{}), parser)
		}
	default:
		log.Fatal("dns config error, parse failed")
	}

}

func doUpdate(dnsConfig map[interface{}]interface{}, parser ip.Parser) {
	yunType := dnsConfig["yun"]
	loadSdk, err := sdk.LoadSdk(yunType.(string))
	if err != nil {
		log.Sugar().Errorf("load yun sdk for %v failed %v. skip", yunType, err)
		return
	}
	ipType := dnsConfig["type"]
	rr := dnsConfig["rr"]
	domain := dnsConfig["domain"]
	if ipType == nil || rr == nil || domain == nil {
		log.Error("dns config rr | domain | type must not be empty. skip")
		return
	}

	parseIPType, err := dns.ParseIPType(ipType.(string))
	if err != nil {
		log.Sugar().Errorf("parse ip dns.type failed %s. skip", err)
		return
	}
	err = loadSdk.Init(viper.GetViper())
	if err != nil {
		log.Sugar().Errorf("init %v sdk failed %s. skip", yunType.(string), err)
		return
	}
	err = loadSdk.UpdateRecord(parser, &dns.UpdateRecord{
		DomainName: domain.(string),
		RR:         rr.(string),
		Type:       parseIPType,
	})
	if err != nil {
		log.Sugar().Errorf(" %v sdk update record failed %s. skip", yunType.(string), err)
	}
}
