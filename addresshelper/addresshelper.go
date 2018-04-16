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
    return &i
}
