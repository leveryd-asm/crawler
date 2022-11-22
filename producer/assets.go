package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return body
}

func QuerySubdomain(queryParam string) []byte {
	// url := assetEndpoint + "/api/info/subdomain/query?limit=10&offset=10"
	url := assetEndpoint + "/api/info/subdomain/query?" + queryParam
	body := httpGet(url)
	// println(string(body))
	return body
}
