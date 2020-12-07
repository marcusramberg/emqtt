package main

import ( "fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	fmt.Println("OH HAI")
	opts := MQTT.NewClientOptions()
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
