package main

import (
	"flag"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	RedirectorFQDN  string
	WebListenSocket string
	RedisHost       string
	Verbose         bool
)

func init() {
	flag.StringVar(&RedirectorFQDN, "fqdn", "redirector.example.org",
		"Which FQDN is used for Redirector")

	flag.StringVar(&WebListenSocket, "listen", ":8080",
		"Which socket should the web service use to bind itself")

	flag.StringVar(&RedisHost, "redis", ":6379",
		"The Redis socket that should be used")

	flag.BoolVar(&Verbose, "verbose", false,
		"Be more verbose")
}

func ValidateCommandArgs() {
}

func main() {
	flag.Parse()
	ValidateCommandArgs()

	conn := OpenConnection(RedisHost)
	defer conn.Close()

	RunRedirector(conn)
}
