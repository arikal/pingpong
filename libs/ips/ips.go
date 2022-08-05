package ips

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

type IpAddress string
type IpAddresses []IpAddress

func (ips IpAddresses) GetRandom() IpAddress {
	rand.Seed(time.Now().UnixNano())

	i := rand.Intn(len(ips))

	return ips[i]
}

func CreateFromFile(f string) IpAddresses {
	b, _ := os.ReadFile(f)
	lines := strings.Split(string(b), "\n")
	ips := make(IpAddresses, len(lines))

	for i, l := range lines {
		ips[i] = IpAddress(l)
	}

	return ips
}

func CreateFromFileArg() IpAddresses {
	f := os.Args[1]

	return CreateFromFile(f)
}
