package main

// Normalisasi database dalam bentuk 3NF bertujuan untuk menghilangkan seluruh atribut atau field yang tidak berhubungan dengan primary key.
// Dengan demikian tidak ada ketergantungan transitif pada setiap kandidat key
// fungsi normalisasi 3NF antara lain :
// 1. Memenuhi semua persyaratan dari bentuk normal kedua.
// 2. Menghapus kolom yang tidak tergantung pada primary key.

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type Rekap struct {
	NoBon     int
	Discount  int
	Total     int
	Bayar     int
	Kembalian int
	KodeKasir string
	Tanggal   string
	Waktu     string
}

type DetailRekap struct {
	NoBon      int
	KodeBarang string
	Harga      int
	Jumlah     int
}

type Barang struct {
	KodeBarang string
	NamaBarang string
	Harga      int
}

type Kasir struct {
	KodeKasir string
	NamaKasir string
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate:
// Buatlah tabel dengan nama rekap, rekap_detail, barang, dan kasir
// Lalu insert data ke masing-masing tabel seperti pada contoh di bagian bawah file ini
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	//sqlStmt := `CREATE TABLE rekap ... ` // TODO: replace this
	sqlStmt := `CREATE TABLE rekap (
		no_bon INTEGER,
		discount INTEGER,
		total INTEGER,
		bayar INTEGER,
		kembalian INTEGER,
		kode_kasir VARCHAR(20),
		tanggal VARCHAR(20),
		waktu VARCHAR(20)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	//_, err = db.Exec(`INSERT INTO rekap ... `) // TODO: replace this
	_, err = db.Exec(`INSERT INTO rekap (no_bon, discount, total, bayar, kembalian, kode_kasir, tanggal, waktu) VALUES 
		(00001, 0, 100000, 100000, 0, "K001", "2020-01-01", "12:00:00"),
		(00001, 0, 10000, 10000, 0, "K001", "2020-01-01", "12:00:00"),
		(00001, 0, 70000, 100000, 30000, "K001", "2020-01-01", "12:00:00"),
		(00002, 0, 1000, 10000, 9000, "K002", "2020-01-01", "12:00:00"),
		(00002, 0, 50000, 100000, 50000, "K002", "2020-01-01", "12:00:00")
	;`)

	if err != nil {
		panic(err)
	}

	//sqlStmt = `CREATE TABLE rekap_detail ... ` // TODO: replace this
	sqlStmt = `CREATE TABLE rekap_detail (
		no_bon INTEGER,
		kode_barang VARCHAR(20),
		harga INTEGER,
		jumlah INTEGER
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	//_, err = db.Exec(`INSERT INTO rekap_detail ... `) // TODO: replace this
	_, err = db.Exec(`INSERT INTO rekap_detail (no_bon, kode_barang, harga, jumlah) VALUES
		(00001, "B001", 10000, 1),
		(00001, "B002", 20000, 1),
		(00001, "B003", 30000, 1),
		(00001, "B004", 40000, 1),
		(00001, "B005", 50000, 1)
	;`)

	if err != nil {
		panic(err)
	}

	//sqlStmt = `CREATE TABLE barang ... ` // TODO: replace this
	sqlStmt = `CREATE TABLE barang (
		kode_barang VARCHAR(20),
		nama_barang VARCHAR(20),
		harga INTEGER
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	//_, err = db.Exec(`INSEERT INTO barang ... `) // TODO: replace this
	_, err = db.Exec(`INSERT INTO barang (kode_barang, nama_barang, harga) VALUES
		("B001", "Kopi", 10000),
		("B002", "Teh", 20000),
		("B003", "Kecap", 30000),
		("B004", "Gula", 40000),
		("B005", "Air Mineral", 50000)
	;`)

	if err != nil {
		panic(err)
	}

	//sqlStmt = `CREATE TABLE kasir ... ` // TODO: replace this
	sqlStmt = `CREATE TABLE kasir (
		kode_kasir VARCHAR(20),
		nama_kasir VARCHAR(20)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	//_, err = db.Exec(`INSERT INTO kasir ... `) // TODO: replace this
	_, err = db.Exec(`INSERT INTO kasir (kode_kasir, nama_kasir) VALUES
		("K001", "Kasir 1"),
		("K002", "Kasir 2"),
		("K003", "Kasir 3")
	;`)

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkNoBonExists:
// checkNoBonExists digunakan untuk menghitung jumlah data yang ada berdasarkan no_bon
func checkNoBonExists(noBon string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	//sqlStmt := `SELECT ... FROM ... WHERE ... = ?;` // TODO: replace this
	sqlStmt := `SELECT COUNT(*) FROM rekap WHERE no_bon = ?;`

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi countRekapDetailByNoBon:
// countRekapDetailByNoBon digunakan untuk menghitung jumlah rekap detail yang ada berdasarkan no_bon
func countRekapDetailByNoBon(noBon string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	//sqlStmt := `SELECT ... FROM ... WHERE ... = ?;` // TODO: replace this
	sqlStmt := `SELECT COUNT(*) FROM rekap_detail WHERE no_bon = ?;`

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return 0, err
	}
	return countBon, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkBarangExists:
func checkBarangExists(kodeBarang string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	//sqlStmt := `...` // TODO: replace this
	sqlStmt := `SELECT COUNT(*) FROM barang WHERE kode_barang = ?;`

	row := db.QueryRow(sqlStmt, kodeBarang)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkKasirExists:
func checkKasirExists(kodeKasir string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	//sqlStmt := `...` // TODO: replace this
	sqlStmt := `SELECT COUNT(*) FROM kasir WHERE kode_kasir = ?;`

	row := db.QueryRow(sqlStmt, kodeKasir)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}
