package components

// Number of individuals of each population
const iNum = 100

// Population struct
type Population struct {
	individuals []*Individual
	generation  int64
}

// CreateNewPopulation function
func CreateNewPopulation(gn int64) *Population {
	pop := Population{
		generation: gn,
	}

	return &pop
}

// Generation return all individuals
func (p *Population) Generation() int64 {
	return p.generation
}

// GetIndividuals return all individuals
func (p *Population) GetIndividuals() []*Individual {
	return p.individuals
}

// AddIndividual add an individual to the population
func (p *Population) AddIndividual(i *Individual) {
	p.individuals = append(p.individuals, i)
}

// RemoveIndividual remove an individual from the population
func (p *Population) RemoveIndividual(i int) {
	p.individuals[i] = p.individuals[len(p.individuals)-1]
	p.individuals = p.individuals[:len(p.individuals)-1]
}
