package sample4

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"omega/internal/glog"
)

type Sample4API struct {
	Sample4Service Sample4Service
}

func ProvideSample4API(p Sample4Service) Sample4API {
	return Sample4API{Sample4Service: p}
}

func (p *Sample4API) FindAll(c *gin.Context) {
	sample4s := p.Sample4Service.FindAll()

	c.JSON(http.StatusOK, gin.H{"sample4s": sample4s})
}

func (p *Sample4API) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sample4 := p.Sample4Service.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"sample4": sample4})
}

func (p *Sample4API) Create(c *gin.Context) {
	var sample4 Sample4
	err := c.BindJSON(&sample4)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdSample4 := p.Sample4Service.Save(sample4)
	glog.Debug(createdSample4)

	c.JSON(http.StatusOK, gin.H{"sample4": createdSample4})
}

func (p *Sample4API) Update(c *gin.Context) {
	var sample4 Sample4
	err := c.BindJSON(&sample4)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	sample4 = p.Sample4Service.FindByID(uint(id))
	if sample4 == (Sample4{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.Sample4Service.Save(sample4)

	c.Status(http.StatusOK)
}

func (p *Sample4API) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sample4 := p.Sample4Service.FindByID(uint(id))
	if sample4 == (Sample4{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.Sample4Service.Delete(sample4)

	c.Status(http.StatusOK)
}
