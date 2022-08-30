package hendlers

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "github.com/NajmiddinAbdulhakim/iman/api-gateway/genproto"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// GetInfoFromNetANDInsertDB
// route /posts [get]
func (h *handler) GetInfoFromNetANDInsertDB(c *gin.Context) {
	var (
		empty       pb.Empty
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*time.Duration(2))
	defer cancel()

	info, err := h.serviceManager.GetService().GetInfoFromAPI(ctx, &empty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		log.Println(`failed to getting info from network: `, err.Error())
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*time.Duration(30))
	defer cancel()
	
	res, err := h.serviceManager.GetService().CreatePost(ctx, &pb.CreatePostsReq{
		Posts: info.Posts,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		log.Println(`failed to creating posts: `, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.Success)
	
	


}
