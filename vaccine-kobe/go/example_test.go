package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func baseURL() string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	URL := os.Getenv("BASE")
	return URL
}
func TestMain(t *testing.T) {
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
	fmt.Println("id", department.Department[0].ID)

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
func TestFetchArticle(t *testing.T) {
	URL := baseURL()
	f := New("article")
	article, err := FetchArticle(URL, f)
	if err != nil {
		t.Fatal(err)
	}
	if article.Articles[0].ID != 1230 {
		t.Fatal("id error")
	}
}
func TestFetchDepartment(t *testing.T) {
	URL := baseURL()
	f := New("department")
	department, err := FetchDepartment(URL, f)
	if err != nil {
		t.Fatal(err)
	}
	if department.Department[0].ID != 8769 {
		t.Fatal("erro id")
	}
}

func TestFetchAvailableDepartment(t *testing.T) {
	URL := baseURL()
	f := New("availableDepartment")
	availableDepartment, err := FetchAvailableDepartment(URL, f)
	if err != nil {
		t.Fatal(err)
	}

	if availableDepartment.DepartmentList[0] != 8769 {
		t.Fatalf("erro id : %v", availableDepartment.DepartmentList[0])
	}
}

func TestFetchItemList(t *testing.T) {
	URL := baseURL()
	f := New("itemList")
	itemList, err := FetchItemList(URL, f)
	if err != nil {
		t.Fatal(err)
	}

	if itemList.Item[0].ID != 1 {
		t.Fatalf("erro id : %v", itemList.Item[0].ID)
	}
}
