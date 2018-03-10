package xmlParserWithLoops

import (
	"fmt"
	"encoding/xml"
)

var washPostXML = []byte(`
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	 <sitemap>
			<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
	 </sitemap>
	 <sitemap>
			<loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
	 </sitemap>
	 <sitemap>
			<loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
	 </sitemap>
</sitemapindex>
`)

type Location struct {
	Loc string `xml:"loc"`
}

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`

}

func (L Location) String() string {
	return fmt.Sprintf(L.Loc);
}

func XMLParserWithLoops() {
	bytes := washPostXML
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	// range returns two values index and value
	// for i := 5 {} or for i:= range
	for _, Location :=  range s.Locations {
		fmt.Printf("\n%s\n", Location)
	}
}