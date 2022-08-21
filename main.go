package main


import (
	//"bytes"
    "strings"
	"encoding/json"
	//"errors"
	"fmt"
	"os"
    //"io/ioutil"
	//"net"
	"net/http"
	//"strconv"
    "log"
    "io"
)

type tokenJSON struct {
    Token string `json:"token"`
    AppID string `json:"appid"`
}
var token string
var appid string

func loadCredentials() {
    content, err := os.ReadFile("./token.json")
    if err != nil {
        log.Fatal("Failed to load token: ", err)
    }
    var tokenjson tokenJSON
    err = json.Unmarshal(content, &tokenjson)
    if err != nil {
        log.Fatal("token.json wrong format: ", err)
    }
    appid, token = tokenjson.AppID, tokenjson.Token
}


func sendMessage(channel string, msg string){
    client := &http.Client{}
    //https://discord.com/channels/1010708479142547547/1010708479742320763
    req, err := http.NewRequest("POST", "https://discord.com/api/v10/channels/"+channel+"/messages", 
        strings.NewReader("{\"content\":\""+msg+"\"}"))
	if err != nil {
		log.Println(": ", err)
        return
	}
    
	req.Header.Set("Authorization", "Bot " + token)
    req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
    if err != nil {
		log.Println(": ",err)
        return
	}
    
    defer resp.Body.Close()
}

func printGetRequest(url string){
    client := &http.Client{}
    //https://discord.com/channels/1010708479142547547/1010708479742320763
    req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(": ", err)
	}
    
	req.Header.Set("Authorization", "Bot " + token)
    //req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
    
    defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
    if err != nil {
		log.Fatal(": ",err)
	}

	fmt.Println(string(b))
}

func main(){
    loadCredentials()
    sendMessage("1010708479742320763","hi there")
    printGetRequest("https://discord.com/api/v10/channels/1010708479742320763")
    printGetRequest("https://discord.com/api/v10/channels/1010708479742320763/messages")
}





