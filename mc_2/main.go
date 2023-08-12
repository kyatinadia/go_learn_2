package main

import (
	"fmt"
	"os"
	"strings"
)

// Struct untuk menyimpan data tokoh
type Tokoh struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	AlasanAOT string
}

// Slice dari struct tokoh
var tokohAOT = []Tokoh{
	{
		Nama:      "Levi Ackerman",
		Alamat:    "Survey Corps",
		Pekerjaan: "Soldier",
		AlasanAOT: "Humanity's Strongest Soldier",
	},
	{
		Nama:      "Eren Yeager",
		Alamat:    "Survey Corps",
		Pekerjaan: "Titan",
		AlasanAOT: "Founding Titan",
	},
}

// Function untuk menampilkan data tokoh berdasarkan nama yang diinput
func getDataTokoh(nama string) *Tokoh {
	for i := 0; i < len(tokohAOT); i++ {
		if tokohAOT[i].Nama == nama {
			return &tokohAOT[i]
		}
	}
	return nil
}

func main() {
	// Ambil argument nama tokoh dari CLI
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [Nama]")
		return
	}
	//gabung argumen menjadi satu string
	nama := strings.Join(os.Args[1:], " ")

	// Panggil function untuk mendapatkan data tokoh berdasarkan nama
	tokoh := getDataTokoh(nama)
	if tokoh == nil {
		fmt.Println("Tokoh tidak ditemukan.")
		return
	}

	// Tampilkan data tokoh
	fmt.Println("Data Tokoh:")
	fmt.Println("Nama:", tokoh.Nama)
	fmt.Println("Alamat:", tokoh.Alamat)
	fmt.Println("Pekerjaan:", tokoh.Pekerjaan)
	fmt.Println("Karakteristik:", tokoh.AlasanAOT)
}
