package main

import (
	"fmt"
	"strings"
)

func main() {
	var judul string
	fmt.Print("Masukan Judul Film Yang Ingin Dicari : ")
	fmt.Scanln(&judul)

	film, status := namaFilm(judul)

	if status == "Is Found" {
		fmt.Println(film, status)
	} else {
		fmt.Println(film, status)
	}
}

func namaFilm(judul string) (film string, status string) {
	filmList := [...]string{"Money-Heist", "Avengers", "Spiderman", "Black-Mirror"}
	judulToLower := strings.ToLower(judul)
	for _, element := range filmList {
		if strings.ToLower(element) == judulToLower {
			film = judulToLower
			status = "is Found"
			return
		}
	}
	film = judulToLower
	status = "is Not Found"
	return
}
