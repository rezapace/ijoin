package main

import (
  "fmt"
  "os"
)

func calculateTip(tagihan float64) (float64, float64) {
  var tipRate float64

  if tagihan >= 50 && tagihan <= 300 {
    tipRate = 0.15
  } else {
    tipRate = 0.20
  }

  tip := tagihan * tipRate
  total := tagihan + tip

  return tip, total
}

func main() {
  var tagihan float64
  fmt.Print("Masukkan nilai tagihan: ")
  fmt.Scan(&tagihan)

  tip, total := calculateTip(tagihan)

  outputText := fmt.Sprintf("Tagihannya: Rp. %.2f\nTipnya: Rp. %.2f\nTotal nilainya: Rp. %.2f\n", tagihan, tip, total)

  // Membuka file output.txt untuk penulisan
  file, err := os.Create("output.txt")
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  defer file.Close()

  // Menulis output ke file
  _, err = file.WriteString(outputText)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  fmt.Println("Output telah ditulis ke output.txt")
