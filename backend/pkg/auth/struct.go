package auth

import (
	"sade-backend/db/cmd"
	"time"
)

type Auth struct {
	userTable *cmd.DataTable
	linkTable *cmd.DataTable
	timeout   time.Duration
}
