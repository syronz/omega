package sample5

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"omega/internal/glog"
)

type Sample5API struct {
	Sample5Service Sample5Service
}

func ProvideSample5API(p Sample5Service) Sample5API {
	return Sample5API{Sample5Service: p}
}

func (p *Sample5API) FindAll(c *gin.Context) {
	sample5s := p.Sample5Service.FindAll()

	c.JSON(http.StatusOK, gin.H{"sample5s": sample5s})
}

func (p *Sample5API) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sample5 := p.Sample5Service.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"sample5": sample5})
}

func (p *Sample5API) Create(c *gin.Context) {
	var sample5 Sample5
	err := c.BindJSON(&sample5)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdSample5 := p.Sample5Service.Save(sample5)
	glog.Debug(createdSample5)

	c.JSON(http.StatusOK, gin.H{"sample5": createdSample5})
}

func (p *Sample5API) Update(c *gin.Context) {
	var sample5 Sample5
	err := c.BindJSON(&sample5)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	sample5 = p.Sample5Service.FindByID(uint(id))
	if sample5 == (Sample5{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.Sample5Service.Save(sample5)

	c.Status(http.StatusOK)
}

func (p *Sample5API) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sample5 := p.Sample5Service.FindByID(uint(id))
	if sample5 == (Sample5{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.Sample5Service.Delete(sample5)

	c.Status(http.StatusOK)
}
