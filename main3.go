package main

// import (
// 	"encoding/base64"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"strings"
// )

// type Planet struct {
// 	Name           string   `json:"name"`
// 	RotationPeriod string   `json:"rotation_period"`
// 	OrbitalPeriod  string   `json:"orbital_period"`
// 	Diameter       string   `json:"diameter"`
// 	Climate        string   `json:"climate"`
// 	Gravity        string   `json:"gravity"`
// 	Terrain        string   `json:"terrain"`
// 	SurfaceWater   string   `json:"surface_water"`
// 	Population     string   `json:"population"`
// 	Residents      []string `json:"residents"`
// 	Films          []string `json:"films"`
// 	Created        string   `json:"created"`
// 	Edited         string   `json:"edited"`
// 	URL            string   `json:"url"`
// }

// type Character struct {
// 	Name      string   `json:"name"`
// 	Height    string   `json:"height"`
// 	Mass      string   `json:"mass"`
// 	HairColor string   `json:"hair_color"`
// 	SkinColor string   `json:"skin_color"`
// 	EyeColor  string   `json:"eye_color"`
// 	BirthYear string   `json:"birth_year"`
// 	Gender    string   `json:"gender"`
// 	Homeworld string   `json:"homeworld"`
// 	Films     []string `json:"films"`
// 	Species   []string `json:"species"`
// 	Vehicles  []string `json:"vehicles"`
// 	Starships []string `json:"starships"`
// 	Created   string   `json:"created"`
// 	Edited    string   `json:"edited"`
// 	URL       string   `json:"url"`
// 	IsGood    *bool    `json:"is_good"`
// }

// type ApiPersonResponse struct {
// 	Count    int         `json:"count"`
// 	Next     string      `json:"next"`
// 	Previous string      `json:"previous"`
// 	Results  []Character `json:"results"`
// }

// type ApiPlanetResponse struct {
// 	Count    int      `json:"count"`
// 	Next     string   `json:"next"`
// 	Previous string   `json:"previous"`
// 	Results  []Planet `json:"results"`
// }

// type RolodexResponse struct {
// 	OracleNotes string `json:"oracle_notes"`
// }

// type PoblacionPlaneta struct {
// 	NumBuenos       int     `json:"numBuenos"`
// 	NumMalos        int     `json:"numMalos"`
// 	IndicedeBalance float64 `json:"indiceDeBalance"`
// }

// func main() {

// 	var personajes []Character
// 	var planetas []Planet

// 	paginasPersonajes := 9
// 	paginasPlanetas := 6

// 	for i := 1; i <= paginasPersonajes; i++ {
// 		personajes = append(personajes, getPersonasData(i)...)
// 	}

// 	for i := 1; i <= paginasPlanetas; i++ {
// 		planetas = append(planetas, getPlanetasData(i)...)
// 	}

// 	mapaDePersonajes := make(map[string]Character)

// 	for i := 0; i < len(personajes); i++ {
// 		personaje := personajes[i]
// 		esBueno, err := getRolodex(personaje.Name)
// 		if err != nil {
// 			fmt.Printf("Error obteniendo el rolodex: %v\n", err)
// 			continue
// 		}
// 		personaje.IsGood = &esBueno
// 		personajes[i] = personaje

// 		if _, ok := mapaDePersonajes[personaje.URL]; !ok {
// 			mapaDePersonajes[personaje.URL] = personaje
// 		}

// 	}
// 	mapaDePlanetas := make(map[string]PoblacionPlaneta)
// 	for i := 0; i < len(planetas); i++ {
// 		mapaDePlanetas[planetas[i].Name] = PoblacionPlaneta{}
// 		for j := 0; j < len(planetas[i].Residents); j++ {
// 			if personaje, ok := mapaDePersonajes[planetas[i].Residents[j]]; ok {
// 				if *personaje.IsGood {
// 					numeroDeBuenosActual := mapaDePlanetas[planetas[i].Name].NumBuenos
// 					numeroDeMalosActual := mapaDePlanetas[planetas[i].Name].NumMalos
// 					mapaDePlanetas[planetas[i].Name] = PoblacionPlaneta{NumBuenos: numeroDeBuenosActual + 1, NumMalos: numeroDeMalosActual}
// 				} else {
// 					numeroDeBuenosActual := mapaDePlanetas[planetas[i].Name].NumBuenos
// 					numeroDeMalosActual := mapaDePlanetas[planetas[i].Name].NumMalos
// 					mapaDePlanetas[planetas[i].Name] = PoblacionPlaneta{NumMalos: numeroDeMalosActual + 1, NumBuenos: numeroDeBuenosActual}
// 				}
// 			}

// 		}

// 	}

// 	for key, value := range mapaDePlanetas {
// 		indiceDeBalance := float64(value.NumBuenos-value.NumMalos) / float64(value.NumBuenos+value.NumMalos)
// 		value.IndicedeBalance = indiceDeBalance
// 		mapaDePlanetas[key] = value
// 	}

// 	for key, value := range mapaDePlanetas {
// 		fmt.Printf("Planeta: %s\n", key)
// 		fmt.Printf("Número de buenos: %d\n", value.NumBuenos)
// 		fmt.Printf("Número de malos: %d\n", value.NumMalos)
// 		fmt.Printf("Índice de balance: %f\n", value.IndicedeBalance)
// 		if value.IndicedeBalance == 0 {
// 			fmt.Println("Neutral")
// 		}
// 	}

// }

// func getPersonasData(page int) []Character {

// 	url := fmt.Sprintf("https://swapi.dev/api/people/?page=%d", page)

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
// 	var resp ApiPersonResponse
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		fmt.Printf("Error decodificando la respuesta: %v\n", err.Error())
// 		return nil
// 	}
// 	//fmt.Println(string(body))

// 	var personajes []Character
// 	for i := 0; i < len(resp.Results); i++ {
// 		personaje := resp.Results[i]
// 		personajes = append(personajes, personaje)
// 	}

// 	return personajes
// }

// func getPlanetasData(page int) []Planet {

// 	url := fmt.Sprintf("https://swapi.dev/api/planets/?page=%d", page)

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
// 	var resp ApiPlanetResponse
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		fmt.Printf("Error decodificando la respuesta: %v\n", err.Error())
// 		return nil
// 	}
// 	//fmt.Println(string(body))

// 	var planetas []Planet
// 	for i := 0; i < len(resp.Results); i++ {
// 		planeta := resp.Results[i]
// 		planetas = append(planetas, planeta)
// 	}

// 	return planetas
// }

// func getRolodex(namePersonaje string) (bool, error) {
// 	fmt.Println("Buscando en el rolodex")
// 	fmt.Println(namePersonaje)
// 	urlAux := "https://makers-challenge.altscore.ai/v1/s1/e3/resources/oracle-rolodex?name="
// 	escapedURL := url.QueryEscape(namePersonaje)
// 	escapedURL = urlAux + escapedURL
// 	client := &http.Client{}

// 	// Crear el request
// 	req, err := http.NewRequest("GET", escapedURL, nil)
// 	if err != nil {
// 		fmt.Printf("Error creando el request: %v\n", err)
// 		return false, errors.New("no se encontró el texto esperado")
// 	}

// 	// Añadir headers
// 	req.Header.Add("API-KEY", "791e97b98aa24e4d8fd52121141e94d8")
// 	req.Header.Add("Accept", "application/json")

// 	// Hacer el request
// 	response, err := client.Do(req)
// 	if err != nil {
// 		fmt.Printf("Error haciendo GET request: %v\n", err)
// 		return false, errors.New("no se encontró el texto esperado")
// 	}
// 	defer response.Body.Close()

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Printf("Error leyendo el body: %v\n", err)
// 		return false, errors.New("no se encontró el texto esperado")
// 	}

// 	// // Imprimir la respuesta
// 	// fmt.Println("Respuesta:")
// 	// fmt.Println(string(body))
// 	var resp RolodexResponse
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		fmt.Printf("Error decodificando la respuesta: %v\n", err.Error())
// 		return false, errors.New("no se encontró el texto esperado")
// 	}
// 	fmt.Println(string(body))

// 	decoded, err := base64.StdEncoding.DecodeString(resp.OracleNotes)
// 	if err != nil {
// 		fmt.Printf("Error al decodificar Base64: %v\n", err)
// 		return false, errors.New("no se encontró el texto esperado")
// 	}
// 	textoBueno := "Light Side of the Force."
// 	textoMalo := "Dark Side of the Force."
// 	if strings.Contains(string(decoded), textoBueno) {
// 		return true, nil
// 	} else if strings.Contains(string(decoded), textoMalo) {
// 		return false, nil
// 	}
// 	return false, errors.New("no se encontró el texto esperado")

// }
