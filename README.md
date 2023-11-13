# ijoin

berisikan Bot untuk melakukan Penghijauan Repository

     ```powershell
     name: CI

        on:
        push:
            branches:
            - master
        workflow_dispatch:
        schedule:
            - cron: "0 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20 * * *"
        pull_request:
            branches:
            - master

        jobs:
        build:
            runs-on: ubuntu-latest

            steps:
            - name: Checkout repository
                uses: actions/checkout@v2

            - name: Pull latest changes
                run: git pull

            - name: Create and populate random file
                run: |
                # Generate a random 9-digit number
                random_number=$((100000000 + RANDOM % 900000000))

                # Create a Go file with the random number as the name
                filename="${random_number}.go"

                # Write the Go code to the file
                echo 'package main

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

                    fmt.Println("Output telah ditulis ke output.txt")' > "$filename"

                echo "Created and populated file: $filename"
                working-directory: ${{ github.workspace }}

            - name: Commit and push random file
                run: |
                git add .
                git config user.name 'rezapace'
                git config user.email 'm.rezahidayat.rh@gmail.com'
                git commit -m "Add random Go file"
                git push
     ```
