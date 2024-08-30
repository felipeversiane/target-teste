package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Revenue struct {
	Day   int     `json:"day"`
	Value float64 `json:"value"`
}

/* Faça a validação do jeito que preferir e do que preferir */
func validate(revenues []Revenue) error {
	for _, r := range revenues {
		if r.Value < 0 {
			return fmt.Errorf("valor de faturamento inválido para o dia %d: %.2f. Os valores devem ser não-negativos.", r.Day, r.Value)
		}
		if r.Day <= 0 {
			return fmt.Errorf("valor de dia inválido: %d. O dia deve ser positivo.", r.Day)
		}
	}
	return nil
}

func main() {
	filePath := "./exec3/revenue.json"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("Arquivo não existe: %v", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	var revenues []Revenue
	if err := json.Unmarshal(byteValue, &revenues); err != nil {
		log.Fatalf("Erro ao descompactar o arquivo: %v", err)
	}

	if err := validate(revenues); err != nil {
		log.Fatalf("Erro de validação: %v", err)
	}

	var sum, min, max float64
	var daysWithRevenue, daysAboveAverage int
	min = -1

	for _, r := range revenues {
		if r.Value > 0 {
			if min == -1 || r.Value < min {
				min = r.Value
			}
			if r.Value > max {
				max = r.Value
			}
			sum += r.Value
			daysWithRevenue++
		}
	}

	if daysWithRevenue == 0 {
		fmt.Println("Nenhum dado de faturamento válido disponível.")
		return
	}

	average := sum / float64(daysWithRevenue)

	var aboveAverageDetails []Revenue
	for _, r := range revenues {
		if r.Value > average {
			aboveAverageDetails = append(aboveAverageDetails, r)
			daysAboveAverage++
		}
	}

	formatBRL := func(amount float64) string {
		return fmt.Sprintf("R$ %.2f", amount)
	}

	fmt.Printf("Menor faturamento: %s\n", formatBRL(min))
	fmt.Printf("Maior faturamento: %s\n", formatBRL(max))
	fmt.Printf("Número de dias com faturamento acima da média: %d\n", daysAboveAverage)

	fmt.Println("Dias com faturamento acima da média:")
	for _, r := range aboveAverageDetails {
		fmt.Printf("Dia %d: %s\n", r.Day, formatBRL(r.Value))
	}
}
