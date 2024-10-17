package main

import (
	"fmt"
	"reflect"
)

type Mobil struct {
	Nama      string
	Kecepatan int
}

func main() {
	// Membuat instance dari struct Mobil
	mobil := Mobil{Nama: "Honda", Kecepatan: 120}

	// Menggunakan reflect untuk mendapatkan nilai dan tipe dari mobil
	val := reflect.ValueOf(&mobil).Elem()

	// Menampilkan nilai field yang ada
	fmt.Printf("Mobil sebelum diubah: %+v\n", mobil)

	// Mengubah nilai field yang ada (contoh: mengubah Kecepatan)
	fieldKecepatan := val.FieldByName("Kecepatan")
	if fieldKecepatan.IsValid() && fieldKecepatan.CanSet() {
		fieldKecepatan.SetInt(150) // mengubah kecepatan menjadi 150
	}

	// Menambahkan field baru ke map (karena kita tidak bisa menambah field baru ke struct)
	fieldName := "Tahun"                   // nama field baru
	fields := make(map[string]interface{}) // map untuk menyimpan field baru

	// Menyimpan field baru ke map
	fields[fieldName] = "2020" // menyimpan nilai tahun

	// Menampilkan mobil setelah diubah
	fmt.Printf("Mobil setelah diubah: %+v\n", mobil)

	// Menampilkan semua field (lama dan baru)
	fmt.Println("\nSemua field:")
	printFields(mobil, fields)

	// Menampilkan field baru yang ditambahkan
	fmt.Println("Field yang ditambahkan:")
	for k, v := range fields {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func printFields(mobil Mobil, fields map[string]interface{}) {
	val := reflect.ValueOf(mobil)

	// Mencetak field dari struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i)

		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}

	// Mencetak field dari map
	for k, v := range fields {
		fmt.Printf("%s: %v\n", k, v)
	}
}
