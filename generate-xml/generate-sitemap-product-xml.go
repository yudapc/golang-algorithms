package generatexml

import (
	"encoding/json"
	"log"

	"github.com/beevik/etree"
)

type Product struct {
	URL []struct {
		Loc        string
		Lastmod    string
		Changefreq string
		Priority   string
	}
}

func SitemapProductGenerator(jsonString string) {
	data := Product{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatal(err.Error())
	}

	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	urlset := doc.CreateElement("urlset")
	urlset.CreateAttr("xmlns", "http://www.sitemaps.org/schemas/sitemap/0.9")

	for _, each := range data.URL {
		url := urlset.CreateElement("url")
		url.CreateElement("loc").SetText(each.Loc)
		url.CreateElement("lastmod").SetText(each.Lastmod)
		url.CreateElement("changefreq").SetText(each.Changefreq)
		url.CreateElement("priority").SetText(each.Priority)
	}

	doc.Indent(2)

	err = doc.WriteToFile("sitemap-product-1.xml")
	if err != nil {
		log.Println(err.Error())
	}
}
