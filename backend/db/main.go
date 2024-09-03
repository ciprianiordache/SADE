package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func New(server, user, password, database, driver string) *Connection {
	return &Connection{
		Server:   server,
		User:     user,
		Pass:     password,
		DataBase: database,
		Driver:   driver,
	}
}

func (db *Connection) Conn() error {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", db.Server, db.User, db.Pass, db.DataBase)
	conn, err := sql.Open(db.Driver, connStr)
	if err != nil {
		return err
	}
	db.conn = conn
	return nil
}

func (db *Connection) Exec(sql string, params ...interface{}) ([]map[string]interface{}, error) {
	if db.conn == nil {
		return nil, fmt.Errorf("db connection is not established")
	}
	stmt, err := db.conn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuesPtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuesPtrs[i] = &values[i]
		}
		if err := rows.Scan(valuesPtrs...); err != nil {
			return nil, err
		}
		rowData := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				rowData[col] = string(b)
			} else {
				rowData[col] = val
			}
		}
		result = append(result, rowData)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
func (db *Connection) CreateTable() {
	err := db.Conn()
	if err != nil {
		fmt.Println("cannot connect to database: ", err)
	}

	createUserTable := `CREATE TABLE IF NOT EXISTS users (
    						id SERIAL PRIMARY KEY,
    						first_name TEXT,
    						last_name TEXT,
    						email TEXT UNIQUE NOT NULL,
	  					password TEXT,
    						role TEXT NOT NULL,
    						verified BOOLEAN NOT NULL DEFAULT FALSE,
    						created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
						);`

	_, err = db.Exec(createUserTable)
	if err != nil {
		fmt.Printf("cannot create users table: %v", err)
	}

	createLinkTable := `CREATE TABLE IF NOT EXISTS links (
    						id SERIAL PRIMARY KEY,
    						email TEXT NOT NULL,
    						link TEXT NOT NULL,
    						expiry INTEGER NOT NULL,
    						created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    						FOREIGN KEY (email) REFERENCES users(email)
						);`

	_, err = db.Exec(createLinkTable)
	if err != nil {
		fmt.Println("cannot create links table: ", err)
	}

	createMediaTable := `CREATE TABLE IF NOT EXISTS media (
    						id SERIAL PRIMARY KEY,
    						uploaded_by INTEGER NOT NULL,
    						client_email TEXT NOT NULL,
    						preview_path TEXT NOT NULL,
    						original_path TEXT NOT NULL,
    						locked BOOLEAN NOT NULL DEFAULT TRUE,
    						price REAL NOT NULL DEFAULT 0.0,
    						created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    						FOREIGN KEY (client_email) REFERENCES users(email),
    						FOREIGN KEY (uploaded_by) REFERENCES users(id)
						);`

	_, err = db.Exec(createMediaTable)
	if err != nil {
		fmt.Println("cannot create media table: ", err)
	}

	createTransactionTable := `CREATE TABLE IF NOT EXISTS transaction (
								  id SERIAL PRIMARY KEY,
								  admin_id INTEGER NOT NULL,
                        		  media_id INTEGER NOT NULL,
                        		  client_email TEXT NOT NULL,
                        		  amount BIGINT NOT NULL,
                        		  currency TEXT NOT NULL,
                        		  description TEXT NOT NULL,
                        		  transaction_id TEXT NOT NULL,
                        		  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        		  FOREIGN KEY (media_id) REFERENCES media(id),
    							  FOREIGN KEY (admin_id) REFERENCES users(id)
								);`
	_, err = db.Exec(createTransactionTable)
	if err != nil {
		fmt.Println("cannot create transaction table: ", err)
	}
}
