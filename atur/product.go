package atur

import (
	"net/http"

	"restgin/basecon"
	"restgin/entitas"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var err error
	p := entitas.Product{Name: c.PostForm("name"), Price: c.PostForm("price"), ImgUrl: c.PostForm("imageurl")}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}
	basecon.Db.Save(&p)
	c.JSON(http.StatusCreated, gin.H{"status": "ok", "result": p})
}

func FetchAllProduct(c *gin.Context) {
	var p []entitas.Product
	var pt []entitas.ProductTamp

	basecon.Db.Find(&p)

	if len(p) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
		return
	}

	for _, item := range p {

		pt = append(pt, entitas.ProductTamp{ID: item.ID, Name: item.Name, Price: item.Price, ImgUrl: item.ImgUrl, CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": pt})
}

func FetchSingleProduct(c *gin.Context) {
	var p entitas.Product
	pID := c.Param("id")

	basecon.Db.First(&p, pID)

	if p.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
		return
	}

	pt := entitas.ProductTamp{ID: p.ID, Name: p.Name, Price: p.Price, ImgUrl: p.ImgUrl, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": pt})
}

func UpdateProduct(c *gin.Context) {
	var p entitas.Product
	pID := c.Param("id")

	basecon.Db.First(&p, pID)

	if p.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
		return
	}

	basecon.Db.Model(&p).Update("name", c.PostForm("name"))
	basecon.Db.Model(&p).Update("price", c.PostForm("price"))
	basecon.Db.Model(&p).Update("imageurl", c.PostForm("imageurl"))
	pt := entitas.ProductTamp{ID: p.ID, Name: p.Name, Price: p.Price, ImgUrl: p.ImgUrl, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Product updated", "result": pt})
}

func DeleteProdcut(c *gin.Context) {
	var p entitas.Product
	pID := c.Param("id")

	basecon.Db.First(&p, pID)

	if p.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
		return
	}

	basecon.Db.Delete(&p)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Product deleted"})
}

func ProductMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
