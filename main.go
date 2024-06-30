package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	urlChan := make(chan string)

	threadsNum, pageURLList, err := readConsolFlags()

	if err != nil {
		panic(err)
	}

	for range threadsNum {
		go worker(urlChan, &wg)
	}

	for _, url := range pageURLList {
		wg.Add(1)
		urlChan <- url
	}

	wg.Wait()

	close(urlChan)

}

func worker(urlChan chan string, wg *sync.WaitGroup) {

	// обработка URL ссылок из канала urlChan
	for url := range urlChan {

		// загрузка данных с сайта
		pageData, err := downloadByURL(url)

		if err != nil {
			log.Println(err)
			wg.Done()
			continue
		}

		// преобразование URL адреса в имя файла
		fileName := convPageURLToFileName(url)

		if err = storeFile(fileName, pageData); err != nil {
			log.Println(err)
			wg.Done()
			continue
		}

		wg.Done()
	}

}

func readConsolFlags() (threadsNum int8, pageURLList []string, err error) {

	// чтение и обработка флагов командной строки

	var threads int
	var urls string

	flag.IntVar(&threads, "threads", 2, "number of threads")
	flag.StringVar(&urls, "urls", "https://google.com,https://ya.ru,https://duckduckgo.com", "list of URLs")
	flag.Parse()

	pageURLList = strings.Split(urls, ",")
	threadsNum = int8(threads)

	return threadsNum, pageURLList, err
}

func downloadByURL(pageURL string) (pageData []byte, err error) {

	// загрузка данных из страницы по URL

	resp, err := http.Get(pageURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	pageData, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return pageData, nil
}

func convPageURLToFileName(pageURL string) (fileName string) {

	// преобразование ссылки в имя файла

	fileName = strings.ReplaceAll(pageURL, "://", "-")
	fileName = strings.ReplaceAll(fileName, ".", "_")
	return fileName + ".txt"
}

func storeFile(fileName string, pageData []byte) (err error) {

	// сохранение данных в файл

	err = ioutil.WriteFile(fileName, pageData, 0644)

	if err != nil {
		return err
	}

	fmt.Println("Created file ", fileName)
	return nil
}
