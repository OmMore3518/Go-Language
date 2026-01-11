package read

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFromFile() (float64, error) {
	byteop, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("Balnce file not found")
	}
	stringop := string(byteop)
	ans, err := strconv.ParseFloat(stringop, 64)
	if err != nil {
		return 1000, errors.New("Cannot Converted to float")
	}
	return ans, nil
}

func WriteIntoFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}