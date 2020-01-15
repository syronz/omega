package sample5

type Sample5Service struct {
	Sample5Controller Sample5Controller
}

func ProvideSample5Service(p Sample5Controller) Sample5Service {
	return Sample5Service{Sample5Controller: p}
}

func (p *Sample5Service) FindAll() []Sample5 {
	return p.Sample5Controller.FindAll()
}

func (p *Sample5Service) FindByID(id uint) Sample5 {
	return p.Sample5Controller.FindByID(id)
}

func (p *Sample5Service) Save(sample5 Sample5) Sample5 {
	s4 := p.Sample5Controller.Save(sample5)

	return s4
}

func (p *Sample5Service) Delete(sample5 Sample5) {
	p.Sample5Controller.Delete(sample5)
}
