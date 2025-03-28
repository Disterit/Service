package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	password := "123456"
	hasher := sha256.New()

	// Записываем пароль в хеш
	hasher.Write([]byte(password))

	// Получаем хеш в виде байтового среза
	hashBytes := hasher.Sum(nil)

	// Конвертируем байты в hex-строку
	hashString := hex.EncodeToString(hashBytes)

	fmt.Println(hashString)
}
