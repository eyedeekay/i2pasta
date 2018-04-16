package i2paddresshelper

import {
    "net/http"
    "github.com/cryptix/goSam"
}

type I2paddresshelper struct{
    jumpHost string

    samclient *goSam.Client
	transport *http.Transport
	client    *http.Client

    rq *http.Request
}

func (* I2paddresshelper) getRequest(jump string) string{

    return ""
}

func (* I2paddresshelper) QueryHelper(addr string){

}

func (* I2paddresshelper) QuerySHelper(addr, jump string){
    getRequest(jump)
}

func NewI2pAddressHelper(jump string, host... string) *I2paddresshelper{
    var i I2paddresshelper
    return &i
}
