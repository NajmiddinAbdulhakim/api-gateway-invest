package api

import (
	han "github.com/NajmiddinAbdulhakim/iman/api-gateway/api/handlers"
	"github.com/NajmiddinAbdulhakim/iman/api-gateway/config"
	"github.com/NajmiddinAbdulhakim/iman/api-gateway/service"
	"github.com/gin-gonic/gin"
)

type Option struct {
	Conf           config.Config
	ServiceManager service.IServiceManager
}

func New(o *Option) *gin.Engine {
	router := gin.New()

	handler := han.New(&han.HandlerConfig{
		ServiceManager: o.ServiceManager,
		Cfg:            o.Conf,
	})

	crud := router.Group("/")
	{
		crud.POST(`posts`, handler.CreatePost)
		crud.GET(`posts/:id`, handler.GetPost)
		crud.PUT(`posts/:id`, handler.UpdatePost)
		crud.DELETE(`posts/:id`, handler.DeletePost)
		crud.GET(`posts`, handler.ListPosts)
	}

	web := router.Group("/")
	{
		web.GET(`get-insert`, handler.GetInfoFromNetANDInsertDB)
	}

	router.Run()

	return router
}
