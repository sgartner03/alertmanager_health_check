package webserver

import (
	"alertmanager_health/metrics"
	"fmt"
	"encoding/json"
	"net/http"
)

type IncrementEndpoint struct {
	Metrics metrics.Metrics
}

func (web IncrementEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req := ReadJSON(r)
	seq := Parse(req.Alerts)
	web.Metrics.IncrementSequence(seq)
	fmt.Fprint(w, seq)
}


func ReadJSON(r *http.Request) Request {
	decoder := json.NewDecoder(r.Body)
    	var req Request 
    	err := decoder.Decode(&req)
    	if err != nil {
       		panic(err)
    	}
	return req
}
    

func Parse(alerts []Alert) []string {
	var arr []string
	for _, alert := range alerts {
		arr = append(arr,alert.Labels.GepardecCluster)
	}
	return arr
}
