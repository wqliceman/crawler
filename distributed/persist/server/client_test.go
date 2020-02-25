package main

import (
    "github.com/wqliceman/crawler/basic/engine"
    "github.com/wqliceman/crawler/basic/model"
    "github.com/wqliceman/crawler/distributed/rpcsupport"
    "log"
    "testing"
)

func TestItemSaver(t *testing.T) {
    const host = ":1234"
    //server
    go serveRpc(host, "test1")

    //client
    client, err := rpcsupport.NewClient(host)
    if err != nil {
        panic(err)
    }

    item := engine.Item{
        Url:  "http://album.zhenai.com/u/1544932507",
        Type: "zhenai",
        Id:   "1544932507",
        Payload: model.Profile{
            Age:        "32岁",
            Car:        "",
            Education:  "",
            Gender:     "女士",
            Height:     "165cm",
            House:      "",
            Hukou:      "工作地:天津宝坻区",
            Income:     "",
            Name:       "微笑人生",
            Occupation: "",
            Weight:     "",
            Xinzuo:     "魔羯座(12.22-01.19)",
        },
    }

    result := ""
    err = client.Call("ItemSaverService.Save", item, &result)
    if err != nil{
        log.Printf("Item Saver Error saving item %v : %v", item, err)
    }
}
