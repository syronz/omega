package sample4

import (
	"omega/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"omega/internal/glog"
)

type Sample4Repository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func ProvideSample4Repostiory(c config.CFG) Sample4Repository {
	return Sample4Repository{DB: c.DB, Log: c.Log}
}

func (p *Sample4Repository) FindAll() []Sample4 {
	var sample4s []Sample4
	p.DB.Find(&sample4s)

	return sample4s
}

func (p *Sample4Repository) FindByID(id uint) Sample4 {
	var sample4 Sample4
	err := p.DB.First(&sample4, id).Error

	sample4.Extra = struct {
		LastVisit string
		Mark      int
	}{
		"2019",
		-15,
	}

	glog.Debug(sample4, id, err)

	return sample4
}

func (p *Sample4Repository) Save(sample4 Sample4) (s4 Sample4) {
	p.DB.Create(&sample4).Scan(&s4)

	p.Log.Debug(s4)
	glog.Debug(s4)
	// err = i.DB.Create(&i.Item).Scan(&item).Error

	return s4
}

func (p *Sample4Repository) Delete(sample4 Sample4) {
	p.DB.Delete(&sample4)
}
