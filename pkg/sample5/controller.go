package sample5

type Sample5Controller struct {
	Sample5Repository Sample5Repository
}

func ProvideSample5Controller(p Sample5Repository) Sample5Controller {
	return Sample5Controller{Sample5Repository: p}
}

func (p *Sample5Controller) FindAll() []Sample5 {
	return p.Sample5Repository.FindAll()
}

func (p *Sample5Controller) FindByID(id uint) Sample5 {
	return p.Sample5Repository.FindByID(id)
}

func (p *Sample5Controller) Save(sample5 Sample5) Sample5 {
	s4 := p.Sample5Repository.Save(sample5)

	return s4
}

func (p *Sample5Controller) Delete(sample5 Sample5) {
	p.Sample5Repository.Delete(sample5)
}
