package i2pasta ///convert

import (
	"os"
    "log"
)

type I2plog struct {
    verbose bool
}

func (i *I2plog)Error(err error, inp ...string) bool {
	if i.verbose {
		for _, i := range inp {
			os.Stderr.WriteString(i)
		}
	}
	if err != nil {
		log.Println(inp)
		log.Fatal(err)
		return false
	}
	return true
}
