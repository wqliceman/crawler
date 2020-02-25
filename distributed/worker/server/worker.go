package main

import (
	"flag"
	"fmt"
	"github.com/wqliceman/crawler/distributed/rpcsupport"
	"github.com/wqliceman/crawler/distributed/worker"
	"log"
)

var port = flag.Int("port", 0,
	"the port for worker listen on")

func main() {
	flag.Parse()
	if *port == 0{
		fmt.Printf("must specify a port")
		return
	}
	log.Fatal(
		rpcsupport.ServeRpc(
			fmt.Sprintf(":%d", *port),
			worker.CrawlService{}))
}
