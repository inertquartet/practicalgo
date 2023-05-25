package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println(githubInfo("inertquartet"))
}

// githubInfo returns name and number of public repos for login
func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: %s", err)
		/*
			log.Printf("error: %s", err)
			os.Exit(1)
		*/
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}
	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.NumRepos, nil
}

type Reply struct {
	Name     string
	NumRepos int `json:"public_repos"`
}

/*
   	JSON <--> GO

   true/false <--> true/false
   string <--> string
   null <-> nil
   number <-> float64 (default), float32, int8, int16, int32, int64, int, uint8, ...
   array <--> []any] ([]interface{})
   object <--> map[string]any, struct

   JSON -> io.Reader -> Go: json.Decoder
   JSON -> []byte -> Go: json.Unmarshal
   Go -> io.Writer -> JSON: json.Encoder
   Go -> []byte -> JSON: json.Marshal
*/
