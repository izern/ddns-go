package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {

}

func TestRegistry(t *testing.T) {

	Registry("name", &AliyunSDK{})
}

func TestLoadSdk(t *testing.T) {
	sdk, err := LoadSdk("aliyun")
	assert.NoError(t, err)
	assert.NotNil(t, sdk)

	sdk, err = LoadSdk("random name")
	assert.Error(t, err)
	assert.Nil(t, sdk)

}
