package ip

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {

}

func Test_GetIP(t *testing.T) {

	parser := NewOSExecParser(
		fmt.Sprintf("ip -o -4 addr show %v scope global | awk '{print $4}' | cut -d/ -f1", "wlan0"),
		fmt.Sprintf("ip -o -6 addr show %v scope global | awk '{print $4}' | cut -d/ -f1", "wlan0"),
	)
	ip, err := parser.GetIPV4()
	assert.NoErrorf(t, err, "failed")
	println(ip)

	ip, err = parser.GetIPV6()
	assert.NoErrorf(t, err, "failed")
	println(ip)

}
