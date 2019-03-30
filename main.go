package main

import (
	server "github.com/drockdriod/proxy-firebase/net"
	_ "github.com/joho/godotenv/autoload"
)

func main(){
	server.Start()
}