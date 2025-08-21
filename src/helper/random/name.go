package random

import (
	"github.com/cynx-io/cynx-core/src/constant"
	"math/rand"
	"strings"
)

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
