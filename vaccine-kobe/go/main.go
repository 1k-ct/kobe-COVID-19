package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/xerrors"
)

type Article struct {
	Articles []struct {
		ID          int         `json:"id"`
		Category    interface{} `json:"category"`
		Header      string      `json:"header"`
		Description string      `json:"description"`
		URL         interface{} `json:"url"`
		Body        string      `json:"body"`
		Information struct {
			Important bool `json:"important"`
		} `json:"information"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"articles"`
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
	URL := os.Getenv("BASE")

	f := New("article")
	article, err := FetchArticle(URL, f)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Println(article.Articles[0].ID)

	f = New("department")
	department, err := FetchDepartment(URL, f)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Println(department.Department[0].Name)

	f = New("availableDepartment")
	availableDepartment, err := FetchAvailableDepartment(URL, f)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Println(availableDepartment.DepartmentList[0])

	f = New("itemList")
	itemList, err := FetchItemList(URL, f)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Println(itemList.Item[0].Name)
}

// ----------------

type Fetcher interface {
	FetchItem(url string) (interface{}, error)
}

var (
	article             = Article{}
	department          = Department{}
	availableDepartment = AvailableDepartment{}
	itemList            = ItemList{}
)

func New(contentType string) Fetcher {
	switch contentType {
	case "article":
		return article
	case "department":
		return department
	case "availableDepartment":
		return availableDepartment
	case "itemList":
		return itemList
	default:
		return nil
	}
}

// --------------------

func FetchArticle(url string, f Fetcher) (*Article, error) {
	url = url + "articles/"
	obj, err := f.FetchItem(url)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	article, ok := obj.(*Article)
	if !ok {
		return nil, xerrors.New("type failed")
	}
	return article, nil
}
func (Article) FetchItem(url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}

	item := &Article{}

	if err := json.Unmarshal(robots, item); err != nil {
		err := xerrors.New("JSON Unmarshal error:")
		return nil, xerrors.New(err.Error())
	}
	return item, nil
}

// ---------
func FetchDepartment(url string, f Fetcher) (*Department, error) {
	url = url + "department/"
	obj, err := f.FetchItem(url)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	department, ok := obj.(*Department)
	if !ok {
		return nil, xerrors.New("type failed")
	}
	return department, nil
}
func (Department) FetchItem(url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}

	item := &Department{}

	if err := json.Unmarshal(robots, item); err != nil {
		err := xerrors.New("JSON Unmarshal error:")
		return nil, xerrors.New(err.Error())
	}
	return item, nil
}
func FetchAvailableDepartment(url string, f Fetcher) (*AvailableDepartment, error) {
	url = url + "available_department/"
	obj, err := f.FetchItem(url)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	availableDepartment, ok := obj.(*AvailableDepartment)
	if !ok {
		return nil, xerrors.New("type failed")
	}
	return availableDepartment, nil
}
func (AvailableDepartment) FetchItem(url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}

	item := &AvailableDepartment{}

	if err := json.Unmarshal(robots, item); err != nil {
		err := xerrors.New("JSON Unmarshal error:")
		return nil, xerrors.New(err.Error())
	}
	return item, nil
}
func FetchItemList(url string, f Fetcher) (*ItemList, error) {
	url = url + "item/"
	obj, err := f.FetchItem(url)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	itemList, ok := obj.(*ItemList)
	if !ok {
		return nil, xerrors.New("type failed")
	}
	return itemList, nil
}
func (ItemList) FetchItem(url string) (interface{}, error) {
	res, err := http.Get(url)
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
