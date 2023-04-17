package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Access_token string `env:"ACCESS_TOKEN,required"`
	Url          string `env:"URL,required"`
}

var Con Config

func main() {
	InitConfig(&Con)
	AutoRedeploy()
}

func InitConfig(c *Config) {
	c.Access_token = os.Getenv("ACCESS_TOKEN")
	c.Url = os.Getenv("URL")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

type Output struct {
	Out string `json:"out"`
}

// type ClientMassage struct {
// 	ExecOut string `json:"execOut"`
// 	ExecErr string `json:"execErr"`
// }

func AutoRedeploy() {
	payload := strings.NewReader("{\"access_token\":\"" + Con.Access_token + "\"}")
	req, err := http.NewRequest("POST", Con.Url, payload)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()
	var resp []byte
	resp, err = io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	// var output Output
	// err = json.Unmarshal(resp, &output)
	// if err != nil {
	// 	log.Println(err)
	// }

	log.Println("status:", response.StatusCode, "\nOutput", resp)
}
