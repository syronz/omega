package sample4

type Sample4Service struct {
	Sample4Repository Sample4Repository
}

func ProvideSample4Service(p Sample4Repository) Sample4Service {
	return Sample4Service{Sample4Repository: p}
}

func (p *Sample4Service) FindAll() []Sample4 {
	return p.Sample4Repository.FindAll()
}

func (p *Sample4Service) FindByID(id uint) Sample4 {
	return p.Sample4Repository.FindByID(id)
}

func (p *Sample4Service) Save(sample4 Sample4) Sample4 {
	s4 := p.Sample4Repository.Save(sample4)

	return s4
}

func (p *Sample4Service) Delete(sample4 Sample4) {
	p.Sample4Repository.Delete(sample4)
}
