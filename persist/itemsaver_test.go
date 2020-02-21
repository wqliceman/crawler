package persist

import (
    "context"
    "encoding/json"
    "github.com/wqliceman/crawler/basic/engine"
    "github.com/wqliceman/crawler/basic/model"
    "gopkg.in/olivere/elastic.v5"
    "testing"
)

func TestSave(t *testing.T){
    item := engine.Item{
        Url:     "",
        Type:    "",
        Id:      "",
        Payload: model.Profile{
            Name:       "iceman",
            Gender:     "男士",
            Age:        "33",
            Height:     "165",
            Weight:     "70",
            Income:     "",
            Marriage:   "",
            Education:  "",
            Occupation: "",
            Hukou:      "",
            Xinzuo:     "",
            House:      "",
            Car:        "",
        },
    }

    id, err := save(item)
    if err != nil{
        panic(err)
    }

    //TODO: Try to start up elastic search
    // here using docker go client
    client, err := elastic.NewClient(
        elastic.SetSniff(false))
    if err != nil{
        panic(err)
    }

    resp,err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
    if err != nil{
        panic(err)
    }

    var actual model.Profile
    err = json.Unmarshal(*resp.Source, &actual )
    if err != nil{
        panic(err)
    }

    if actual != profile{
        t.Errorf("got %v; expected: %v", actual, profile)
    }

}
