package main
import (
    "fmt"
    "sort"
    "strconv"
    "time"
)

// Struktur data untuk riwayat tidur
type SleepRecord struct {
    Date       string    // Tanggal tidur
    SleepTime  time.Time // Jam tidur
    WakeTime   time.Time // Jam bangun
    Duration   float64   // Durasi tidur (dalam jam)
    Quality    string    // Kualitas tidur
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

