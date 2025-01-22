package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {

	numRequest := 1000
	var msgArray []string
	for i := 0; i < numRequest; i++ {
		msgArray = append(msgArray, getMsgs())

	}
	//fmt.Println(msgArray)

}

func getMsgs() string {
	url := "https://makers-challenge.altscore.ai/v1/s1/e8/actions/door"
	client := &http.Client{}

	// Crear el request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error creando el request: %v\n", err)
		return ""
	}

	// Añadir headers
	req.Header.Add("API-KEY", "791e97b98aa24e4d8fd52121141e94d8")
	req.Header.Add("Accept", "application/json")

	// Hacer el request
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error haciendo GET request: %v\n", err)
		return ""
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error leyendo el body: %v\n", err)
		return ""
	}

	// imprimir headers

	bodyMsg := string(body)
	//fmt.Println(string(body))

	if response.StatusCode == 200 {
		fmt.Printf("\n\n")
		fmt.Println(string(body))
		for name, values := range response.Header {
			// Loop over all values for the name.
			for _, value := range values {
				// if name != "Content-Type" && name != "Content-Length" && name != "Access-Control-Allow-Origin" && name != "Access-Control-Allow-Methods" && name != "Access-Control-Allow-Headers" && name != "Access-Control-Max-Age" {
				// 	fmt.Println(name, value)
				// }

				if name == "Set-Cookie" {
					fmt.Println(name, value)
					cockieValue := getCockieValue(value)
					return getMsgsCockie(cockieValue)
				}

			}
		}
	}

	if response.StatusCode != 200 {
		fmt.Printf("\n\n")
		fmt.Println("Error en la respuesta")
		fmt.Println(string(body))
		for name, values := range response.Header {
			// Loop over all values for the name.
			for _, value := range values {
				if name != "Content-Type" && name != "Content-Length" && name != "Access-Control-Allow-Origin" && name != "Access-Control-Allow-Methods" && name != "Access-Control-Allow-Headers" && name != "Access-Control-Max-Age" {
					fmt.Println(name, value)
				}
			}
		}
	}

	return bodyMsg

}

func getMsgsCockie(cockie string) string {
	url := "https://makers-challenge.altscore.ai/v1/s1/e8/actions/door"
	client := &http.Client{}

	// Crear el request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error creando el request: %v\n", err)
		return ""
	}

	// Añadir headers
	req.Header.Add("API-KEY", "791e97b98aa24e4d8fd52121141e94d8")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cookie", "gryffindor="+cockie)

	// Hacer el request
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error haciendo GET request: %v\n", err)
		return ""
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error leyendo el body: %v\n", err)
		return ""
	}

	// imprimir headers

	bodyMsg := string(body)
	//fmt.Println(string(body))

	if response.StatusCode == 200 {
		fmt.Printf("\n\n")
		fmt.Println(string(body))
		for name, values := range response.Header {
			// Loop over all values for the name.
			for _, value := range values {
				if name == "Set-Cookie" {
					fmt.Println(name, value)
					cockieValue := getCockieValue2(value)
					return getMsgsCockie(cockieValue)
				}

			}
		}
	}

	if response.StatusCode != 200 {
		fmt.Printf("\n\n")
		fmt.Println("Error en la respuesta")
		fmt.Println(string(body))
		for name, values := range response.Header {
			// Loop over all values for the name.
			for _, value := range values {
				if name != "Content-Type" && name != "Content-Length" && name != "Access-Control-Allow-Origin" && name != "Access-Control-Allow-Methods" && name != "Access-Control-Allow-Headers" && name != "Access-Control-Max-Age" {
					fmt.Println(name, value)
				}
			}
		}
	}

	return bodyMsg

}

func getCockieValue(in string) string {
	regex := regexp.MustCompile(`gryffindor="([^"]+)"`)
	match := regex.FindStringSubmatch(in)

	if len(match) > 1 {
		fmt.Println("Cadena original:", match)
		decodedBytes, err := base64.StdEncoding.DecodeString(match[1])
		if err != nil {
			fmt.Println("Error al decodificar:", err)
			return match[1]
		}

		// Convertir los bytes decodificados a una cadena
		decodedString := string(decodedBytes)
		fmt.Println("Cadena decodificada:", decodedString)
		return match[1]
	} else {
		fmt.Println("No se encontró el valor.")
		return ""
	}

}

func getCockieValue2(in string) string {
	regex := regexp.MustCompile(`=(.*?);`)
	match := regex.FindStringSubmatch(in)

	if len(match) > 1 {
		fmt.Println("Cadena original:", match)
		decodedBytes, err := base64.StdEncoding.DecodeString(match[1])
		if err != nil {
			fmt.Println("Error al decodificar:", err)
			return match[1]
		}

		// Convertir los bytes decodificados a una cadena
		decodedString := string(decodedBytes)
		fmt.Println("Cadena decodificada:", decodedString)
		return match[1]
	} else {
		fmt.Println("No se encontró el valor.")
		return ""
	}

}
