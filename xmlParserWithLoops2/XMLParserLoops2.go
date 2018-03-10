package xmlParserWithLoops2


import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}
// dont need both for loops as it is a string only

type News struct {
	Titles []string `xml: "url>news>title"`
	Keywords []string `xml: "url>news>keywords"`
	Locations []string `xml: "url>loc"`
}

func XMLParserWithLoops2() {
	var s SitemapIndex
	var n News

	resp,_ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location :=  range s.Locations {
		resp,_ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		
	}
}