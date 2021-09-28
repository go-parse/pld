/*
 * Copyright 2021 Vasiliy Vdovin
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pld

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Proxy struct {
	Address   string
	Port      string
	Anonymity string
	Country   string
}

func Get() []Proxy {

	r := make([]Proxy, 0)

	var d []struct {
		Updated string `json:"UPDATED"`
		Total   string `json:"TOTAL"`
		Lista   []struct {
			IP      string `json:"IP"`
			Port    string `json:"PORT"`
			ISO     string `json:"ISO"`
			Anon    string `json:"ANON"`
			Country string `json:"COUNTRY"`
		} `json:"LISTA"`
	}

	req, err := http.NewRequest("GET", "https://www.proxy-list.download/api/v0/get?l=en&t=http", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Host = "www.proxy-list.download"
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15")
	req.Header.Set("Accept-Language", "en-gb")
	req.Header.Set("Referer", "https://www.proxy-list.download/HTTP")
	req.Header.Set("Connection", "keep-alive")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if b, e := ioutil.ReadAll(resp.Body); e == nil {

		if e := json.Unmarshal(b, &d); e != nil {
			log.Fatal(e)
		}

		for _, l := range d {

			for _, i := range l.Lista {

				r = append(r, Proxy{
					Address:   i.IP,
					Port:      i.Port,
					Anonymity: i.Anon,
					Country:   i.ISO,
				})
			}
		}
	}

	return r
}
