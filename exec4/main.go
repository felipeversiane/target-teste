package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func calculatePercentage(value, total float64) float64 {
	return (value / total) * 100
}

/* Faça a validação do jeito que preferir e do que preferir */
func validate(revenues map[string]float64) error {
	for state, revenue := range revenues {
		if revenue < 0 {
			return fmt.Errorf("valor de faturamento inválido para o estado %s: R$ %.2f. Os valores devem ser não-negativos.", state, revenue)
		}
	}
	return nil
}

func main() {
	filePath := "./exec4/revenue.json"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	var revenues map[string]float64
	if err := json.Unmarshal(byteValue, &revenues); err != nil {
		log.Fatalf("Erro ao descompactar JSON: %v", err)
	}

	if err := validate(revenues); err != nil {
		log.Fatalf("Erro de validação: %v", err)
	}

	var totalRevenue float64
	for _, revenue := range revenues {
		totalRevenue += revenue
	}

	fmt.Printf("Faturamento Total: R$ %.2f\n", totalRevenue)

	fmt.Println("Percentual de representação por estado:")
	for state, revenue := range revenues {
		percentage := calculatePercentage(revenue, totalRevenue)
		fmt.Printf("%s: %.2f%%\n", state, percentage)
	}
}
