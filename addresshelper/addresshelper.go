package i2paddresshelper

import (
    "net/http"
    "github.com/cryptix/goSam"
)

type I2paddresshelper struct{
    jumpHost string

    samclient *goSam.Client
	transport *http.Transport
	client    *http.Client

    rq *http.Request
}

func (i *I2paddresshelper) getRequest(jump string) string{

    return ""
}

func (i *I2paddresshelper) QueryHelper(addr string){

}

func (i *I2paddresshelper) QuerySHelper(addr, jump string){
    i.getRequest(jump)
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
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: i.transport,
	}
    return &i
}
