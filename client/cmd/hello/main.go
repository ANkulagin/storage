package main

import (
	"fmt"
	"github.com/ANkulagin/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("client: ", file)
}
