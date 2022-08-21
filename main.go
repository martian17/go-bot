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



//main functions

//when payload == nil
//for optimization but not sure if im gonna use it
func makeRequestNil (method string, url string) (*http.Response, error) {
    url = "https://discord.com/api/v10"+url
    client := &http.Client{}
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        log.Println("could not create request: ")
        return nil, err
	}
    
	req.Header.Set("Authorization", "Bot " + token)
    //req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
    if err != nil {
		log.Println("request do failed: ")
        return nil, err
	}
    return resp, nil
}

func makeRequest (method string, url string, payload string) (*http.Response, error) {
    url = "https://discord.com/api/v10"+url
    client := &http.Client{}
    req, err := http.NewRequest(method, url, strings.NewReader(payload))
    if err != nil {
        log.Println("could not create request: ")
        return nil, err
	}
    
	req.Header.Set("Authorization", "Bot " + token)
    req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
    if err != nil {
		log.Println("request do failed: ")
        return nil, err
	}
    return resp, nil
}

func sendMessage(channel string, msg string) error {
    _,err := makeRequest("POST","/channels/"+channel+"/messages",msg)
    return err
}

func printRequest(method string, url string, payload string) {
	resp, err := makeRequest(method,url,payload)
    if err != nil {
		log.Println(err)
        return
	}
    
    defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
    if err != nil {
		log.Println("request body read failed: ",err)
        return
	}

	fmt.Println(string(b))
}


func main(){
    loadCredentials()
    
    //https://discord.com/channels/1010708479142547547/1010708479742320763
    
    //example codes
    //post message
    //sendMessage("1010708479742320763","hi there")
    
    //get channel state
    //printRequest("GET","/channels/1010708479742320763","")
    
    //get channel message history
    //printRequest("GET","/channels/1010708479742320763/messages","")
    
    //get list of registered slash commands
    printRequest("GET","/applications/"+appid+"/commands","")
    
    //delete command
    //printRequest("DELETE","/applications/"+appid+"/commands/1010341986890420224","")
    
}





