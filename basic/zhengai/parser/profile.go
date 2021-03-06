package parser

import (
    "github.com/wqliceman/crawler/basic/engine"
    "github.com/wqliceman/crawler/basic/model"
    "github.com/wqliceman/crawler/distributed/config"
    "regexp"
    "strings"
)

var purpleRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var pinkRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+)</div>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(
    contents []byte,
    url string,
    name string,
    gender string ) engine.ParseResult{
    profile := model.Profile{}
    profile.Name = name
    profile.Gender = gender
    match := purpleRe.FindAllSubmatch(contents, -1)

    if match != nil && len(match) >=7 {
        profile.Marriage = string(match[0][1])
        profile.Age =  string(match[1][1])
        profile.Xinzuo = string(match[2][1])
        profile.Height= string(match[3][1])

        index := 4
        temp := string(match[4][1])
        if strings.Contains(temp, "kg") {
            profile.Weight = string(match[index][1])
            index++
        }
        profile.Hukou = string(match[index][1]) //[4:]
        index++
        profile.Income = string(match[index][1])
        //index++
        //profile.Occupation = string(match[index][1])
        //index++
        //profile.Education = string(match[index][1])
    }

    id := extractString([]byte(url), idUrlRe)
    //pink reg
    result := engine.ParseResult{
        Items:    []engine.Item{{
            Url : url,
            Type: "zhenai",
            Id: id,
            Payload:profile,
        }},
    }

    // 取本页面内，猜你喜欢的的
    var guessRe = regexp.MustCompile(`href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)
    ms := guessRe.FindAllSubmatch(contents, -1)
    for _, m := range ms {
        result.Requests = append(result.Requests, engine.Request{
            Url:   string(m[1]),
            Parser:  engine.NewFuncParser(
                ParseCity, config.ParseCity),
        })
    }

    // 取本页面其它城市链接

    return result
}

func extractString(c []byte, r *regexp.Regexp) string {
    match := r.FindSubmatch(c)
    if match != nil && len(match) >= 2 {
        return string(match[1])
    } else {
        return ""
    }
}

type ProfileParser struct {
    userName string
    gender string
}

func (p *ProfileParser) Parse(
    contents []byte, url string) engine.ParseResult {
    return parseProfile(contents, url, p.userName,p.gender)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
    var profileArgs = make(map[string]string)
    profileArgs["Name"] = p.userName
    profileArgs["Gender"] = p.gender

    return config.ParseProfile,profileArgs
}

func NewProfileParser(name string,
    gender string) *ProfileParser {
    return &ProfileParser{
        userName: name,
        gender:   gender,
    }
}