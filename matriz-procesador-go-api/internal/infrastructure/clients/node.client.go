package clients

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type NodeClient struct {
    apiUrl string
}

type NodeClientInterface interface {
	SendToStats(data interface{}) (*http.Response, error)
}


func NewNodeClient(url string) *NodeClient {
    return &NodeClient{apiUrl: url}
}

func (nc *NodeClient) SendToStats(data interface{}) (*http.Response, error) {
    log.Println("Enviando datos a Node API")
    jsonData, err := json.Marshal(data)
    if err != nil { return nil, err }
    
    return http.Post(nc.apiUrl, "application/json", bytes.NewBuffer(jsonData))
}