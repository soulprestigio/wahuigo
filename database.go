package database

import (
 "database/sql"
 _ "github.com/mattn/go-sqlite3"
 "log"
 "os" 
)

// Declare db here
var db *sql.DB 

// GetDB returns the database connection.
func GetDB() *sql.DB {
 return db
}

func CloseDB() {
 if db != nil {
  db.Close()
 }
}


// InitializeDB initializes the database connection.
func InitializeDB() { 
 var err error
 // Check if the database file exists
 if _, err := os.Stat("./wahui.db"); os.IsNotExist(err) {
  db, err = sql.Open("sqlite3", "./wahui.db")
  if err != nil {
   log.Fatal(err)
  }
 } else {
  db, err = sql.Open("sqlite3", "./wahui.db")
  if err != nil {
   log.Fatal(err)
  }
 }

 // Create the boards table if it doesn't exist
 _, err := db.Exec(`
  CREATE TABLE IF NOT EXISTS boards (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name TEXT NOT NULL,
   description TEXT NOT NULL,
   url TEXT NOT NULL
  );
 `)
 if err != nil {
  log.Fatal(err)
 }
 

}
