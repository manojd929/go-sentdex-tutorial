package xmlServer

/* Response XML
<?xml version="1.0" encoding="UTF-8"?>
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
*/

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml" // used to unmarshal here
)

type Location struct {
	Loc string `xml:"loc"` // in xml it is the loc tag
}

type SitemapIndex struct {
	// Locations have to be capitalized for unmarshaling in encoding/xml
	// this is a slice , an array but not of fixed size
	Locations []Location `xml:"sitemap"`

}

func (L Location) String() string {
	return fmt.Sprintf(L.Loc);
}

func XMLServer() {
	// _ is used to denote the variable is not going to be used further
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	// fmt.Println(resp)
	bytes, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(bytes)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations) // it is not a string yet
	// how did the above function get called here
}