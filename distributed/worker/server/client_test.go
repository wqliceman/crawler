package main

import (
    "fmt"
    "github.com/wqliceman/crawler/distributed/config"
    "github.com/wqliceman/crawler/distributed/rpcsupport"
    "github.com/wqliceman/crawler/distributed/worker"
    "testing"
    "time"
)

func TestCrawlService(t *testing.T){
    const host = ":9000"
    go rpcsupport.ServeRpc(host, worker.CrawlService{})
    time.Sleep(time.Second*1)

    client, err := rpcsupport.NewClient(host)
    if err != nil{
        panic(err)
    }

    var args =make(  map[string]string)
    args["Name"] =  "风中客"
    args["Gender"] = "男士"

    req := worker.Request{
        Url:    "http://album.zhenai.com/u/1237010509",
        Parser: worker.SerializedParser{
                Name: config.ParseProfile,
                Args: args,
        },
    }

    var result worker.ParseResult
    err = client.Call(
        config.CrawlServiceRpc, req, &result)
    if err != nil{
        t.Error(err)
    }else {
        fmt.Println(result)
    }
}
