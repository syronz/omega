package sample5

import (
	"omega/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"omega/internal/glog"
)

type Sample5Repository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func ProvideSample5Repostiory(c config.CFG) Sample5Repository {
	return Sample5Repository{DB: c.DB, Log: c.Log}
}

func (p *Sample5Repository) FindAll() []Sample5 {
	var sample5s []Sample5
	p.DB.Find(&sample5s)

	return sample5s
}

func (p *Sample5Repository) FindByID(id uint) Sample5 {
	var sample5 Sample5
	err := p.DB.First(&sample5, id).Error

	sample5.Extra = struct {
		LastVisit string
		Mark      int
	}{
		"2019",
		-15,
	}

	glog.Debug(sample5, id, err)

	return sample5
}

func (p *Sample5Repository) Save(sample5 Sample5) (s4 Sample5) {
	p.DB.Create(&sample5).Scan(&s4)

	p.Log.Debug(s4)
	glog.Debug(s4)
	// err = i.DB.Create(&i.Item).Scan(&item).Error

	return s4
}

func (p *Sample5Repository) Delete(sample5 Sample5) {
	p.DB.Delete(&sample5)
}
