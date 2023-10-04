package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/HironixRotifer/mongodb-service-advertisements/database"
	"github.com/HironixRotifer/mongodb-service-advertisements/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AdvertisementCollection = database.AdvertisementData(database.Client, "Advertisement")
var Validate = validator.New()

func GetAdvertisementById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var searchAdvertisement []models.Advertisements
		var queryParam = c.Query("name")

		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid search index"})
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		searchQueryDB, err := AdvertisementCollection.Find(ctx, bson.M{"name": bson.M{"$regex": queryParam}})
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, "something went wrong while fetching data")
		}

		err = searchQueryDB.All(ctx, &searchAdvertisement)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusBadRequest, "invalid")
			return
		}

		defer searchQueryDB.Close(ctx)
		if err := searchQueryDB.Err(); err != nil {
			log.Panicln(err)
			c.IndentedJSON(http.StatusBadRequest, "invalid request")
			return
		}

		defer cancel()

		c.IndentedJSON(http.StatusOK, searchAdvertisement)
	}
}

func GetAdvertisements() gin.HandlerFunc {
	return func(c *gin.Context) {
		var advertisementList []models.Advertisements
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := AdvertisementCollection.Find(ctx, bson.D{{}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong while fetching data")
			return
		}

		err = cursor.All(ctx, &advertisementList)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		defer cursor.Close(ctx)
		if err := cursor.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusBadRequest, "invalid")
			return
		}

		defer cancel()

		c.IndentedJSON(http.StatusOK, advertisementList)
	}
}

func CreateAdvertisement() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var advertisement models.Advertisements

		if err := c.BindJSON(&advertisement); err != nil {
			log.Panicln(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		validationErr := Validate.Struct(advertisement)
		if validationErr != nil {
			log.Panicln(validationErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		}

		advertisement.ID = primitive.NewObjectID()

		_, inserterr := AdvertisementCollection.InsertOne(ctx, advertisement)
		if inserterr != nil {
			log.Println(inserterr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the advertisement didn`t created"})
			return
		}

		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Successfully created",
			"data":    advertisement.ID,
		})
	}
}
