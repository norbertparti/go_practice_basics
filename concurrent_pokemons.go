package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Pokemon struct {
	Name string `json:"name"`
}

const baseUrl string = "https://pokeapi.co/api/v2/pokemon/%d"

func getPokemon(id int) (pokemon *Pokemon, err error) {
	url := fmt.Sprintf(baseUrl, id)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&pokemon)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

func getPokemonsSequential() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start))
	}()

	pokemonIds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pokemonMap := make(map[int]*Pokemon, len(pokemonIds))

	for _, id := range pokemonIds {
		pokemon, err := getPokemon(id)
		if err != nil {
			continue
		}
		pokemonMap[id] = pokemon
		fmt.Printf("Fetched pokemon %d with name %v\n", id, pokemon.Name)
	}
}
