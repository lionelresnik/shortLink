package main

import (
	"log"
	"net/http"
	"smartShortLink/database"
	"time"
)

const timeToRecheck = time.Second * 60
const statusCodeOK = 200

func GarbageCollector() {
	for {
		urls, err := dbInstance.GetAllDomains()
		if err != nil {
			log.Printf("unable to get all domains")
			return
		}
		log.Printf("*************** Garbage Collector Started- will remove inactive urls ***************")
		pingWebSites(urls)
		log.Printf("*************** Garbage Collector Completed ***************")

		time.Sleep(timeToRecheck)
	}
}

func pingWebSites(urls []database.Domains) {

	for k := range urls {
		domain := urls[k]
		url := "http://" + domain.Domain
		log.Printf("will ping domain:" + url)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("unable to ping the site")
			continue
		}
		log.Printf("for the domain: " + url + " status code is: " + resp.Status)

		if resp.StatusCode != statusCodeOK {
			removeInvalidUrl(domain)
		}

	}

}

func removeInvalidUrl(domain database.Domains) {
	dbInstance.RemoveDomain(domain)
}
