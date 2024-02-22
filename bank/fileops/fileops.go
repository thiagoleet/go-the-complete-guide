package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Fail to find balance file")
	}

	balanceText := string(data)

	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return defaultValue, errors.New("Fail to parse stored value")
	}

	return balance, nil
}
