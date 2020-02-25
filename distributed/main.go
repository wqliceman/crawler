package main

import (
	"flag"
	"github.com/wqliceman/crawler/basic/engine"
	"github.com/wqliceman/crawler/basic/scheduler"
	"github.com/wqliceman/crawler/basic/zhengai/parser"
	"github.com/wqliceman/crawler/distributed/config"
	itemsaver "github.com/wqliceman/crawler/distributed/persist/client"
	"github.com/wqliceman/crawler/distributed/rpcsupport"
	worker "github.com/wqliceman/crawler/distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "", "worker hosts (comma sep)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(
	    *itemSaverHost)
	if err != nil {
		panic(err)
	}


	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 70,
		ItemChan:    itemChan,
		RequestProcessor: processor,
	}
	e.Run(
		engine.Request{
			Url: "http://www.zhenai.com/zhenghun",
			Parser: engine.NewFuncParser(
				parser.ParseCityList, config.ParseCityList),

			//Url:        "http://www.zhenai.com/zhenghun/shanghai",
			//ParserFunc: parser.ParseCity,
		})
}

func createClientPool(hosts []string) chan *rpc.Client{
	var clients []*rpc.Client
	for _, h :=range hosts{
		client, err := rpcsupport.NewClient(h)
		if err == nil{
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		}else {
			log.Printf("Error connection to %s: %v",
				h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func(){
		for{
			for _, client := range clients{
				out <-client
			}
		}
	}()
	return out
}
