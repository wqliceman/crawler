package parser

import (
	"github.com/wqliceman/crawler/basic/engine"
	"regexp"
)

var (
	cityRe    = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
	cityUrlRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)`)
)

func ParseCity(
	contents []byte, _ string) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: ProfileParser(
					string(m[2]), string(m[3])),
			})
	}

	return result
	// 获取其他城市链接
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}

func ProfileParser(
	name string, gender string) engine.ParserFunc {
		return func(c []byte, url string) engine.ParseResult{
			return ParseProfile(c, url, name, gender)
		}
	}