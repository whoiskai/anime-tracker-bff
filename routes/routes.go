package routes

import (
	"github.com/gin-gonic/gin"
	handle_anime "github.com/whoiskai/anime-tracker-bff/handlers"
)

type Routes struct{}

func (c Routes) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", handle_anime.TestHandle)
		api.GET("/animes", handle_anime.GetAllAnime)
		api.POST("/animes", handle_anime.CreateAnime)
		api.GET("/animes/:id", handle_anime.GetAnime)
		api.PUT("/animes/:id", handle_anime.UpdateAnime)
		api.DELETE("/animes/:id", handle_anime.DeleteAnime)
	}
	r.Run(":8080")
}