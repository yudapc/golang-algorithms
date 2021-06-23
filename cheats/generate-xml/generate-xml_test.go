package generatexml

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/beevik/etree"
)

func TestGenerateXML(t *testing.T) {
	const jsonString = `{
        "Title": "Noval Agung",
        "URL": "https://novalagung.com",
        "Content": {
            "Articles": [{
                "Categories": [ "Server" ],
                "Title": "Connect to Oracle Server using Golang and Go-OCI8 on Ubuntu",
                "URL": "/go-oci8-oracle-linux/"
            }, {
                "Categories": [ "Server", "VPN" ],
                "Title": "Easy Setup OpenVPN Using Docker DockVPN",
                "URL": "/easy-setup-openvpn-docker/"
            }, {
                "Categories": [ "Server" ],
                "Info": "popular article",
                "Title": "Setup Ghost v0.11-LTS, Ubuntu, Nginx, Custom Domain, and SSL",
                "URL": "/ghost-v011-lts-ubuntu-nginx-custom-domain-ssl/"
            }]
        }
    }`
	data := Document{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatal(err.Error())
	}

	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	website := doc.CreateElement("website")

	website.CreateElement("title").SetText(data.Title)
	website.CreateElement("url").SetText(data.URL)

	content := website.CreateElement("contents")

	for _, each := range data.Content.Articles {
		article := content.CreateElement("article")
		article.CreateElement("title").SetText(each.Title)
		article.CreateElement("url").SetText(each.URL)

		for _, category := range each.Categories {
			article.CreateElement("category").SetText(category)
		}

		if each.Info != "" {
			article.CreateAttr("info", each.Info)
		}
	}

	doc.Indent(2)

	err = doc.WriteToFile("output.xml")
	if err != nil {
		log.Println(err.Error())
	}
}

func TestGenerateSitemapProductXML(t *testing.T) {
	const jsonString = `{
        "URL": [{
            "Loc": "https://localhost/product-slug-sku001.html",
            "Lastmod": "2021-06-22",
            "Changefreq": "monthly",
            "Priority": "0.5"
        },{
            "Loc": "https://localhost/product-slug-sku002.html",
            "Lastmod": "2021-06-22",
            "Changefreq": "monthly",
            "Priority": "0.5"
        }]
    }`
	SitemapProductGenerator(jsonString)
}
