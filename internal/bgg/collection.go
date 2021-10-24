package bgg

type Collection struct {
	Total int              `xml:"totalitems,attr"`
	Items []CollectionItem `xml:"item"`
}

type CollectionItem struct {
	ID        int    `xml:"objectid,attr"`
	Type      string `xml:"subtype,attr"`
	Name      string `xml:"name"`
	Image     string `xml:"image"`
	Thumbnail string `xml:"thumbnail"`
}
