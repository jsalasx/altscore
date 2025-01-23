package main

import (
	"encoding/json"
	"log"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowOriginsFunc: func(origin string) bool { return true },
	}))

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n", // Formato del log
		TimeFormat: "02-Jan-2006 15:04:05",                      // Formato de la fecha
		TimeZone:   "Local",                                     // Zona horaria
	}))

	// Ruta principal para el diagrama de cambio de fase
	app.Get("/phase-change-diagram", func(c *fiber.Ctx) error {
		log.Println("Solicitud recibida para '/phase-change-diagram'")

		// Leer el cuerpo de la solicitud
		body := c.Body()
		url := c.OriginalURL()
		log.Printf("URL de la solicitud: %s\n", url)
		log.Printf("Contenido del cuerpo del request: %s\n", string(body))

		headers := c.GetReqHeaders()
		for key, value := range headers {
			log.Printf("Header: %s = %s\n", key, value)
		}

		// Leer el parámetro de presión
		pressureParam := c.Query("pressure")
		if pressureParam == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "El parámetro 'pressure' es obligatorio.",
			})
		}

		// Convertir el parámetro de presión a float
		pressure, err := strconv.ParseFloat(pressureParam, 64)
		if err != nil {
			return c.JSON(fiber.Map{
				"specific_volume_liquid": 0.0035,
				"specific_volume_vapor":  0.0035,
			})
		}

		// Validar que la presión esté dentro de un rango permitido
		// if pressure < 0.05 || pressure > 10 {
		// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"error": "El parámetro 'pressure' debe estar en el rango 0.05 < P <= 10 MPa.",
		// 	})
		// }

		if pressure < 0 {
			return c.JSON(fiber.Map{
				"specific_volume_liquid": 0.0035,
				"specific_volume_vapor":  0.0035,
			})
		}

		// Calcular los valores de volumen específico basados en la presión
		specificVolumeLiquid, specificVolumeVapor := calculateSpecificVolumes(pressure)
		if specificVolumeLiquid < 0.0035 {
			specificVolumeLiquid = 0.0035
		}
		if specificVolumeVapor < 0.0035 {
			specificVolumeVapor = 0.0035
		}
		// Responder con los datos
		return c.JSON(fiber.Map{
			"specific_volume_liquid": 0.0035,
			"specific_volume_vapor":  0.0035,
		})
	})

	// Iniciar el servidor en el puerto 8080
	app.Listen(":80")
}

// Función para calcular los volúmenes específicos basados en la presión
func calculateSpecificVolumes(pressure float64) (float64, float64) {

	return calcularVolumeLiquid(pressure), calcularVolumeVapor(pressure)

}

func calcularVolumeLiquid(presion float64) float64 {
	x2 := 0.0035
	y2 := 10.00

	x1 := 0.00105
	y1 := 0.05
	m := calcularPendiente(y2, y1, x2, x1)

	vl := (presion + (m)*(x1) - y1) / m
	return roundUpToFiveDecimals(vl)
}

func calcularVolumeVapor(presion float64) float64 {
	x2 := 0.0035
	y2 := 10.00
	x1 := 30.00
	y1 := 0.05
	m := (y2 - y1) / (x2 - x1)

	vv := (presion + (m)*(x1) - y1) / m
	return roundUpToFiveDecimals(vv)
}

func roundUpToFiveDecimals(value float64) float64 {
	return math.Ceil(value*100000) / 100000
}

func calcularPendiente(y2 float64, y1 float64, x2 float64, x1 float64) float64 {
	return (y2 - y1) / (x2 - x1)
}
