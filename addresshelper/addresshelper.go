package i2paddresshelper

import (
	"github.com/eyedeekay/i2pasta/nup"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/eyedeekay/gosam"
)

type I2paddresshelper struct {
	jumpHost string
	samHost  string
	samPort  string

	samclient *goSam.Client
	transport *http.Transport
	client    *http.Client

	rq    *http.Request
	l     i2pasta.I2plog
	c     bool
	aherr error
}

func (i *I2paddresshelper) fixUrl(addr, jump string) string {
	trimmedjumphost := strings.Replace(jump, "http://", "", -1)
	trimmedjumpurl := strings.Replace(addr, "http://", "", -1)
	rval := strings.Replace(trimmedjumphost+"/jump/"+trimmedjumpurl, "//", "/", -1)
	log.Println("http://"+rval)
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

type Option func(*I2paddresshelper) error

func SetVerbose(b bool) func(*I2paddresshelper) error {
	return func(c *I2paddresshelper) error {
		c.l.Verbose = b
		return nil
	}
}

func SetAddr(s string) func(*I2paddresshelper) error {
	return func(c *I2paddresshelper) error {
		c.samHost = s
		return nil
	}
}

func SetJump(s string) func(*I2paddresshelper) error {
	return func(c *I2paddresshelper) error {
		c.jumpHost = s
		return nil
	}
}

func SetPort(s string) func(*I2paddresshelper) error {
	return func(c *I2paddresshelper) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		if port < 65536 && port > -1 {
			c.samPort = s
		} else {
			c.samPort = "7656"
		}
		return nil
	}
}

func NewI2pAddressHelper(jump, host, port string) (*I2paddresshelper, error) {
	log.Println("addresshelper.go ", jump, len(host), host[0])
	return NewI2pAddressHelperFromOptions(SetJump(jump), SetAddr(host), SetPort(port), SetVerbose(true))
}

func NewI2pAddressHelperFromOptions(opts ...func(*I2paddresshelper) error) (*I2paddresshelper, error) {
	var i I2paddresshelper
	i.samHost = "127.0.0.1"
	i.samPort = "7656"
	i.jumpHost = "http://inr.i2p"
	for _, o := range opts {
		if err := o(&i); err != nil {
			return nil, err
		}
	}
	i.samclient, i.aherr = goSam.NewClient(i.samHost + ":" + i.samPort)
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
	i.client.Get(i.fixUrl("", i.jumpHost))
	return &i, nil
}
