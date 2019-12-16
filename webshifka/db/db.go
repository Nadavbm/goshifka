package db

import (
	"fmt"
	"log"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/nadavbm/goshifka/commshifka/env"
)

const (
	host = GetEnv("DB_HOST", "shifkadb")
	port = GetEnvInt("DB_PORT", 5432)
	user = GetEnv("DB_USER", "")
	password = GetEnv("DB_PASSWORD", "")
	dbname = GetEnv("DB_NAME", "")
)

