package main

import (
	"fmt"
	"io/ioutil"
	"secretsanta"
)

type playerInfo struct {
	name       string
	cannotHave string
}

var players = []playerInfo{
	{"Jeff", "Kat"},
	{"Kat", "Jeff"},
	{"Steve", "Liz"},
	{"Liz", "Steve"},
	{"Nancy", "Doug"},
	{"Doug", "Nancy"},
	{"Greg", "Megan"},
	{"Megan", "Greg"},
	{"Julie", ""},
	{"Charlotte", "Gabe"},
	{"Gabe", "Charlotte"},
	{"Chris", "Christine"},
	{"Christine", "Chris"},
}

func main() {
	ss := secretsanta.New()
	for _, player := range players {
		ss.AddPlayer(player.name, player.cannotHave)
	}
	ssmap, err := ss.Randomize()
	if err != nil {
		fmt.Println("Secret Santa failed! Error:", err)
		return
	}

	for k, v := range ssmap {
		filename := "/tmp/" + k + "_SecretSanta.txt"
		ioutil.WriteFile(filename, []byte(v), 0666)
	}
}
