package utils

import (
    "math/rand"
    "time"
)

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
  
func GetRandomUrl(size int) string {
  
    rand.Seed(time.Now().Unix())
  
    str := make([]rune, size)
  
    // Generating Random string
    for i := range str {
        str[i] = runes[rand.Intn(len(runes))]
    }
  
    return string(str)
}