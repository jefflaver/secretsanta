package secretsanta

import (
	"fmt"
	"testing"
)

func TestSecretSanta_AddPlayer(t *testing.T) {
	ss := New()
	if ss.AddPlayer("Jeff", "") != nil {
		t.Fatal("Unable to add single Player")
	}
	if ss.AddPlayer("New", "Person") != nil {
		t.Fatal("Unable to add new player with cannot have restriction")
	}
	if ss.AddPlayer("Jeff", "") == nil {
		t.Fatal("Duplicate player allowed to be added")
	}
	if ss.AddPlayer("Jeff", "Test") == nil {
		t.Fatal("Duplicate player allowed to be added, with cannot have restriction")
	}
	if ss.AddPlayer("", "") == nil {
		t.Fatal("Nameless player allowed to be added")
	}
}

func TestUnsolveableGame(t *testing.T) {
	ss := New()
	ss.AddPlayer("Jeff", "Kat")
	ss.AddPlayer("Kat", "Jeff")
	if ssmap, err := ss.Randomize(); err == nil {
		t.Fatal("Randomize returned success despite players having no choices, output:", ssmap)
	}
	ss.AddPlayer("Bob", "")
	if ssmap, err := ss.Randomize(); err == nil {
		t.Fatal("Randomize returned success despite unsolveable game, output:", ssmap)
	}
}

func TestSolveableGameResult(t *testing.T) {
	ss := New()
	type testData struct {
		name, notHave string
	}
	data := []testData{
		{"Jeff", "Julie"},
		{"Julie", "Kat"},
		{"Kat", "Jeff"},
	}
	for _, v := range data {
		ss.AddPlayer(v.name, v.notHave)
	}
	ssmap, err := ss.Randomize()
	if err != nil {
		t.Fatal("Randomize returned unsuccessful for solveable game")
	}
	if v, ok := ssmap["Jeff"]; !ok || v != "Kat" {
		t.Fatal("Jeff didn't end up having Kat, when only possible solution")
	}
	if v, ok := ssmap["Julie"]; !ok || v != "Jeff" {
		t.Fatal("Julie didn't end up having Jeff, when only possible solution")
	}
	if v, ok := ssmap["Kat"]; !ok || v != "Julie" {
		t.Fatal("Kat didn't end up having Julie, when only possible solution")
	}
}

func ExampleSecretSanta_Randomize() {
	// Create a new SecretSanta game
	ss := New()

	// Add two players to the game
	ss.AddPlayer("Jeff", "")
	ss.AddPlayer("Kat", "")

	// Run a simulation
	ssmap, _ := ss.Randomize()

	// Kat will have Jeff, and Jeff will have Kat
	fmt.Println(ssmap["Jeff"])
	fmt.Println(ssmap["Kat"])
	// Output:
	// Kat
	// Jeff
}
