package config

const(
    // Parser names
    ParseCity       = "ParseCity"
    ParseCityList   = "ParseCityList"
    ParseProfile    = "ParseProfile"
    NilParser       = "NilParser"

    // ElasticSearch Index
    ElasticIndex    = "dating_profile"

    // RPC Endpoints
    ItemSaverRpc    = "ItemSaverService.Save"
    CrawlServiceRpc = "CrawlService.Process"

    // Rate limits
    Qps             = 20
)
