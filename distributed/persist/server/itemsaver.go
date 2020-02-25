package main

import (
    "flag"
    "fmt"
    "github.com/wqliceman/crawler/distributed/config"
    "github.com/wqliceman/crawler/distributed/persist"
    "github.com/wqliceman/crawler/distributed/rpcsupport"
    "gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0,
    "the port for item saver listen on")

func main() {
    flag.Parse()
    if *port == 0{
        fmt.Println("must specify a port for item saver")
    }
    serveRpc(
        fmt.Sprintf(":%d", *port),
        config.ElasticIndex)
}
func serveRpc(host,index string) error{
    client, err := elastic.NewClient(
        elastic.SetSniff(false))
    if err != nil{
        return err
    }

    return rpcsupport.ServeRpc( host,
        &persist.ItemSaverService{
            Client: client,
            Index:  index,
        })
}