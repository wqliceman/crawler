package client

import (
    "github.com/wqliceman/crawler/basic/engine"
    "github.com/wqliceman/crawler/distributed/config"
    "github.com/wqliceman/crawler/distributed/worker"
    "net/rpc"
)

func CreateProcessor(
    clientChan chan *rpc.Client) engine.Processor{

    return func(req engine.Request)(
            engine.ParseResult, error){
        sReq := worker.SerializeRequest(req)
        var sResult worker.ParseResult

        tempClient := <- clientChan
        err := tempClient.Call(config.CrawlServiceRpc, sReq, &sResult)
        if err != nil{
            return engine.ParseResult{}, err
        }

        return worker.DeserializeResult(sResult),nil
    }
}
