package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"sync"
// )

// type SondaSilenciosaResponse struct {
// 	Distance string `json:"distance"`
// 	Time     string `json:"time"`
// }

// type SondaSilenciosaData struct {
// 	Distance float64 `json:"distance"`
// 	Time     float64 `json:"time"`
// }

// func main() {
// 	numRequests := 10000 // Total de solicitudes
// 	numGoroutines := 100 // Número de gorutinas
// 	responses := make([]*SondaSilenciosaData, 0, numRequests)
// 	responseChannel := make(chan *SondaSilenciosaData, numRequests)

// 	var wg sync.WaitGroup

// 	for i := 0; i < numGoroutines; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for j := 0; j < numRequests/numGoroutines; j++ {
// 				response := getData()
// 				if response != nil {
// 					if response.Distance == "failed to measure, try again" {
// 						continue
// 					}

// 					if response.Time == "failed to measure, try again" {
// 						continue
// 					}
// 					distanceStr := strings.Replace(response.Distance, " AU", "", 1)
// 					distance, err := strconv.ParseFloat(distanceStr, 64)
// 					if err != nil {
// 						fmt.Printf("Error al convertir distancia: %v\n", err)
// 						return
// 					}

// 					timeStr := strings.Replace(response.Time, " hours", "", 1)
// 					time, err := strconv.ParseFloat(timeStr, 64)
// 					if err != nil {
// 						fmt.Printf("Error al convertir tiempo: %v\n", err)
// 						return
// 					}
// 					dataAux := &SondaSilenciosaData{
// 						Distance: distance, // Convertir AU a km
// 						Time:     time,     // Convertir horas a segundos
// 					}
// 					responseChannel <- dataAux
// 				}
// 			}
// 		}()
// 	}

// 	// Cerramos el canal una vez que todas las gorutinas terminen
// 	go func() {
// 		wg.Wait()
// 		close(responseChannel)
// 	}()

// 	// Recopilamos las respuestas del canal
// 	for resp := range responseChannel {
// 		responses = append(responses, resp)
// 	}

// 	fmt.Printf("Se obtuvieron %d respuestas\n", len(responses))
// 	velocidad := calcularVelocidad(responses)
// 	fmt.Println(velocidad)
// }

// func getData() *SondaSilenciosaResponse {
// 	url := "https://makers-challenge.altscore.ai/v1/s1/e1/resources/measurement"

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
// 	var resp SondaSilenciosaResponse
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		fmt.Printf("Error decodificando la respuesta: %v\n", err)
// 		return nil
// 	}
// 	fmt.Println(string(body))
// 	return &resp
// }

// func calcularVelocidad(datos []*SondaSilenciosaData) string {
// 	var distanciaTotal float64
// 	var tiempoTotal float64
// 	for _, dato := range datos {
// 		if dato != nil {
// 			distanciaTotal += dato.Distance
// 			tiempoTotal += dato.Time
// 		}
// 	}

// 	distanciaPromedio := distanciaTotal / float64(len(datos))
// 	tiempoPromedio := tiempoTotal / float64(len(datos))
// 	VelocidadEnAU := distanciaPromedio / tiempoPromedio
// 	fmt.Println("Velocidad en AU: ", VelocidadEnAU)
// 	distanciaPromedioEnKms := distanciaPromedio * 149597870.7
// 	tiempoPromedioEnSegs := tiempoPromedio * 3600

// 	velocidad := distanciaPromedioEnKms / tiempoPromedioEnSegs

// 	velocidadEnKmsPorHora := velocidad / 3600
// 	fmt.Println("Distancia promedio en kms: ", velocidadEnKmsPorHora)

// 	return fmt.Sprintf("La velocidad promedio es: %f km/s", velocidad)

// }
