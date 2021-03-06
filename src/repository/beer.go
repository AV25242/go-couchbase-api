package repository

import (
	"fmt"
	"utils"
)

type Beer struct {
	Id     string  `json:id`
	Abv    float64 `json:abv`
	Ibu    float64 `json:ibu`
	Name   string  `json:name`
	Srm    float64 `json:srm`
	Style  string  `json:style`
	Type   string  `json:type`
	Upc    float64 `json:upc`
	Update string  `json:update`
}

func GetBeer(id string) (Beer, error) {
	bucket := utils.Bucket()

	var beer Beer
	_, err := bucket.Get(id, &beer)
	if err != nil {
		fmt.Println(err.Error())
		return beer, err
	}
	return beer, err

}

func AddBeer(beer Beer) (Beer, error) {

	bucket := utils.Bucket()
	_, err := bucket.Insert(beer.Id, &beer, 0)
	if err != nil {
		fmt.Println(err.Error())
		return beer, err
	}
	return beer, err

}
func ModifyBeer(beer Beer) (Beer, error) {

	bucket := utils.Bucket()
	_, err := bucket.Upsert(beer.Id, &beer, 0)
	if err != nil {
		fmt.Println(err.Error())
		return beer, err
	}
	return beer, err

}

func RemoveBeer(id string) error {

	bucket := utils.Bucket()
	_, err := bucket.Remove(id, 0)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err

}
