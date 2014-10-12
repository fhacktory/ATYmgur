package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type imgur struct {
	CLIENT_ID     string
	CLIENT_SECRET string
	auth_url      string
	token_url     string
	token         oauth2.Token
	pin           string
	cl            http.Client
	trans         *oauth2.Transport
}

func (i *imgur) init() {
	conf, err := oauth2.NewConfig(&oauth2.Options{
		ClientID:     i.CLIENT_ID,
		ClientSecret: i.CLIENT_SECRET,
	}, i.auth_url, i.token_url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Init oauth starting...")

	url_def := conf.AuthCodeURL("state", "online", "auto")
	url_pin, _ := url.Parse(url_def)
	url_pin_q := url_pin.Query()
	url_pin_q.Set("response_type", "pin")
	url_pin.RawQuery = url_pin_q.Encode()
	fmt.Printf("Visit the URL and get the code: \n%v\n\n", url_pin)

	err = json.Unmarshal([]byte(CONFIG.Key), &i.token)
	if err != nil {
		fmt.Print("Enter the PIN Number of the ATM Machine : ")
		fmt.Scanf("%s", &i.pin)
		//hack because we can't change the grant_type in this fuckin lib
		resp, err := http.PostForm("https://api.imgur.com/oauth2/token",
			url.Values{
				"client_id":     {i.CLIENT_ID},
				"client_secret": {i.CLIENT_SECRET},
				"grant_type":    {"pin"},
				"pin":           {i.pin}})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Good PIN trying request")
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		CONFIG.Key = string(body)
		err = json.Unmarshal(body, &i.token)
		if err != nil {
			fmt.Println(string(body))
		}
	} else {
		fmt.Println("Already auth ! lets use that :)")
	}

	i.token.TokenType = "Bearer" //HACK AGAIN T_T it's Bearer not bearer

	i.trans = conf.NewTransport()
	i.trans.SetToken(&i.token)

	i.cl = http.Client{Transport: i.trans}
	save_conf()
}
