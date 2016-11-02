package secretsanta

import (
	"errors"
	"math/rand"
	"time"
)

type playerInfo struct {
	name       string
	cannotHave string
}

// SecretSanta object which can be used to generate a game of gift giving randomness
type SecretSanta struct {
	players []playerInfo
}

// New creates a SecretSanta object which is initialized and ready to use
func New() *SecretSanta {
	return &SecretSanta{make([]playerInfo, 0, 4)}
}

// AddPlayer puts an additional player into the game - a name is mandatory and cannot be
// repeated.  The optional cannotHave string is the name of a player within the game who
// the added player cannot be assigned
func (ss *SecretSanta) AddPlayer(name string, cannotHave string) error {
	if name == "" {
		return errors.New("Name is empty")
	}
	for _, player := range ss.players {
		if name == player.name {
			return errors.New("Player already exists")
		}
	}

	ss.players = append(ss.players, playerInfo{name, cannotHave})

	return nil
}

// Randomize returns a map of Secret Santas -> Receipent
func (ss *SecretSanta) Randomize() (map[string]string, error) {
	gameMap := make(map[string][]string)

	// For each player, generate a list of the players they can have
	for _, player := range ss.players {
		gameMap[player.name] = make([]string, 0)
		for _, otherPlayer := range ss.players {
			if player.name == otherPlayer.name {
				continue // The player should skip themselves
			}
			if player.cannotHave == otherPlayer.name {
				continue // Player should skip people they cannot have
			}
			gameMap[player.name] = append(gameMap[player.name], otherPlayer.name)
		}
	}

	rounds := 0
	for {
		ssmap := make(map[string]string)
		gameMapCopy := make(map[string][]string)
		for k, v := range gameMap {
			gameMapCopy[k] = make([]string, len(gameMap[k]))
			copy(gameMapCopy[k], v)
		}

		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		player := ss.players[rand.Intn(len(ss.players))].name
		first := player

		for i := 0; i < len(ss.players)-1; i++ {
			others := gameMapCopy[player]
			var pick string
			for len(others) > 0 {
				pickIdx := random.Intn(len(others))
				pick = others[pickIdx]
				if _, ok := gameMapCopy[pick]; ok {
					break
				} else {
					others = append(others[:pickIdx], others[pickIdx+1:]...)
				}
			}
			delete(gameMapCopy, player)
			ssmap[player] = pick
			player = pick
		}

		for _, other := range gameMapCopy[player] {
			if other == first {
				ssmap[player] = first
				delete(gameMapCopy, player)
				break
			}
		}
		if len(gameMapCopy) == 0 {
			return ssmap, nil
		}

		rounds++
		if rounds > 1000 {
			break
		}
	}

	return nil, errors.New("Could not find solution")
}
