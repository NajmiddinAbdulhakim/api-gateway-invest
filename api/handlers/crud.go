package hendlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "github.com/NajmiddinAbdulhakim/iman/api-gateway/genproto"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreatePost creates post
// route /posts [post]
func (h *handler) CreatePost(c *gin.Context) {
	var (
		body        pb.Post
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(`failed to bind json: `, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	res, err := h.serviceManager.CRUDService().CreatePost(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(`failed to create post: `, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetPost gets post by id
// route /posts/{id} [get]
func (h *handler) GetPost(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	getId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": `argument not integer`,
		})
		log.Println(`failed to argument not integer: `, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	res, err := h.serviceManager.CRUDService().GetPostById(
		ctx, &pb.PostByIdReq{
			Id: getId,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(`failed to getting post by id: `, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdatePost updates post
// route /posts/{id} [put]
func (h *handler) UpdatePost(c *gin.Context) {
	var (
		body        pb.Post
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(`failed to bind json: `, err.Error())
		return
	}
	body.Id, err = strconv.ParseInt(c.Param(`id`), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": `argument not integer for update`,
		})
		log.Println(`failed to argument not integer for update: `, err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	res, err := h.serviceManager.CRUDService().UpdatePost(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(`failed to update post: `, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// DeletePost deletes post by id
// route /posts/{id} [delete]
func (h *handler) DeletePost(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	getId, err := strconv.ParseInt(c.Param(`id`), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": `argument not integer for delete`,
		})
		log.Println(`failed to argument not integer for delete `, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	res, err := h.serviceManager.CRUDService().DeletePost(
		ctx, &pb.PostByIdReq{
			Id: getId,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(`failed to delete post by id: `, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// ListPosts returns list of posts
// route /posts [get]
func (h *handler) ListPosts(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	page, err := strconv.ParseInt(c.Query(`page`), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": `argument not integer for page`,
		})
		log.Println(`failed to argument not integer for page: `, err.Error())
		return
	}

	limit, err := strconv.ParseInt(c.Query(`limit`), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": `argument not integer for limit`,
		})
		log.Println(`failed to argument not integer fot limit: `, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	res, err := h.serviceManager.CRUDService().ListPosts(
		ctx, &pb.PostListReq{
			Limit: limit,
			Page: page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(`failed to getting list of posts: `, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}