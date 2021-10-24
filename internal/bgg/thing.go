package bgg

type Thing struct {
	Item Item `xml:"item"`
}

type Item struct {
	ID          int     `xml:"id,attr"`
	Type        string  `xml:"type,attr"`
	Image       string  `xml:"image"`
	Thumbnail   string  `xml:"thumbnail"`
	Names       []Value `xml:"name"`
	Description string  `xml:"description"`
	MinPlayers  IValue  `xml:"minplayers"`
	MaxPlayers  IValue  `xml:"maxplayers"`
	PlayTime    IValue  `xml:"playingtime"`
	MinPlayTime IValue  `xml:"minplaytime"`
	MaxPlayTime IValue  `xml:"maxplaytime"`
	Links       []Value `xml:"link"`
	Statistics  Stats   `xml:"statistics>ratings"`
}

type Value struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

type IValue struct {
	Type  string `xml:"type,attr"`
	Value int    `xml:"value,attr"`
}

type Stats struct {
	Rating Value `xml:"average"`
	Weight Value `xml:"averageweight"`
}

func (i *Item) Name() string {
	for _, n := range i.Names {
		if n.Type == "primary" {
			return n.Value
		}
	}
	return ""
}

func (i *Item) Categories() []string {
	c := []string{}
	for _, l := range i.Links {
		if l.Type == "boardgamecategory" {
			c = append(c, l.Value)
		}
	}
	return c
}

func (i *Item) Mechanics() []string {
	m := []string{}
	for _, l := range i.Links {
		if l.Type == "boardgamemechanic" {
			m = append(m, l.Value)
		}
	}
	return m
}
