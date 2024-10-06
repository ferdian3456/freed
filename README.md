Import Semua Library
go get github.com/go-playground/validator
go get github.com/jackc/pgx/v5/stdlib
go get github.com/google/uuid
go get github.com/julienschmidt/httprouter

Masukkan DB_URI di current directory
export DB_URI=postgresql://ferdian:password@localhost:5432/freed?sslmode=disable

Instal golang migrate buat migrasi database
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
