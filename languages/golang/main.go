package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	startTime := time.Now().UnixMilli()
	fmt.Printf("🔹 Proceso iniciado en %d ms\n", startTime)

	// Cargar .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se pudo cargar el archivo .env")
	}

	// Leer variables de entorno
	mongoHost := getEnv("MONGO_HOST", "localhost")
	mongoPort := getEnv("MONGO_PORT", "27017")
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	var mongoCredential string
	if mongoUser != "" && mongoPassword != "" {
		mongoCredential = fmt.Sprintf("%s:%s@", mongoUser, mongoPassword)
	}
	uri := fmt.Sprintf("mongodb://%s%s:%s", mongoCredential, mongoHost, mongoPort)

	// Conexión a MongoDB
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	fmt.Println("✅ Conectado a MongoDB")

	db := client.Database("sample_db")
	collection := db.Collection("mycollection")

	// Crear documento
	doc := bson.D{
		{Key: "name", Value: "Martín Rivas"},
		{Key: "age", Value: 42},
		{Key: "source", Value: "Go"},
	}
	insertResult, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
	id := insertResult.InsertedID.(primitive.ObjectID)
	fmt.Printf("📝 Documento insertado con _id: %s\n", id.Hex())

	// Buscar por _id
	var result bson.M
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("🔍 Documento encontrado por ID:", result)

	// Buscar todos los documentos
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var allDocs []bson.M
	if err = cursor.All(context.TODO(), &allDocs); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("📄 Documentos encontrados (%d): %v\n", len(allDocs), allDocs)

	endTime := time.Now().UnixMilli()
	fmt.Printf("🔹 Proceso finalizado en %d ms\n", endTime)
	fmt.Printf("⏱️ Duración total: %d ms\n", endTime-startTime)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}