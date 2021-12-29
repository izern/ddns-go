package ip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {

}

func TestNewUnixIPParser(t *testing.T) {
	parser := NewUnixIPParser("wlan0")
	ip, err := parser.GetIPV6()
	assert.NoErrorf(t, err, "failed")
	println(ip)
	ip, err = parser.GetIPV4()
	assert.NoErrorf(t, err, "failed")
	println(ip)

}
