package authController

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/GulabSinghSikarwar/preProjBackend/database"
	"github.com/GulabSinghSikarwar/preProjBackend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func createUserHelper(user *models.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Id = primitive.NewObjectID()
		user.User_id = user.Id.Hex()

		var token, refresh_token string
		user.Token = &token
		user.Refresh_Token = &refresh_token

	}

}
func Sign_up() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User

		err := c.BindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		validate := validator.New()
		validationErr := validate.Struct(user)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})

		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()

		if err != nil {
			log.Panic(err)

			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking database "})
			return
		}
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking database "})
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this email or phone exsist "})
		}

		createUserHelper(&user)

	}

}
