package random

import (
	"github.com/cynxees/cynx-core/src/constant"
	"math/rand"
	"strings"
)

var adjectives = []string{
	"Ancient", "Bold", "Creative", "Daring", "Elegant", "Fancy", "Gentle",
	"Happy", "Icy", "Jolly", "Kind", "Lively", "Mighty", "Noble", "Odd",
	"Peaceful", "Quick", "Royal", "Silent", "Tough", "Unique", "Vivid", "Wild", "Young", "Zesty",
}

var verbs = []string{
	"Running", "Flying", "Jumping", "Singing", "Dancing", "Hunting",
	"Racing", "Shining", "Glowing", "Roaring", "Sprinting", "Soaring",
	"Climbing", "Leaping", "Whispering",
}

func RandomAnimalName(separator string) string {

	animal := constant.Animals[rand.Intn(len(constant.Animals))]
	verb := constant.Verbs[rand.Intn(len(constant.Verbs))]
	adj := constant.Adjectives[rand.Intn(len(constant.Adjectives))]
	randomNumbers := RandomNumbers(3, 3)

	seed := RandomIntInRange(0, 3)
	switch seed {
	case 0:
		// Verb + Animal
		return strings.Join([]string{verb, animal, randomNumbers}, separator)
	case 1:
		// Adjective + Animal
		return strings.Join([]string{adj, animal, randomNumbers}, separator)
	case 2:
		// Adjective + Verb + Animal
		return strings.Join([]string{adj, verb, animal, randomNumbers}, separator)
	case 3:
		// Verb + Adjective + Animal
		return strings.Join([]string{verb, adj, animal, randomNumbers}, separator)
	}

	return strings.Join([]string{verb, adj, animal, randomNumbers}, separator)
}
