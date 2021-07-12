package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type TheClient struct {
	UserName string
	Apples   []Apple
}

func NewTheClient(userName string) *TheClient {
	client := &TheClient{UserName: userName}
	client.load()
	return client
}

func (c *TheClient) GetAppleByHash(hash string) (*Apple, error) {
	for _, apple := range c.Apples {
		curHash := HashAppleV3(&apple)
		if curHash == hash {
			return &apple, nil
		}
	}

	return nil, fmt.Errorf("Cannot find an apple by hash " + hash)
}

func (c *TheClient) AddApple(apple *Apple) {
	c.Apples = append(c.Apples, *apple)
	c.save()
}

func (c *TheClient) UpdateApple(apple *Apple) {
	hash := HashAppleV3(apple)
	c.DeleteAppleByHash(hash)
	c.AddApple(apple)
}

func (c *TheClient) DeleteAppleByHash(hash string) {
	var filteredApples []Apple

	for _, apple := range c.Apples {
		curHash := HashAppleV3(&apple)
		if curHash == hash {
			continue
		}
		filteredApples = append(filteredApples, apple)
	}

	c.Apples = filteredApples
	c.save()
}

// =================== Private helpers ======================

func (c *TheClient) save() {
	jsonString, err := json.Marshal(c.Apples)

	if err != nil {
		log.Fatalln("Unable to save file", err)
	}

	_ = ioutil.WriteFile("/tmp/"+c.UserName+".json", jsonString, 0644)
}

func (c *TheClient) load() {
	file, err := ioutil.ReadFile("/tmp/" + c.UserName + ".json")

	if err != nil {
		return
	}

	var apples []Apple

	err = json.Unmarshal(file, &apples)

	if err != nil {
		log.Fatalln("Unable to read file", err)
	}

	c.Apples = apples
}
