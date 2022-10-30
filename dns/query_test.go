package dns

import (
	"fmt"
	"github.com/hktalent/jaeles/libs"
	"testing"
)

func TestQueryDNS(t *testing.T) {
	opt := libs.Options{
		Concurrency: 3,
		Threads:     5,
		Verbose:     true,
		NoDB:        true,
		NoOutput:    true,
	}

	dnsRcord := libs.Dns{
		Domain: "github.com",
	}
	QueryDNS(&dnsRcord, opt)
	fmt.Println(dnsRcord)
}
