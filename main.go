package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
)

func main() {
	get("tears of the sun")

}
func get(keyword string){
	keyword = strings.Replace(keyword, " ", "_", -1)
	url := "https://v2.sg.media-imdb.com/suggests/"+string(keyword[0])+"/"+keyword+".json"

	req, err1 := http.NewRequest("GET", url, nil)
	if err1 == nil {
		req.Header.Add("content-type", "application/json")
		res, err2 := http.DefaultClient.Do(req)
		if err2 == nil {
			defer res.Body.Close()
			body, err3 := ioutil.ReadAll(res.Body)
			if err3 == nil {
				response := strings.Replace(string(body), "imdb$"+keyword+"(", "", -1)
				response = strings.TrimSuffix(response, ")")
				var m APIResponse
				if err := json.Unmarshal([]byte(response), &m); err != nil {
					panic(err)
				}else {
					imdb := IMDB{m.D[0].L, m.D[0].ID, m.D[0].S, m.D[0].Y,m.D[0].Q,m.D[0].I[0].(string)}
					fmt.Println(imdb)
				}
			}
		}
	}
}