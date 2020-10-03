package main

import (
	"net/http"
    "log"
	"fmt"
	"time"
	"github.com/PuerkitoBio/goquery"
	"path"
	"regexp"
	"encoding/json"
	"encoding/csv"
	"os"
	"strconv"
	"path/filepath"
)

func fetch_page_urls() []string {
	// Define Timeout of 30 Seconds
	client := &http.Client{
        Timeout: 30 * time.Second,
    }

	// Request URL
	url := "https://www.moneycontrol.com/india/stockpricequote"
    response, err := client.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

	// Get Content of the Page
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }

	//Select the Exact Tag Required
	var hrefs []string
	document.Find("div.alph_pagn").Each(func(c_index int, container *goquery.Selection) {
		container.Find("a").Each(func(e_index1 int, element *goquery.Selection) {
			href, exists := element.Attr("href")
			if exists {
				hrefs = append(hrefs, "https://www.moneycontrol.com"+href)
			} else {
				fmt.Println("Not Found!")
			}
		})
	})
	hrefs = hrefs[1:] // Pop first element as site returns base URL

	return hrefs
}

func fetch_share_urls(urls []string) []string {
	client := &http.Client{
        Timeout: 30 * time.Second,
    }

	var share_urls []string
	bse_url := "https://priceapi.moneycontrol.com/pricefeed/bse/equitycash/"

	for _, url := range urls {
		response, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal("Error loading HTTP response body. ", err)
		}

		document.Find("a.bl_12").Each(func(index int, element *goquery.Selection) {
				href, exists := element.Attr("href")
				if exists {
					temp := path.Base(href)
					match, _:= regexp.Match("[A-Z]+[0-9]*", []byte(temp))
					if match {
						share_urls = append(share_urls, bse_url + temp)
					}
				}
		})
	}

	return share_urls
}

func json_data(urls []string) [][]string {
	client := &http.Client{
        Timeout: 30 * time.Second,
	}
	
	var name string
	var acronym string
	var current_price string
	var price_change_per_day string
	var price_change_per_week string
	var percentage_change_per_day string
	var percentage_change_per_week string	
	var json_obj map[string] interface {}
	var data map[string]interface{}
	var record = []string{"Sno.", "Acronym", "Name", "Current Price", "Price Change Per Day", "Price Change Per Week", "Percentage Change Per Day", "Percentage Change Per Week",}
	var sheet [][]string
	var ok bool
	var count int = 1

	sheet = append(sheet, record)

	for i, url := range urls {
		fmt.Printf("\r\t\t\t")
		fmt.Printf("\rStatus: %d/%d", i+1, len(urls))
		var record []string

		response, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		if response.StatusCode == http.StatusOK {
			err = json.NewDecoder(response.Body).Decode(&json_obj)
			if err  != nil {
				log.Fatal(err)
			}
			if json_obj["code"] == "200" {
				record = append(record, strconv.Itoa(count))
				data = json_obj["data"].(map[string]interface{})
		
				acronym, ok = data["symbol"].(string)
				if !ok || (acronym == "") {acronym = "Nan"}
				record = append(record, acronym)
			
				name, ok = data["company"].(string)
				if !ok || (name == "") {name = "Nan"}
				record = append(record, name)
			
				current_price, ok = data["pricecurrent"].(string)
				if !ok || (current_price == "") {current_price = "Nan"}
				record = append(record, current_price)
			
				price_change_per_day, ok = data["pricechange"].(string)
				if !ok || (price_change_per_day == "") {price_change_per_day = "Nan"}
				record = append(record, price_change_per_day)
			
				price_change_per_week, ok = data["cl1wChange"].(string)
				if !ok || (price_change_per_week == "") {price_change_per_week = "Nan"}
				record = append(record, price_change_per_week)
			
				percentage_change_per_day, ok = data["pricepercentchange"].(string)
				if !ok || (percentage_change_per_day == "") {percentage_change_per_day = "Nan"}
				record = append(record, percentage_change_per_day)
			
				percentage_change_per_week, ok = data["cl1wPerChange"].(string)
				if !ok || (percentage_change_per_week == "") {percentage_change_per_week = "Nan"}
				record = append(record, percentage_change_per_week)
				
				sheet = append(sheet, record)
				count++
			}
			
		}
		
	}
	fmt.Printf("\r\t\t\t\r")
	return sheet
}

func write_csv(name string, sheet [][]string) string {
	name =  name[0:len(name)-len(filepath.Ext(name))]
	name = filepath.Base(name)
	filename := name + " " + time.Now().Format("02-Jan-2006 150405") + ".csv"

	file, err := os.Create(filename)
    if err != nil {log.Fatal(err)}
    defer file.Close()

    csv_writer := csv.NewWriter(file)
    defer csv_writer.Flush()

    for _, record := range sheet {
		err := csv_writer.Write(record)
		if err != nil {log.Fatal(err)}
		}

	return filename
}

func main() {
	start_time := time.Now()
	fmt.Printf("Loading... Please wait!")

	page_urls := fetch_page_urls()
	share_urls := fetch_share_urls(page_urls)
	sheet := json_data(share_urls)

	filename := write_csv(os.Args[0], sheet)
	fmt.Printf("\nExported to: \"%s\"\n", filename)

	end_time := time.Now()
	elapsed := end_time.Sub(start_time)
	fmt.Printf("Time Elapsed: %v (%d Seconds)\n", elapsed, int64(elapsed.Seconds()))
}