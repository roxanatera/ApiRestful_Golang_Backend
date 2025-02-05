package connect

import (
	"context"
	"fmt"
	"os"

	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("La variable de entorno %s no está configurada", key))
	}
	return value
}

func Connect() rel.Repository {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}

	user := getEnv("DB_USER")
	pass := getEnv("DB_PASSWORD")
	server := getEnv("DB_SERVER")
	port := getEnv("DB_PORT")
	name := getEnv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?clientFoundRows=true&charset=utf8&parseTime=True&loc=Local",
		user, pass, server, port, name,
	)

	adapter, err := mysql.Open(dsn)
	if err != nil {
		panic(err)
	}

	// Prueba la conexión
	ctx := context.Background()
	if err := adapter.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Conexión exitosa a la base de datos")

	repo := rel.New(adapter)
	return repo
}