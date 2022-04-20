package main

import (
	"sync"

	"github.com/1k-ct/clonefile/kobe-COVID-19/vaccine-kobe/go/csv2json"
)

func main() {
	infectedPersonInfos := []struct {
		url      string
		fileName string
	}{
		{url: "https://www.city.kobe.lg.jp/documents/32576/kansensya.csv",
			fileName: "kansensya.csv"},
		{url: "https://www.city.kobe.lg.jp/documents/32576/kensa.csv",
			fileName: "kensa.csv"},
		{url: "https://www.city.kobe.lg.jp/documents/32576/kansensyazokusei.csv",
			fileName: "kansensyazokusei.csv"},
	}

	var wg sync.WaitGroup

	for _, v := range infectedPersonInfos {
		wg.Add(1)
		go func(url, fileName string) {
			defer wg.Done()
			csv2json.FetchKansensyaJson(url, fileName)
		}(v.url, v.fileName)
	}
	wg.Wait()
}
