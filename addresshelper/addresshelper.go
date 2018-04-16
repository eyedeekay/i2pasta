package i2paddresshelper

import (
    "net/http"

    "github.com/eyedeekay/i2pasta/nup"

    "github.com/cryptix/goSam"
)

type I2paddresshelper struct{
    jumpHost string

    samclient *goSam.Client
	transport *http.Transport
	client    *http.Client

    rq *http.Request
    l i2pasta.I2plog
}
func (i *I2paddresshelper) fixUrl(addr string) string{
    return addr
}

func (i *I2paddresshelper) getHostname(addr, jump string) string{
    resp, err := i.client.Get(i.fixUrl(addr))
    if l.Error(err, "Sent request.") {
        resp.Body.Close()
        result, rerr = ioutil.ReadAll(resp.Body)
        if l.Error(rerr, "Read response.") {
            if location := string(resp.Header.Get("Location")); location != "" {
                contents := strings.SplitN(location, "=", 2)
                if len(contents) == 2 {
                    hostname := strings.Replace(strings.Replace(strings.Replace(contents[0], "http://", "", -1), "?i2paddresshelper", "", -1), "/", "", -1)
                    return hostname
                }
            }
        }
	}
    return ""
}

func (i *I2paddresshelper) QueryHelper(addr string){
    return i.getHostname(addr, i.jumpHost)

}

func (i *I2paddresshelper) QuerySHelper(addr, jump string)string{
    return i.getHostname(addr, jump)
}

func NewI2pAddressHelper(jump string, host... string) *I2paddresshelper{
    var i I2paddresshelper
    if len(host) == 1 {
        i.samclient = goSam.NewClient(host[0] + ":7656")
    }else if len(host) == 2 {
        i.samclient = goSam.NewClient(host[0] + ":" + host[1])
    }
    i.jumpHost = jump
    i.transport = &http.Transport{
		Dial: i.samclient.Dial,
	}
	Log("i2pdig.go Setting up Client")
	i.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) l.Error {
			return http.ErrUseLastResponse
		},
		Transport: i.transport,
	}
    return &i
}
