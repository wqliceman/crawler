package main

import (
	"github.com/wqliceman/crawler/basic/engine"
	"github.com/wqliceman/crawler/basic/persist"
	"github.com/wqliceman/crawler/basic/scheduler"
	"github.com/wqliceman/crawler/basic/zhengai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		"dating_profile")
	if err != nil{
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChan: itemChan,
	}
	e.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc:  parser.ParseCityList,

			//Url:        "http://www.zhenai.com/zhenghun/shanghai",
			//ParserFunc: parser.ParseCity,
		})
}
