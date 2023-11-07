package angga

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

//const (
//	host     = "localhost"
//	port     = 5431
//	user     = "postgres"
//	password = "password"
//	dbname   = "movies_db"
//)

func connectPostgresDB() *sql.DB {
	connstring := "user=postgres dbname=postgres password='password' host=localhost port=5431 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func list() {
	fmt.Println("=-=-=-=----- Film ----=-=-=-=")
	fmt.Println("1. Tambahkan Film")
	fmt.Println("2. Lihat Daftar Film")
	fmt.Println("3. Ubah Data Film")
	fmt.Println("4. Hapus Film,")
	fmt.Println("5. Keluar")
}

func insert(db *sql.DB, judul string, sutradara string, negara string, tahun int, gendre string) {
	_, err := db.Exec("INSERT INTO movies(judul,sutradara,negara,tahun,gendre) VALUES($1,$2,$3,$4,$5)", judul, sutradara, negara, tahun, gendre)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value inserted")
	}
}

func add(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Judul : ")
	str1, _ := reader.ReadString('\n')
	judul := strings.Trim(strings.Trim(str1, "\n"), " ")

	fmt.Print("Sutradara : ")
	str2, _ := reader.ReadString('\n')
	sutradara := strings.Trim(strings.Trim(str2, "\n"), " ")

	fmt.Print("negara : ")
	str3, _ := reader.ReadString('\n')
	negara := strings.Trim(strings.Trim(str3, "\n"), " ")

	fmt.Print("tahun : ")
	var tahun int
	fmt.Scan(&tahun)

	fmt.Print("gendre : ")
	str4, _ := reader.ReadString('\n')
	gendre := strings.Trim(strings.Trim(str4, "\n"), " ")

	var addMovie movie
	addMovie.judul = judul
	addMovie.sutradara = sutradara
	addMovie.negara = negara
	addMovie.tahun = tahun
	addMovie.gendre = gendre

	insert(db, addMovie.judul, addMovie.sutradara, addMovie.negara, addMovie.tahun, addMovie.gendre)

}

func View(rows *sql.Rows, err error) {
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer rows.Close()

	var result []movie

	for rows.Next() {
		var each = movie{}
		var err = rows.Scan(&each.id, &each.judul, &each.sutradara, &each.negara, &each.tahun, &each.gendre)

		if err != nil {
			fmt.Print(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Print(err.Error())
		return
	}

	for i := 0; i < len(result); i++ {
		fmt.Println(i+1, ".", "Judul : ", result[i].judul, " | Sutradara : ", result[i].sutradara, " | Negara : ", result[i].negara, " | Tahun : ", result[i].tahun, " | Gendre : ", result[i].gendre)
	}
}

func update(db *sql.DB) {
	fmt.Print("Judul Yang ingin diganti : ")
	var judul string
	fmt.Scan(&judul)

	fmt.Print("Judul : ")
	var judulPengganti string
	fmt.Scan(&judulPengganti)

	fmt.Print("Sutradara : ")
	var sutradaraPengganti string
	fmt.Scan(&sutradaraPengganti)

	fmt.Print("negara : ")
	var negaraPengganti string
	fmt.Scan(&negaraPengganti)

	fmt.Print("tahun : ")
	var tahunPengganti int
	fmt.Scan(&tahunPengganti)

	fmt.Print("gendre : ")
	var gendrePengganti string
	fmt.Scan(&gendrePengganti)

	_, err := db.Exec("UPDATE movies SET judul=$1, sutradara=$2, negara=$3, tahun=$4, gendre=$5 WHERE judul=$6", judulPengganti, sutradaraPengganti, negaraPengganti, tahunPengganti, gendrePengganti, judul)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data updated")
	}

}

func delete(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Judul Yang ingin dihapus : ")
	judul, _ := reader.ReadString('\n')

	_, err := db.Exec("DELETE FROM MOVIES WHERE JUDUL=$1", judul)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil dihapus !")
	}
}

type movie struct {
	id        int
	judul     string
	sutradara string
	negara    string
	tahun     int
	gendre    string
}

func Main() {

	db := connectPostgresDB()

	for {

		list()

		var inp int
		fmt.Print("Pilih Aksi : ")
		fmt.Scan(&inp)

		switch inp {
		case 1:
			add(db)
		case 2:
			var optView int
			fmt.Println("Opsi Pencarian :")
			fmt.Println("1. Semua Data")
			fmt.Println("2. Judul")
			fmt.Println("3. Tahun")
			fmt.Println("4. Gendre")
			fmt.Print("Opsi : ")
			fmt.Scan(&optView)

			switch optView {
			case 1:
				rows, err := db.Query("SELECT * FROM movies")
				View(rows, err)
			case 2:
				var judul string
				fmt.Print("Masukan Judul : ")
				fmt.Scan(&judul)

				rows, err := db.Query("SELECT * FROM movies WHERE JUDUL = $1", judul)
				View(rows, err)
			case 3:
				var tahun string
				fmt.Print("Masukan Tahun : ")
				fmt.Scan(&tahun)

				rows, err := db.Query("SELECT * FROM movies WHERE TAHUN LIKE '%' || $1 || '%'", tahun)
				View(rows, err)
			case 4:
				var gendre string
				fmt.Print("Masukan Gendre : ")
				fmt.Scan(&gendre)

				rows, err := db.Query("SELECT * FROM movies WHERE GENDRE LIKE '%' || $1 || '%'", gendre)
				View(rows, err)
			}
		case 3:
			update(db)
		case 4:
			delete(db)
		case 5:
			os.Exit(0)
		}

	}

}
