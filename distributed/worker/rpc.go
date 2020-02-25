package worker

import "github.com/wqliceman/crawler/basic/engine"

type CrawlService struct {}

func (CrawlService) Process(
    req Request, result *ParseResult) error{
    enginReq, err:= DeserializeRequest(req)
    if err != nil{
        return err
    }

    enginResult, err := engine.Worker(enginReq)
    if err != nil{
        return err
    }
    *result = SerializeResult(enginResult)
    return nil
}
