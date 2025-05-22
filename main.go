package main

import (
	"fmt"
	"sort"
	"time"
)

// Struktur data untuk riwayat tidur
type SleepRecord struct {
	Date      string    // Tanggal tidur
	SleepTime time.Time // Jam tidur
	WakeTime  time.Time // Jam bangun
	Duration  float64   // Durasi tidur (dalam jam)
	Quality   string    // Kualitas tidur
}

// Slice untuk menyimpan semua riwayat tidur
var sleepRecords []SleepRecord

// Fungsi untuk menampilkan menu utama
func displayMenu() {
	fmt.Println("=== APLIKASI PEMANTAUAN TIDUR ===")
	fmt.Println("[1] Lihat Riwayat Tidur")
	fmt.Println("[2] Tambahkan Riwayat Tidur Baru")
	fmt.Println("[3] Edit Riwayat Tidur")
	fmt.Println("[4] Hapus Riwayat Tidur")
	fmt.Println("[5] Cari Riwayat Tidur Berdasarkan Tanggal")
	fmt.Println("[6] Urutkan Riwayat Tidur Berdasarkan Durasi atau Tanggal")
	fmt.Println("[7] Laporan Pola Tidur")
	fmt.Println("[0] Keluar")
	fmt.Print("Pilih menu> ")
}

// Fungsi untuk menambahkan riwayat tidur baru
func addSleepRecord() {
	var date string
	var sleepTimeString, wakeTimeString string
	var quality string

	fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
	fmt.Scanln(&date)

	fmt.Print("Masukkan waktu tidur (HH:MM): ")
	fmt.Scanln(&sleepTimeString)
	sleepTime, _ := time.Parse("15:04", sleepTimeString)

	fmt.Print("Masukkan waktu bangun (HH:MM): ")
	fmt.Scanln(&wakeTimeString)
	wakeTime, _ := time.Parse("15:04", wakeTimeString)

	duration := wakeTime.Sub(sleepTime).Hours()
	fmt.Print("Masukkan kualitas tidur (baik/sedang/jelek): ")
	fmt.Scanln(&quality)

	newRecord := SleepRecord{
		Date:      date,
		SleepTime: sleepTime,
		WakeTime:  wakeTime,
		Duration:  duration,
		Quality:   quality,
	}
	sleepRecords = append(sleepRecords, newRecord)
	fmt.Println("Riwayat tidur berhasil ditambahkan!")
}

// Fungsi untuk menampilkan semua riwayat tidur
func displaySleepRecords() {
	if len(sleepRecords) == 0 {
		fmt.Println("Tidak ada riwayat tidur.")
		return
	}
	fmt.Println("=== RIWAYAT TIDUR ===")
	for i, record := range sleepRecords {
		fmt.Printf("[%d] Tanggal: %s, Tidur: %s, Bangun: %s, Durasi: %.2f jam, Kualitas: %s\n",
			i+1, record.Date, record.SleepTime.Format("15:04"), record.WakeTime.Format("15:04"),
			record.Duration, record.Quality)
	}
}

// Fungsi untuk mengedit riwayat tidur
func editSleepRecord() {
	displaySleepRecords()
	if len(sleepRecords) == 0 {
		return
	}

	var index int
	fmt.Print("Pilih nomor riwayat yang ingin diedit: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(sleepRecords) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	var date string
	var sleepTimeString, wakeTimeString string
	var quality string

	fmt.Print("Masukkan tanggal baru (YYYY-MM-DD): ")
	fmt.Scanln(&date)

	fmt.Print("Masukkan waktu tidur baru (HH:MM): ")
	fmt.Scanln(&sleepTimeString)
	sleepTime, _ := time.Parse("15:04", sleepTimeString)

	fmt.Print("Masukkan waktu bangun baru (HH:MM): ")
	fmt.Scanln(&wakeTimeString)
	wakeTime, _ := time.Parse("15:04", wakeTimeString)

	duration := wakeTime.Sub(sleepTime).Hours()
	fmt.Print("Masukkan kualitas tidur baru (baik/sedang/jelek): ")
	fmt.Scanln(&quality)

	sleepRecords[index-1] = SleepRecord{
		Date:      date,
		SleepTime: sleepTime,
		WakeTime:  wakeTime,
		Duration:  duration,
		Quality:   quality,
	}
	fmt.Println("Riwayat tidur berhasil diperbarui!")
}

// Fungsi untuk menghapus riwayat tidur
func deleteSleepRecord() {
	displaySleepRecords()
	if len(sleepRecords) == 0 {
		return
	}

	var index int
	fmt.Print("Pilih nomor riwayat yang ingin dihapus: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(sleepRecords) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	sleepRecords = append(sleepRecords[:index-1], sleepRecords[index:]...)
	fmt.Println("Riwayat tidur berhasil dihapus!")
}

// Fungsi untuk mencari riwayat tidur berdasarkan tanggal
func searchSleepRecordByDate(date string) {
	found := false
	for _, record := range sleepRecords {
		if record.Date == date {
			fmt.Printf("Tanggal: %s, Tidur: %s, Bangun: %s, Durasi: %.2f jam, Kualitas: %s\n",
				record.Date, record.SleepTime.Format("15:04"), record.WakeTime.Format("15:04"),
				record.Duration, record.Quality)
			found = true
		}
	}
	if !found {
		fmt.Println("Riwayat tidur tidak ditemukan.")
	}
}

// Fungsi untuk mengurutkan riwayat tidur berdasarkan durasi atau tanggal
func sortSleepRecords() {
	fmt.Println("Pilih cara pengurutan:")
	fmt.Println("[1] Urutkan berdasarkan durasi tidur")
	fmt.Println("[2] Urutkan berdasarkan tanggal")
	var choice int
	fmt.Print("Pilihan> ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		sort.Slice(sleepRecords, func(i, j int) bool {
			return sleepRecords[i].Duration < sleepRecords[j].Duration
		})
		fmt.Println("Riwayat tidur telah diurutkan berdasarkan durasi.")
	case 2:
		sort.Slice(sleepRecords, func(i, j int) bool {
			return sleepRecords[i].Date < sleepRecords[j].Date
		})
		fmt.Println("Riwayat tidur telah diurutkan berdasarkan tanggal.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Fungsi untuk membuat laporan pola tidur
func generateSleepReport() {
	if len(sleepRecords) == 0 {
		fmt.Println("Tidak ada riwayat tidur untuk dibuat laporan.")
		return
	}

	// Menghitung total durasi tidur dalam 7 hari terakhir
	today := time.Now().Format("2006-01-02")
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02")

	totalDuration := 0.0
	count := 0
	for _, record := range sleepRecords {
		if record.Date >= sevenDaysAgo && record.Date <= today {
			totalDuration += record.Duration
			count++
		}
	}

	averageDuration := totalDuration / float64(count)
	fmt.Printf("\n=== LAPORAN POLA TIDUR ===\n")
	fmt.Printf("Total durasi tidur dalam 7 hari terakhir: %.2f jam\n", totalDuration)
	fmt.Printf("Rata-rata durasi tidur per minggu: %.2f jam\n", averageDuration)
}

// Main function
func main() {
	for {
		displayMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			displaySleepRecords()
		case 2:
			addSleepRecord()
		case 3:
			editSleepRecord()
		case 4:
			deleteSleepRecord()
		case 5:
			var date string
			fmt.Print("Masukkan tanggal yang ingin dicari (YYYY-MM-DD): ")
			fmt.Scanln(&date)
			searchSleepRecordByDate(date)
		case 6:
			sortSleepRecords()
		case 7:
			generateSleepReport()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi! Sampai jumpa.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

		fmt.Println("\nTekan Enter untuk melanjutkan...")
		fmt.Scanln()
	}
}
