package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	influxClient influxdb2.Client
	influxToken  string
)

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("received: %s, %s\n", msg.Payload(), msg.Topic())
	f, err := strconv.ParseFloat(string(msg.Payload()), 64)
	if err != nil {
		fmt.Println("woo")
		panic("e")
	}
	writeAPI := influxClient.WriteAPI("VincenzoHackaton", "Hackaton2022")
	if msg.Topic() == "/temp" {
		writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature temp=%f", f))
	} else if msg.Topic() == "/humidity" {
		writeAPI.WriteRecord(fmt.Sprintf("stat,unit=humidity humidity=%f", f))
	}
	writeAPI.Flush()
}

func subscribe(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 0, nil)
	token.Wait()
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("connected")
}

func connectLostHandler(client mqtt.Client, err error) {
	fmt.Printf("connect lost %v\n", err)
}

func createInfluxClient() {
	influxClient = influxdb2.NewClient("http://192.168.15.90:8086", influxToken)
}

func main() {
	influxToken = os.Getenv("INFLUX_TOKEN")

	createInfluxClient()
	var broker = "192.168.15.90"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("mqtt_receiver")
	opts.SetDefaultPublishHandler(messageHandler)
	opts.SetConnectionLostHandler(connectLostHandler)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	subscribe(client, "/temp")
	subscribe(client, "/humidity")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	influxClient.Close()
	client.Disconnect(250)
	os.Exit(0)
}
