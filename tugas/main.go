package main

import (
	"fmt"
	"log"

	"github.com/golang-beginer-4/tugas/model"
)

func main() {
	var accounts []model.Account
	account := model.Account{}

	for {
		fmt.Println("\nMasukan angka untuk memilih:")
		fmt.Println("1. Tambah akun")
		fmt.Println("2. Tambah saldo")
		fmt.Println("3. Kurangi saldo")
		fmt.Println("4. Lihat semua akun")
		fmt.Println("5. Keluar")
		var pilihan int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var id, name, email string
			var saldo float64

			fmt.Println("\nMasukkan informasi akun baru:")
			fmt.Print("ID: ")
			fmt.Scan(&id)
			fmt.Print("Nama: ")
			fmt.Scan(&name)
			fmt.Print("Email: ")
			fmt.Scan(&email)
			fmt.Print("Saldo awal: ")
			fmt.Scan(&saldo)

			newAccount, err := account.CreateUser(id, name, email, saldo)
			if err != nil {
				log.Fatalf("Gagal membuat akun: %v", err)
			}

			fmt.Println("Akun berhasil dibuat:", newAccount)

		case 2:
			var id string
			var jumlah float64
			fmt.Print("Masukkan id akun untuk menambah saldo: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan jumlah saldo yang ingin ditambahkan: ")
			fmt.Scan(&jumlah)

			for i := range accounts {
				if accounts[i].Id == id {
					accounts[i].TambahSaldo(jumlah)
					fmt.Printf("Saldo berhasil ditambahkan. Saldo sekarang: %.2f\n", accounts[i].Saldo.Amount)
					break
				}
			}

		case 3:
			var id string
			var jumlah float64
			fmt.Print("Masukkan id akun untuk mengurangi saldo: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan jumlah saldo yang ingin dikurangi: ")
			fmt.Scan(&jumlah)

			for i := range accounts {
				if accounts[i].Id == id {
					accounts[i].KurangiSaldo(jumlah)
					fmt.Printf("Saldo sekarang: %.2f\n", accounts[i].Saldo.Amount)
					break
				}
			}

		case 4:
			fmt.Println("\nDaftar akun:")
			for _, user := range accounts {
				fmt.Printf("User: %+v\n", user)
			}

		case 5:
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
