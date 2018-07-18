
GOPATH=$(PWD)/.go

deps:
	go get -u github.com/cryptix/gosam
	go get -u github.com/eyedeekay/i2pasta/nup

test:
	cd addresshelper && go test
	cd convert && go test
	cd httping && go test
	cd nup && go test
