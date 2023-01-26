package repository

import (
	"basic_echo/models"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitializeMongoConnection(ctx *echo.Echo) error {
	var connectionString string
	var err error
	err = godotenv.Load(".env")

	if err == nil {
		connectionString = "mongodb+srv://donnukrit:" + os.Getenv("MONGO_PASSWORD") + "@computer-store-database.0rlefqu.mongodb.net/"

	} else {
		err = errors.New("failed to load env")
		return err

	}

	fmt.Println(connectionString)
	mongoClient, err = mongo.NewClient(options.Client().ApplyURI(connectionString))

	mongoClient.Connect(context.TODO())
	if err == nil {
		fmt.Println("Mongo Connect : Success")
	} else {
		fmt.Println("Mongo Connect : Failed")
		fmt.Println(err.Error())
	}
	return err
}
func RepCreateNewProduct(newProduct models.Product) interface{} {

	products := RepGetAllProducts()

	var err interface{}

	collection := mongoClient.Database("store").Collection("products")

	for _, p := range products {
		if strings.EqualFold(p.Name, newProduct.Name) {

			err = models.ErrorData{
				Code:         4001,
				ErrorTitle:   "Duplicate",
				ErrorMessage: "This item already has in store",
			}
			break

		}

	}
	if err == nil {

		currentTime := time.Now().Unix()
		uuid := uuid.New()
		newProduct.Uuid = uuid.String()
		newProduct.Created_at = currentTime
		newProduct.Updated_at = currentTime
		collection.InsertOne(context.TODO(), newProduct)

	}

	return err
}

func RepGetAllProducts() []models.Product {

	filter := bson.D{}
	productsCollection := mongoClient.Database("store").Collection("products")

	cursor, _ := productsCollection.Find(context.TODO(), filter)

	var result []models.Product
	cursor.All(context.TODO(), &result)
	return result
}

func RepGetAllProductByPriceRange(min float32, max float32) []models.Product {
	var result []models.Product
	productsCollection := mongoClient.Database("store").Collection("products")
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "price", Value: bson.D{{Key: "$gte", Value: min}}}},
				bson.D{{Key: "price", Value: bson.D{{Key: "$lte", Value: max}}}},
			},
		},
	}
	cursor, _ := productsCollection.Find(context.TODO(), filter)
	cursor.All(context.TODO(), &result)

	return result
}

func RepGetProductByCategory() {}

func RepDeleteProduct() {}

func RepGetSortedProductsByPrice() {

}
