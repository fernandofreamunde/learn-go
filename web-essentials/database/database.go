package main

// we need to run: go get -u github.com/go-sql-driver/mysql
import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {

	// Config the db connection... do not forget checking errors
	db, err := sql.Open("mysql", "user:password@(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	// Init the 1st connection to the db, see if works...
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Create a new table
		query := `
			CREATE TABLE users(
				id INT AUTO_INCREMENT,
				username TEXT NOT NULL,
				password TEXT NOT NULL,
				created_at DATETIME,
				PRIMARY KEY (id)
			);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Insert a new user
		username  := "Fernando"
		password  := "supersecret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{ // Query single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query single user
		type user struct{
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users[]user
		for rows.Next() {
			var u user
			
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{ //Deleting user
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
