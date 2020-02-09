package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/whoiskai/anime-tracker-bff/config"
	model "github.com/whoiskai/anime-tracker-bff/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	log "github.com/sirupsen/logrus"
)

// Static Collection
const AnimeCollection = "anime"

// MongoConfig get DB from config
func MongoConfig() *mgo.Database {
	db, err := config.GetMongoDB()
	if err != nil {
		log.Error(err)
	}
	return db
}

// TestHandle for testing
func TestHandle(c *gin.Context) {
	db := *MongoConfig()
	log.Info("MONGO RUNNING: ", db)

	c.JSON(200, gin.H{
		"test": "hello",
	})
}

// GetAllAnime returns all the anime
func GetAllAnime(c *gin.Context) {
	db := *MongoConfig()

	animes := model.Animes{}
	err := db.C(AnimeCollection).Find(bson.M{}).All(&animes)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get All Anime",
		})
		return
	}

	c.JSON(200, gin.H{
		"animes": &animes,
	})
}

// GetAnime get one specific anime base on ID
func GetAnime(c *gin.Context) {
	db := *MongoConfig()

	id := c.Param("id")
	idParse, errParse := strconv.Atoi(id)
	if errParse != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	anime := model.Anime{}
	err := db.C(AnimeCollection).Find(bson.M{"id": &idParse}).One(&anime)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get Anime",
		})
		return
	}

	c.JSON(200, gin.H{
		"anime": &anime,
	})
}

// CreateAnime creates a single anime
func CreateAnime(c *gin.Context) {
	db := *MongoConfig()

	anime := model.Anime{}
	err := c.Bind(&anime)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return
	}

	anime.CreatedAt = time.Now()
	anime.UpdatedAt = time.Now()

	err = db.C(AnimeCollection).Insert(anime)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Insert Anime",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success Insert Anime",
		"anime": &anime,
	})
}

// UpdateAnime updates a specific anime base on ID
func UpdateAnime(c *gin.Context) {
	db := *MongoConfig()

	id := c.Param("id")
	idParse, errParse := strconv.Atoi(id)
	if errParse != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	anime := model.Anime{}
	err := c.Bind(&anime)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return
	}

	anime.ID = idParse
	anime.UpdatedAt = time.Now()

	err = db.C(AnimeCollection).Update(bson.M{"id": &idParse}, anime)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Update Anime",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Succes Update Anime",
		"anime":    &anime,
	})
}

// DeleteAnime deletes a specific anime base on ID
func DeleteAnime(c *gin.Context) {
	db := *MongoConfig()

	id := c.Param("id")                   // Get Param
	idParse, errParse := strconv.Atoi(id) // Convert String to Int
	if errParse != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	err := db.C(AnimeCollection).Remove(bson.M{"id": &idParse})
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Delete Anime",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Succes Delete Anime",
	})
}