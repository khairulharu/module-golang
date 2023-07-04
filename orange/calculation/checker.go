package calculation

func CheckNilai(nilai int) string {
	if nilai >= 90 && nilai <= 100 {
		return "Grade A"
	} else if nilai >= 80 && nilai <= 89 {
		return "Grade B"
	} else if nilai >= 75 && nilai <= 79 {
		return "Grade C"
	} else {
		return "Grade Anda D atau inputan yang anda masukkan salah"
	}
}
