package sqlte

import (
	"database/sql"
	"log"
	"os"
)

type adaptor struct {
  db *sql.DB
}

func NewAdaptor() *adaptor  {
  os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records. 

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

  sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File,
  if err != nil {
		log.Fatal(err)
	}

  _adaptor := &adaptor{
    db: sqliteDatabase,
  }

  _adaptor.setupTables()

  return _adaptor
}

func (this *adaptor) setupTables() {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		role TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT NOT NULL
	);


	CREATE TABLE IF NOT EXISTS drugs (
		id TEXT PRIMARY KEY,
		name TEXT,
		serial_no TEXT,
		tag_id TEXT,
		verdict TEXT,
		archived BOOLEAN
	);

	CREATE TABLE IF NOT EXISTS convicts (
		drug_id TEXT,
		name TEXT,
		is_arrested BOOLEAN,
		FOREIGN KEY (drug_id) REFERENCES drugs(id)
	);
	`

	log.Println("Create Users table...")
	statement, err := this.db.Prepare(createTableQuery) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec() // Execute SQL Statements
	log.Println("Users table created")
}
