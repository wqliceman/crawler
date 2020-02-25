package persist

import (
    "github.com/wqliceman/crawler/basic/engine"
    "github.com/wqliceman/crawler/basic/persist"
    "gopkg.in/olivere/elastic.v5"
    "log"
)

type ItemSaverService struct{
    Client *elastic.Client
    Index string
}

func (s *ItemSaverService) Save(item engine.Item,
    result *string) error{
    err := persist.Save(s.Client, s.Index, item)
    if err != nil{
        *result = "ok"
        log.Printf("Save Item error : %v, %v\n", item, err)
    }else{
        log.Printf("Save Item Saved : %v\n", item)
    }
    return err
}
