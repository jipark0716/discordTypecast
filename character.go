package main

import (
	"fmt"
	"regexp"
	"strconv"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type Character struct {
	Name string
	Doc *goquery.Document
	ItemLevel float64
	CharacterLevel int64
	ExpeditionLevel int64
	ClassName string
	Abilitys []string
}

func SearchCharacter(name string) (this *Character, err error) {
	this = &Character{
		Name: name,
	}
	err = this.Crow()
	return
}

func (this *Character) Crow() (err error) {
	this.loadDoc()
	this.ItemLevel = this.itemLevel()
	this.CharacterLevel = this.characterLevel()
	this.ExpeditionLevel = this.expeditionLevel()
	this.ClassName = this.className()
	this.Abilitys = this.ability()
	return
}

func (this *Character) loadDoc() (err error) {
	resp, err := http.Get(fmt.Sprintf("https://lostark.game.onstove.com/Profile/Character/%s", this.Name))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	this.Doc, err = goquery.NewDocumentFromReader(resp.Body)

	return
}

func (this *Character) itemLevel() (result float64) {
	re := regexp.MustCompile(`^.*v\.|,`)
	str := re.ReplaceAllString(this.Doc.Find(".level-info2__item").Text(), "")
	result, _ = strconv.ParseFloat(str, 64)
	return
}

func (this *Character) characterLevel() (result int64) {
	re := regexp.MustCompile(`^.*v\.`)
	str := re.ReplaceAllString(this.Doc.Find(".level-info__item").Text(), "")
	result, _ = strconv.ParseInt(str, 0, 64)
	return
}

func (this *Character) expeditionLevel() (result int64) {
	re := regexp.MustCompile(`^.*v\.`)
	str := re.ReplaceAllString(this.Doc.Find(".level-info__expedition").Text(), "")
	result, _ = strconv.ParseInt(str, 0, 64)
	return
}

func (this *Character) className() string {
	return this.Doc.Find(".profile-character-info__img").AttrOr("alt", "")
}

func (this *Character) ability() (result []string) {
	this.Doc.Find(".profile-ability-engrave .swiper-slide li").Each(func(i int, s *goquery.Selection) {
		result = append(result, s.Find("span").Text())
	})
	return
}

func (this *Character) ToDiscordMessage() (result string) {

}
