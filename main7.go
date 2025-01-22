package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var damagedSystem string

// Mapeo de sistemas dañados a sus códigos
var systemCodes = map[string]string{
	"navigation": "NAV-01",
}

func main() {
	// Inicializar el sistema dañado al iniciar el servidor
	systemKeys := []string{"navigation"}

	// Asignar el valor a la variable global
	damagedSystem = systemKeys[0]
	fmt.Printf("Sistema dañado seleccionado: %s\n", damagedSystem)

	http.HandleFunc("/status", handleStatus)
	http.HandleFunc("/repair-bay", handleRepairBay)
	http.HandleFunc("/teapot", handleTeapot)

	fmt.Println("API funcionando en http://localhost:80")
	http.ListenAndServe(":80", nil)
}

// Manejar la llamada GET /status
func handleStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Construir la respuesta JSON
	response := map[string]string{
		"damaged_system": damagedSystem,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Manejar la llamada GET /repair-bay
func handleRepairBay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtener el código asociado al sistema dañado
	code, ok := systemCodes[damagedSystem]
	if !ok {
		http.Error(w, "Sistema desconocido", http.StatusInternalServerError)
		return
	}

	// Generar la página HTML
	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Repair</title>
</head>
<body>
<div class="anchor-point">%s</div>
</body>
</html>`, code)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

// Manejar la llamada POST /teapot
func handleTeapot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Retornar código de estado 418 (I'm a teapot)
	w.WriteHeader(http.StatusTeapot)
	w.Write([]byte("418 I'm a teapot"))
}
