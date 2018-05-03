package i2paddresshelper

import (
	"net/http"
	"strings"
    "log"
	"github.com/eyedeekay/i2pasta/nup"

	"github.com/cryptix/goSam"
)

type I2paddresshelper struct {
	jumpHost string
    samHost string
    samPort strings

	samclient *goSam.Client
	transport *http.Transport
	client    *http.Client

	rq    *http.Request
	l     i2pasta.I2plog
	aherr error
}

func (i *I2paddresshelper) Dial(network, addr string) (net.Conn, error) {
    portIdx := strings.Index(addr, ":")
	if portIdx >= 0 {
		addr = addr[:portIdx]
	}
	addr, i.aherr := i.samclientLookup(addr)
	if i.aherr != nil {
		return nil, i.aherr
	}

	id, _, i.aherr := i.samclientCreateStreamSession("")
	if i.aherr != nil {
		return nil, i.aherr
	}
    i.samclient, i.aherr = goSam.NewClient(i.samHost + ":" + i.samPort)
    if i.aherr != nil {
		return nil, i.aherr
	}
	i.aherr = newC.StreamConnect(id, addr)
	if i.aherr != nil {
		return nil, i.aherr
	}

	return newC.SamConn, nil
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
        log.Println("addresshelper.go ", jump, len(host), host[0])
        i.samHost = host[0]
        i.samPort = "7656"
		i.samclient, i.aherr = goSam.NewClient(i.samHost + ":" + i.samPort)
        i.l.Error(i.aherr, "addresshelper.go SAM client connection error")
	} else if len(host) == 2 {
        log.Println("addresshelper.go ", len(host), jump, host[0], host[1])
        i.samHost = host[0]
        i.samPort = host[1]
		i.samclient, i.aherr = goSam.NewClient(i.samHost + ":" + i.samPort)
        i.l.Error(i.aherr, "addresshelper.go SAM client connection error")
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
    i.client.Get(i.fixUrl("",jump))
	return &i
}
