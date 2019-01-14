package main

import (
	"bufio"
	"gocarbe/api"
	"gocarbe/listen"
	"gocarbe/shared"
	"gocarbe/store"
	"os"
)

func setup(configArray []string) {
	shared.SetUpLog()
	store.SetUpDb()
	api.SetUpApi(configArray[0])
}

func main() {
	keyFile, err := os.Open("config.txt")
	if err != nil {
		shared.LOG(err.Error())
	}

	defer keyFile.Close()

	scanner := bufio.NewScanner(keyFile)

	configArray := make([]string, 0)
	for scanner.Scan() {
		config := scanner.Text()
		configArray = append(configArray, config)
	}

	setup(configArray)

	listen.Start()

	store.CloseDb()
	shared.LOG("END OF LOG")
	shared.CloseLog()
}
