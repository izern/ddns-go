package sdk

import (
	"errors"
	"sync"
)

var sdkRegistry sync.Map

func init() {

}

func Registry(name string, sdk YunSdk) {
	sdkRegistry.LoadOrStore(name, sdk)
}

func LoadSdk(name string) (YunSdk, error) {
	load, ok := sdkRegistry.Load(name)
	if !ok {
		return nil, errors.New(name + " not registry, not support yun sdk")
	}
	return load.(YunSdk), nil
}
