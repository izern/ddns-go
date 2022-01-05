package sdk

import (
	"errors"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/izern/ddns-go/dns"
	"github.com/izern/ddns-go/ip"
	"github.com/izern/logging"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	Registry("aliyun", &AliyunSDK{})
}

var log = logging.GetLogger("sdk.aliyun")

type AliyunSDK struct {
	client *alidns20150109.Client
}

func (sdk *AliyunSDK) Init(value *viper.Viper) error {

	accesskey := value.GetString("yun.aliyun.accesskey")
	if accesskey == "" {
		return errors.New("aliyun.accesskey must not be empty")
	}
	accessKeySecret := value.GetString("yun.aliyun.accessKeySecret")
	if accessKeySecret == "" {
		return errors.New("aliyun.accessKeySecret must not be empty")
	}
	endpoint := value.GetString("yun.aliyun.endpoint")
	if endpoint == "" {
		endpoint = "alidns.cn-hangzhou.aliyuncs.com"
	}
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &accesskey,
		// 您的AccessKey Secret
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &endpoint,
	}
	// 访问的域名
	_result, err := alidns20150109.NewClient(config)
	if err != nil {
		return err
	}
	sdk.client = _result
	return nil

}

func (sdk *AliyunSDK) UpdateRecord(parse ip.Parser, record *dns.UpdateRecord) error {
	if parse == nil {
		return errors.New("ip.Parser must not be nil")
	}
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(record.DomainName),
		RRKeyWord:  tea.String(record.RR),
		Type:       tea.String(record.Type.ToString()),
		PageSize:   tea.Int64(500),
	}
	resp, _err := sdk.client.DescribeDomainRecords(describeDomainRecordsRequest)
	if _err != nil {
		return _err
	}

	records := resp.Body.DomainRecords
	if records == nil || len(records.Record) == 0 {
		return errors.New("not found dns record for " + record.ToString())
	}

	var recordId string
	var oldIp string

	for _, recordsRecord := range records.Record {
		if *recordsRecord.RR != record.RR {
			continue
		}
		recordId = *recordsRecord.RecordId
		oldIp = *recordsRecord.Value
	}
	if recordId == "" {
		return errors.New("not found dns record for " + record.ToString())
	}
	newIp, err := parse.GetIP(record.Type)
	if err != nil {
		return err
	}
	if newIp == oldIp {
		log.Info("the same ip. skip",
			zap.String("domain", record.DomainName),
			zap.String("RR", record.RR),
			zap.String("ip", newIp),
		)
		return nil
	}
	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: tea.String(recordId),
		Type:     tea.String(record.Type.ToString()),
		RR:       tea.String(record.RR),
		Value:    tea.String(newIp),
	}

	_, err = sdk.client.UpdateDomainRecord(updateDomainRecordRequest)
	if err != nil {
		return err
	}
	log.Info("update dns ",
		zap.String("RR", record.RR),
		zap.String("domain", record.DomainName),
		zap.String("type", record.Type.ToString()),
	)

	return nil

}
