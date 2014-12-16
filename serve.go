package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
)

var (
	host  string
	port  int
	cache bool
	dir   string
)

func init() {
	const (
		DEFAULT_HOST = "0.0.0.0"
		HOST_USAGE   = "host to run server on"
	)
	flag.StringVar(&host, "host", DEFAULT_HOST, HOST_USAGE)
	flag.StringVar(&host, "h", DEFAULT_HOST, HOST_USAGE)

	const (
		DEFAULT_PORT = 80
		PORT_USAGE   = "port to run server on"
	)
	flag.IntVar(&port, "port", DEFAULT_PORT, PORT_USAGE)
	flag.IntVar(&port, "p", DEFAULT_PORT, PORT_USAGE)

	const (
		DEFAULT_CACHE = false
		CACHE_USAGE   = "enable caching"
	)
	const (
		DEFAULT_DIR = "."
		DIR_USAGE   = "dir to run server on"
	)
	flag.StringVar(&dir, "dir", DEFAULT_DIR, DIR_USAGE)
	flag.StringVar(&dir, "d", DEFAULT_DIR, DIR_USAGE)

	flag.BoolVar(&cache, "enable-caching", DEFAULT_CACHE, CACHE_USAGE)
	flag.BoolVar(&cache, "c", DEFAULT_CACHE, CACHE_USAGE)
}

func main() {
	flag.Parse()

	address := host + ":" + strconv.Itoa(port)
	dirPath, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Serving '%s/' on 'http://%s'\n", dir, address)
	log.Fatal(ServeDir(dirPath, address, cache))
}
