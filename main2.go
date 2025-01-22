package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// type ResonanciaResponse struct {
// 	Id        string   `json:"id"`
// 	Resonance int      `json:"resonance"`
// 	Position  Position `json:"position"`
// }

// type Position struct {
// 	X float64 `json:"x"`
// 	Y float64 `json:"y"`
// 	Z float64 `json:"z"`
// }

// func main() {

// 	var totalPages int = 34
// 	var sumResonance int = 0
// 	for i := 1; i <= totalPages; i++ {
// 		resp := getData(i)
// 		if resp != nil && len(resp) > 0 {
// 			for j := 0; j < len(resp); j++ {
// 				// fmt.Printf("Id: %s, Resonancia: %d, Posición: %f, %f, %f\n", resp[j].Id, resp[j].Resonance, resp[j].Position.X, resp[j].Position.Y, resp[j].Position.Z)
// 				sumResonance += resp[j].Resonance
// 			}

// 		}
// 	}

// 	fmt.Printf("Suma de resonancias: %d\n", sumResonance)
// 	promedio := float64(sumResonance) / float64(100)
// 	fmt.Printf("Promedio de resonancias: %f\n", promedio)

// }

// func getData(page int) []ResonanciaResponse {

// 	url := fmt.Sprintf("https://makers-challenge.altscore.ai/v1/s1/e2/resources/stars?page=%d&sort-by=id&sort-direction=desc", page)

// 	client := &http.Client{}

// 	// Crear el request
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		fmt.Printf("Error creando el request: %v\n", err)
// 		return nil
// 	}

// 	// Añadir headers
// 	req.Header.Add("API-KEY", "791e97b98aa24e4d8fd52121141e94d8")
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
// 	var resp []ResonanciaResponse
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		fmt.Printf("Error decodificando la respuesta: %v\n", err.Error())
// 		return nil
// 	}
// 	fmt.Println(string(body))
// 	return resp
// }
