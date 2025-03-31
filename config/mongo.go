package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var Collection *mongo.Collection

func ConectarMongo() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}

	DB = client.Database(dbName)
	Collection = DB.Collection(collectionName)

	fmt.Println("✅ Conectado a MongoDB correctamente.")
}

func DesconectarMongo(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)
	if err != nil {
		log.Println("⚠️ Error al desconectar MongoDB:", err)
	} else {
		fmt.Println("🧹 Desconectado de MongoDB.")
	}
}
