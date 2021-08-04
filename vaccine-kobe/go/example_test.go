package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"golang.org/x/xerrors"
)

func baseURL() string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	URL := os.Getenv("BASE")
	return URL
}
func writeFile(filename string, obj interface{}) error {
	buf, err := json.Marshal(obj)
	if err != nil {
		return xerrors.New(err.Error())
	}
	if err := os.WriteFile("./../../test/"+filename, buf, 0644); err != nil {
		return xerrors.New(err.Error())
	}
	return nil
}
func TestMain(t *testing.T) {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal(err)
	// }
	// URL := os.Getenv("BASE")

	// f := New("article")
	// article, err := FetchArticle(URL, f)
	// if err != nil {
	// 	log.Fatalf("%+v\n", err)
	// }

	// if err := writeFile("article.json", article); err != nil {
	// 	t.Fatalf("%+v\n", err)
	// }

	// fmt.Println(article.Articles[0].ID)

	// f = New("department")
	// if f == nil {
	// 	t.Fatalf("%+v\n", f)
	// }
	// department, err := FetchDepartment(URL, f)
	// if err != nil {
	// 	log.Fatalf("%+v\n", err)
	// }
	// fmt.Println("id", department.Department[0].ID)

	// f = New("availableDepartment")
	// if f == nil {
	// 	t.Fatalf("%+v\n", f)
	// }
	// availableDepartment, err := FetchAvailableDepartment(URL, f)
	// if err != nil {
	// 	log.Fatalf("%+v\n", err)
	// }
	// fmt.Println(availableDepartment.DepartmentList[0])

	// f = New("itemList")
	// if f == nil {
	// 	t.Fatalf("%+v\n", f)
	// }
	// itemList, err := FetchItemList(URL, f)
	// if err != nil {
	// 	log.Fatalf("%+v\n", err)
	// }
	// fmt.Println(itemList.Item[0].Name)
}

type testArticle struct{}

func (d *testArticle) FetchItem(url string) (interface{}, error) {
	data, err := ioutil.ReadFile("./../../test/article.json")
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	item := &Article{}

	if err := json.Unmarshal(data, item); err != nil {
		err := xerrors.New("JSON Unmarshal error:")
		return nil, xerrors.New(err.Error())
	}
	return item, nil
}
func TestFetchArticle(t *testing.T) {
	f := &testArticle{}

	article, err := FetchArticle("", f)
	if err != nil {
		t.Fatal(err)
	}
	if article.Articles[0].ID != 1230 {
		t.Fatal("id error")
	}
}

type testDepartment struct{}

func (d *testDepartment) FetchItem(url string) (interface{}, error) {
	data, err := ioutil.ReadFile("./../../test/department.json")
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	item := &Department{}

	if err := json.Unmarshal(data, item); err != nil {
		err := xerrors.New("JSON Unmarshal error:")
		return nil, xerrors.New(err.Error())
	}
	return item, nil
}
func TestFetchDepartment(t *testing.T) {
	f := &testDepartment{}
	department, err := FetchDepartment("", f)
	if err != nil {
		t.Fatal(err)
	}
	if department.Department[0].ID != 8769 {
		t.Fatal("erro id")
	}
}

type testAvailableDepartment struct{}

func (d *testAvailableDepartment) FetchItem(url string) (interface{}, error) {
	data, err := ioutil.ReadFile("./../../test/availableDepartment.json")
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	item := &AvailableDepartment{}

	if err := json.Unmarshal(data, item); err != nil {
		err := xerrors.New("JSON Unmarshal error:")
		return nil, xerrors.New(err.Error())
	}
	return item, nil
}
func TestFetchAvailableDepartment(t *testing.T) {
	f := &testAvailableDepartment{}
	availableDepartment, err := FetchAvailableDepartment("", f)
	if err != nil {
		t.Fatal(err)
	}

	if availableDepartment.DepartmentList[0] != 8769 {
		t.Fatalf("erro id : %v", availableDepartment.DepartmentList[0])
	}
}

type testItemList struct{}

func (d *testItemList) FetchItem(url string) (interface{}, error) {
	data, err := ioutil.ReadFile("./../../test/itemList.json")
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	item := &ItemList{}

	if err := json.Unmarshal(data, item); err != nil {
		err := xerrors.New("JSON Unmarshal error:")
		return nil, xerrors.New(err.Error())
	}
	return item, nil
}
func TestFetchItemList(t *testing.T) {
	f := &testItemList{}
	itemList, err := FetchItemList("", f)
	if err != nil {
		t.Fatal(err)
	}

	if itemList.Item[0].ID != 1 {
		t.Fatalf("erro id : %v", itemList.Item[0].ID)
	}
}
