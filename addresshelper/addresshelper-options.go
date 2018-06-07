package i2paddresshelper

import (
	"strconv"
)

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
