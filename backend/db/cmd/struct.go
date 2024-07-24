package cmd

import "sade-backend/db"

type DataTable struct {
	tableName string
	conn      *db.Connection
}
