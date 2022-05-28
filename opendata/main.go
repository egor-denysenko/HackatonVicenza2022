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

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	client := getInfluxClient()
	queryAPI := client.QueryAPI("VincenzoHackaton")
	response := make([]map[string]interface{}, 0)
	// from_timestamp_str := r.URL.Query().Get("from")
	// to_timestamp_str := r.URL.Query().Get("to")
	// var from_timestamp, to_minestamp int
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
	//
	// }
	result, err := queryAPI.Query(context.Background(), `from(bucket:"Hackaton2022")|> range(start: -30m)`)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("error executing query"))
		return
	}
	for result.Next() {
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		fmt.Printf("value: %v\n", result.Record().Values())
		response = append(response, result.Record().Values())
	}

	if result.Err() != nil {
		fmt.Printf("query parsing error: %s\n", result.Err().Error())
		return
	}
	client.Close()
	json.NewEncoder(w).Encode(response)
}

func main() {
	influxToken = os.Args[1]
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET").Queries("api_key", "MyApiKey")

	log.Panic(http.ListenAndServe(":8000", r))
}
