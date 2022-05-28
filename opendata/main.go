package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	influxToken string
)

func getInfluxClient() influxdb2.Client {
	return influxdb2.NewClient("http://192.168.15.90:8086", influxToken)
}

func execQuery(query string) []map[string]interface{} {
	client := getInfluxClient()
	queryAPI := client.QueryAPI("VincenzoHackaton")
	response := make([]map[string]interface{}, 0)
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil
	}
	for result.Next() {
		response = append(response, result.Record().Values())
	}

	if result.Err() != nil {
		fmt.Printf("query parsing error: %s\n", result.Err().Error())
		return nil
	}
	client.Close()
	return response
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// from_timestamp_str := r.URL.Query().Get("from")
	// to_timestamp_str := r.URL.Query().Get("to")
	// var from_timestamp, to_minestamp int
	var response []map[string]interface{}
	// if from_timestamp_str != "" && to_timestamp_str != "" {
	// 	from_timestamp, err := strconv.Atoi(from_timestamp_str)
	// 	if err != nil {
	// 		w.Write([]byte("error"))
	// 		return
	// 	}
	//
	// 	to_timestamp, err := strconv.Atoi(to_timestamp_str)
	// 	if err != nil {
	// 		w.Write([]byte("error"))
	// 		return
	// 	}
	//        response = execQuery(`from(bucket:"Hackaton2022")|> range(start: -30m) |> filter(fn: (r) => r._start >)`)
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	response = execQuery(`from(bucket:"Hackaton2022")|> range(start: -30m)`)
	json.NewEncoder(w).Encode(response)
}

func main() {
	influxToken = os.Getenv("INFLUX_TOKEN")
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET").Queries("api_key", "MyApiKey")

	log.Panic(http.ListenAndServe(":8000", r))
}
