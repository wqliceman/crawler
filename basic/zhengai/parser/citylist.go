package parser

import (
	"fmt"
	"github.com/wqliceman/crawler/basic/engine"
	"github.com/wqliceman/crawler/distributed/config"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 20
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCity, config.ParseCity),
			})
		limit--
		if limit == 0 {
			//break
		}
	}
	fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
