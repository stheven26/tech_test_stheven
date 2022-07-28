package main

import (
	"fmt"
)

// soal 1
var (
	seratusRibu     float64 = 100000
	limaPuluhRibu   float64 = 50000
	duaPuluhRibu    float64 = 20000
	sepuluhRibu     float64 = 10000
	limaRibu        float64 = 5000
	duaRibu         float64 = 2000
	seribu          float64 = 1000
	limaRatusRupiah float64 = 500
	duaRatusRupiah  float64 = 200
	seratusRupiah   float64 = 100
)

func PecahanMataUang(num float64) {
	counterSeratusRibu := 0
	counterLimaPuluhRibu := 0
	counterDuaPuluhRibu := 0
	counterSepuluhRibu := 0
	counterLimaRibu := 0
	counterDuaRibu := 0
	counterSeribu := 0
	counterLimaRatusRupiah := 0
	counterDuaRatusRupiah := 0
	counterSeratusRupiah := 0

	for ; num-seratusRibu >= 0; num = num - seratusRibu {
		counterSeratusRibu++
	}
	for ; num-limaPuluhRibu >= 0; num -= limaPuluhRibu {
		counterLimaPuluhRibu++
	}
	for ; num-duaPuluhRibu >= 0; num -= duaPuluhRibu {
		counterDuaPuluhRibu++
	}
	for ; num-sepuluhRibu >= 0; num -= sepuluhRibu {
		counterSepuluhRibu++
	}
	for ; num-limaRibu >= 0; num -= limaRibu {
		counterLimaRibu++
	}
	for ; num-duaRibu >= 0; num -= duaRibu {
		counterDuaRibu++
	}
	for ; num-seribu >= 0; num -= seribu {
		counterSeribu++
	}
	for ; num-limaRatusRupiah >= 0; num -= limaRatusRupiah {
		counterLimaRatusRupiah++
	}
	for ; num-duaRatusRupiah >= 0; num -= duaRatusRupiah {
		counterDuaRatusRupiah++
	}
	for ; num-seratusRupiah >= 0; num -= seratusRupiah {
		counterSeratusRupiah++
	}
	if num > 0 {
		counterSeratusRupiah++
	}

	fmt.Println("Rp. 100.000:", counterSeratusRibu)
	fmt.Println("Rp. 50.000:", counterLimaPuluhRibu)
	fmt.Println("Rp. 20.000:", counterDuaPuluhRibu)
	fmt.Println("Rp. 10.000:", counterSepuluhRibu)
	fmt.Println("Rp. 5.000:", counterLimaRibu)
	fmt.Println("Rp. 2.000:", counterDuaRibu)
	fmt.Println("Rp. 1.000:", counterSeribu)
	fmt.Println("Rp. 500:", counterLimaRatusRupiah)
	fmt.Println("Rp. 200:", counterDuaRatusRupiah)
	fmt.Println("Rp. 100:", counterSeratusRupiah)
}

//soal 2
func insert(str, str2 string) bool {
	if len(str2)-len(str) != -1 {
		return false
	}

	return true
}

func remove(str, str2 string) bool {
	if len(str2)-len(str) > 1 {
		return false
	}

	// selisih := 0
	// i := 0
	// j := 0
	// for i < len(str) && j < len(str2) {
	// 	if str[i] != str2[j] {
	// 		selisih++
	// 		str2 = str2[:j] + str2[j+1:]
	// 	}
	// }

	// if selisih > 1 {
	// 	return false
	// }

	return true
}

func replace(str, str2 string) bool {

	if len(str)-len(str2) != 0 {
		return false
	}

	selisih := 0

	for i := 0; i < len(str); i++ {
		if str[i] != str2[i] {
			selisih++
		}
	}

	if selisih > 1 {
		return false
	}
	return true
}

func Correction(str, str2 string) bool {
	ganti := replace(str, str2)
	hapus := remove(str, str2)
	tambah := insert(str, str2)

	if ganti == true || hapus == true || tambah == true {
		return true
	}

	return false
}

func main() {
	PecahanMataUang(2050)
	soal2 := Correction("telkom", "tlkom")

	fmt.Println("soal 2:", soal2)
}
