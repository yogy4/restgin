# Untuk pengoperasian(studi kasus dalam hal ini menggunakan OS Arch Linux, database PostgreSQL 11.5, bahasa pemrograman Go versi go1.13.1 serta pengujian menggunakan tool Postman) dan berikut beberapa langkahnya

1. Install package-package yang berkaitan(dalam mode root), ketikan di terminal
pacman -S go postgresql \
yaourt postman(kemudian pilih angka untuk menentukan versi) \
2. Konfigurasi package-package
* untuk golang \
setup gopath(buat direktori, ex: mkdir -p ~/go/src; export path, ex: export PATH="$PATH:$HOME/go/bin") \
install library pendukung(posisi di path folder, ex: go/src) \
go get -u github.com/gin-gonic/gin \
go get -u github.com/jinzhu/gorm \
go get -u github.com/dgrijalva/jwt-go \
go get -u golang.org/x/crypto/bcrypt \
* untuk postgresql \
setup database(masuk ke postgres, ex: sudo -i -u postgres; buat account baru, ex: createuser interactive --pwprompt; lalu buat database, ex: createdb -O user dbname) \
3. Setup aplikasi(agar bisa jalan, folder aplikasi ditaruh di path folder; ubah konfigurasi database di basecon/conf.go pada fungsi init sesuai yang telah dibuat pada step 2; terakhir jalankan aplikasi, ex: go run coba.go atau go build; ./hasil_compile, ex: ./restgin)
4. Alamat url berada di file coba.go
