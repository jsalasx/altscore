package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"sort"
// 	"strconv"
// )

// type PokemonTypeResponse struct {
// 	Count    int           `json:"count"`
// 	Next     string        `json:"next"`
// 	Previous interface{}   `json:"previous"` // Puede ser null, por lo que se usa interface{}
// 	Results  []PokemonType `json:"results"`
// }

// type PokemonType struct {
// 	Name string `json:"name"`
// 	URL  string `json:"url"`
// }

// type PokemonPagResponse struct {
// 	Count    int          `json:"count"`
// 	Next     string       `json:"next"`
// 	Previous interface{}  `json:"previous"` // Puede ser null, por lo que se usa interface{}
// 	Results  []PokemonPag `json:"results"`
// }

// type PokemonPag struct {
// 	Name   string   `json:"name"`
// 	URL    string   `json:"url"`
// 	Height float64  `json:"height"`
// 	Types  []string `json:"types"`
// }

// type PokemonTypeMapa struct {
// 	Name              string  `json:"name"`
// 	URL               string  `json:"url"`
// 	NumeroDePokemons  int     `json:"numero_de_pokemons"`
// 	SumaDeAlturas     int     `json:"suma_de_alturas"`
// 	PromedioDeAlturas float64 `json:"promedio_de_alturas"`
// }

// type PokemonInfoHeigth struct {
// 	Height int                `json:"height"`
// 	Name   string             `json:"name"`
// 	Types  []PokemonInfoTypes `json:"types"`
// }

// type PokemonInfoTypes struct {
// 	Slot int                `json:"slot"`
// 	Type PokemonInfoTypeOne `json:"type"`
// }

// type PokemonInfoTypeOne struct {
// 	Name string `json:"name"`
// 	URL  string `json:"url"`
// }

// func main() {
// 	types := getPokemonTypesData()
// 	mapTypes := make(map[string]PokemonTypeMapa)

// 	for i := 0; i < len(types); i++ {
// 		if _, ok := mapTypes[types[i].URL]; !ok {
// 			mapTypes[types[i].URL] = PokemonTypeMapa{
// 				URL:               types[i].URL,
// 				Name:              types[i].Name,
// 				NumeroDePokemons:  0,
// 				SumaDeAlturas:     0,
// 				PromedioDeAlturas: 0,
// 			}
// 		}

// 	}

// 	numPagesPokemons := 4
// 	var pokemons []PokemonPag
// 	for i := 0; i < numPagesPokemons; i++ {
// 		pokemons = append(pokemons, getPokemons(i)...)
// 	}
// 	fmt.Println("Cantidad de pokemons: ", len(pokemons))
// 	for i := 0; i < len(pokemons); i++ {
// 		pokemonInfo := getPokemonHeigth(pokemons[i].URL)
// 		if pokemonInfo == nil {
// 			fmt.Printf("Error obteniendo la info del pokemon: %s\n", pokemons[i].Name)
// 			continue
// 		}

// 		pokemons[i].Height = float64(pokemonInfo.Height)
// 		for j := 0; j < len(pokemonInfo.Types); j++ {
// 			if pokemonType, ok := mapTypes[pokemonInfo.Types[j].Type.URL]; ok {
// 				if pokemonInfo.Types[j].Type.Name == pokemonType.Name {
// 					var PokemonTypeMapa = PokemonTypeMapa{
// 						URL:               pokemonType.URL,
// 						Name:              pokemonType.Name,
// 						NumeroDePokemons:  pokemonType.NumeroDePokemons + 1,
// 						SumaDeAlturas:     pokemonType.SumaDeAlturas + (pokemonInfo.Height),
// 						PromedioDeAlturas: 0,
// 					}
// 					mapTypes[pokemonInfo.Types[j].Type.URL] = PokemonTypeMapa
// 				}
// 			}
// 		}

// 	}

// 	for key, value := range mapTypes {
// 		fmt.Println(value.Name)
// 		fmt.Println(value.NumeroDePokemons)
// 		if value.NumeroDePokemons > 0 {
// 			promedioAltura := float64(value.SumaDeAlturas) / float64(value.NumeroDePokemons)
// 			value.PromedioDeAlturas, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", promedioAltura), 64)
// 		}
// 		mapTypes[key] = value
// 	}

// 	mapTypesJson, _ := json.Marshal(mapTypes)
// 	fmt.Printf("Mapa de tipos de pokemons: %v\n", string(mapTypesJson))
// 	publicResponse(mapTypes)
// }

// func getPokemonTypesData() []PokemonType {

// 	url := "https://pokeapi.co/api/v2/type/?offset=0&limit=23"

// 	client := &http.Client{}

// 	// Crear el request
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		fmt.Printf("Error creando el request: %v\n", err)
// 		return nil
// 	}

// 	// Añadir headers
// 	//req.Header.Add("API-KEY", "791e97b98aa24e4d8fd52121141e94d8")
// 	req.Header.Add("Accept", "application/json")

// 	// Hacer el request
// 	response, err := client.Do(req)
// 	if err != nil {
// 		fmt.Printf("Error haciendo GET request: %v\n", err)
// 		return nil
// 	}
// 	defer response.Body.Close()

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Printf("Error leyendo el body: %v\n", err)
// 		return nil
// 	}

// 	// // Imprimir la respuesta
// 	// fmt.Println("Respuesta:")
// 	// fmt.Println(string(body))
// 	var resp PokemonTypeResponse
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		fmt.Printf("Error decodificando la respuesta: %v\n", err.Error())
// 		return nil
// 	}
// 	//fmt.Println(string(body))

// 	var types []PokemonType
// 	for i := 0; i < len(resp.Results); i++ {
// 		pokemonType := resp.Results[i]
// 		types = append(types, pokemonType)
// 	}

// 	return types
// }

// func getPokemons(page int) []PokemonPag {
// 	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/?offset=%d&limit=500", page*500)

// 	client := &http.Client{}

// 	// Crear el request
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		fmt.Printf("Error creando el request: %v\n", err)
// 		return nil
// 	}

// 	// Añadir headers
// 	//req.Header.Add("API-KEY", "791e97b98aa24e4d8fd52121141e94d8")
// 	req.Header.Add("Accept", "application/json")

// 	// Hacer el request
// 	response, err := client.Do(req)
// 	if err != nil {
// 		fmt.Printf("Error haciendo GET request: %v\n", err)
// 		return nil
// 	}
// 	defer response.Body.Close()

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Printf("Error leyendo el body: %v\n", err)
// 		return nil
// 	}

// 	// // Imprimir la respuesta
// 	// fmt.Println("Respuesta:")
// 	// fmt.Println(string(body))
// 	var resp PokemonPagResponse
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		fmt.Printf("Error decodificando la respuesta: %v\n", err.Error())
// 		return nil
// 	}
// 	//fmt.Println(string(body))

// 	var pokemons []PokemonPag
// 	for i := 0; i < len(resp.Results); i++ {
// 		pokemon := resp.Results[i]
// 		pokemons = append(pokemons, pokemon)
// 	}

// 	return pokemons
// }

// func getPokemonHeigth(url string) *PokemonInfoHeigth {
// 	client := &http.Client{}

// 	// Crear el request
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		fmt.Printf("Error creando el request: %v\n", err)
// 		return nil
// 	}

// 	// Añadir headers
// 	//req.Header.Add("API-KEY", "791e97b98aa24e4d8fd52121141e94d8")
// 	req.Header.Add("Accept", "application/json")

// 	// Hacer el request
// 	response, err := client.Do(req)
// 	if err != nil {
// 		fmt.Printf("Error haciendo GET request: %v\n", err)
// 		return nil
// 	}
// 	defer response.Body.Close()

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Printf("Error leyendo el body: %v\n", err)
// 		return nil
// 	}

// 	// // Imprimir la respuesta
// 	// fmt.Println("Respuesta:")
// 	// fmt.Println(string(body))
// 	var resp PokemonInfoHeigth
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		fmt.Printf("Error decodificando la respuesta: %v\n", err.Error())
// 		return nil
// 	}
// 	//fmt.Println(string(body))

// 	return &resp
// }

// type HeightsResponse struct {
// 	Heights map[string]float64 `json:"heights"`
// }

// func publicResponse(data map[string]PokemonTypeMapa) error {

// 	// Crear la estructura HeightsResponse
// 	var heights HeightsResponse = HeightsResponse{Heights: make(map[string]float64)}

// 	// Poblar el mapa Heights
// 	for _, value := range data {
// 		heights.Heights[value.Name] = value.PromedioDeAlturas
// 	}

// 	// Extraer las claves y ordenarlas alfabéticamente
// 	keys := make([]string, 0, len(heights.Heights))
// 	for key := range heights.Heights {
// 		keys = append(keys, key)
// 	}
// 	sort.Strings(keys)

// 	// Crear un nuevo mapa ordenado (opcional, solo para serialización ordenada)
// 	orderedHeights := HeightsResponse{Heights: make(map[string]float64)}
// 	for _, key := range keys {
// 		orderedHeights.Heights[key] = heights.Heights[key]
// 	}

// 	// Serializar el nuevo mapa ordenado
// 	bodyPost, err := json.Marshal(orderedHeights)
// 	if err != nil {
// 		fmt.Printf("Error codificando la respuesta: %v\n", err)
// 		return err
// 	}
// 	fmt.Printf("Respuesta ordenada: %v\n", string(bodyPost))
// 	client := &http.Client{}
// 	url := "https://makers-challenge.altscore.ai/v1/s1/e6/solution"
// 	// Crear el request
// 	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyPost))
// 	if err != nil {
// 		fmt.Printf("Error creando el request: %v\n", err)
// 		return err
// 	}

// 	// Añadir headers
// 	req.Header.Add("API-KEY", "791e97b98aa24e4d8fd52121141e94d8")
// 	req.Header.Add("Accept", "application/json")

// 	// Hacer el request
// 	response, err := client.Do(req)
// 	if err != nil {
// 		fmt.Printf("Error haciendo GET request: %v\n", err)
// 		return err
// 	}
// 	defer response.Body.Close()

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Printf("Error leyendo el body: %v\n", err)
// 		return err
// 	}

// 	// // Imprimir la respuesta
// 	fmt.Println("Respuesta______________________________________________________________asdasdas:")
// 	fmt.Println("Respuesta:")
// 	fmt.Println(string(body))

// 	//fmt.Println(string(body))

// 	return nil

// }
