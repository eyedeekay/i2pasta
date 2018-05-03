package i2paddresshelper

import (
	"net/http"
	"strings"

	"github.com/eyedeekay/i2pasta/nup"

	"github.com/cryptix/goSam"
)

type I2paddresshelper struct {
	jumpHost string

	samclient *goSam.Client
	transport *http.Transport
	client    *http.Client

	rq    *http.Request
	l     i2pasta.I2plog
	aherr error
}

func (i *I2paddresshelper) fixUrl(addr, jump string) string {
    rval := strings.Replace(strings.Replace(jump, "http://", "", -1) + "/jump/" + addr, "//", "/", -1)
	return "http://" + rval
}

func (i *I2paddresshelper) getHostinfo(addr, jump string) (string, string) {
	resp, err := i.client.Get(i.fixUrl(addr, jump))
	if i.l.Error(err, "Sent request.") {
		resp.Body.Close()
		if location := string(resp.Header.Get("Location")); location != "" {
			contents := strings.SplitN(location, "=", 2)
			if len(contents) == 2 {
				hostname := strings.Replace(strings.Replace(strings.Replace(contents[0], "http://", "", -1), "?i2paddresshelper", "", -1), "/", "", -1)
				b64 := contents[1]
				return hostname, b64
			}
		}
	}
	return addr, "jumperror"
}

func (i *I2paddresshelper) QueryHelper(addr string) (string, string) {
	for _, jh := range strings.SplitN(i.jumpHost, ",", -1) {
		return i.getHostinfo(addr, jh)
	}
	return addr, "jumperror"
}

func (i *I2paddresshelper) QuerySHelper(addr, jump string) (string, string) {
	return i.getHostinfo(addr, jump)
}

/* Not sure if I can do this here, but if I can, I should.
func (i *I2paddresshelper) CheckRedirect(req *http.Request, via []*http.Request) error {
	var err error
	return err
}
*/

func NewI2pAddressHelper(jump string, host ...string) *I2paddresshelper {
	var i I2paddresshelper

	if len(host) == 1 {
        i.l.Log("addresshelper.go ", len(host, host[0])
		i.samclient, i.aherr = goSam.NewClient(host[0] + ":7656")
        i.l.Fatal(i.aherr, "addresshelper.go SAM client connection error")
	} else if len(host) == 2 {
        i.l.Log("addresshelper.go ", len(host, host[0], host[1])
		i.samclient, i.aherr = goSam.NewClient(host[0] + ":" + host[1])
        i.l.Fatal(i.aherr, "addresshelper.go SAM client connection error")
	}
	i.jumpHost = jump
	i.transport = &http.Transport{
		Dial: i.samclient.Dial,
	}
	i.l.Log("addresshelper.go Setting up Client")
	i.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: i.transport,
	}
	return &i
}
