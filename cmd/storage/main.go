package main

import (
	"fmt"
	"github.com/ANkulagin/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("it works", st)
}
