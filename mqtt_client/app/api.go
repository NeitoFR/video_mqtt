package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	
	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"

	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load("./.env")

	// _initMqttConnection(mqttInfo{os.Getenv("mqtt_server"), os.Getenv("mqtt_port")})
}
