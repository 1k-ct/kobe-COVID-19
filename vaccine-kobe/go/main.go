package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/xerrors"
)

type Article struct {
	Item []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Interval    int    `json:"interval"`
		Information struct {
			Message             string `json:"message"`
			DisplayedNameKana   string `json:"displayed_name_kana"`
			VaccineManufacturer string `json:"vaccine_manufacturer"`
		} `json:"information"`
	} `json:"item"`
}
type Department struct {
	Department []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Number      string `json:"number"`
		Information struct {
			Area              string   `json:"area"`
			Text              string   `json:"text"`
			Access            string   `json:"access"`
			Address1          string   `json:"address1"`
			Address2          string   `json:"address2"`
			Address3          string   `json:"address3"`
			Homepage          string   `json:"homepage"`
			Postcode          string   `json:"postcode"`
			PhoneNumber       string   `json:"phone_number"`
			DisplayedName     string   `json:"displayed_name"`
			NearestStation    []string `json:"nearest_station"`
			DisplayedNameKana string   `json:"displayed_name_kana"`
		} `json:"information"`
		Item []int `json:"item"`
	} `json:"department"`
}
type AvailableDepartment struct {
	DepartmentList []int `json:"department_list"`
}
type ItemList struct {
	Item []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Interval    int    `json:"interval"`
		Information struct {
			Message             string `json:"message"`
			DisplayedNameKana   string `json:"displayed_name_kana"`
			VaccineManufacturer string `json:"vaccine_manufacturer"`
		} `json:"information"`
	} `json:"item"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	baseUrl := os.Getenv("BASE")
	baseUrl = baseUrl + "item/"
	item, err := FetchItem(baseUrl, "a")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(item.Item[0].Name)
}

func FetchItem(baseUrl string, contentType string) (*ItemList, error) {
	res, err := http.Get(baseUrl)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}

	item := &ItemList{}

	if err := json.Unmarshal(robots, item); err != nil {
		err := xerrors.New("JSON Unmarshal error:")
		return nil, xerrors.New(err.Error())
	}
	return item, nil
}
