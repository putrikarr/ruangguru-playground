package laptop

type Sleeping struct {
	Laptop *Laptop
}

func (s Sleeping) Press() {
	// TODO: answer here
	s.Laptop.ChangeCurrentState("On")
	s.Laptop.ChangeState(On{s.Laptop})
}

func (s Sleeping) CanTurnOnLaptop() bool {
	return true
}

func (s Sleeping) Sleep() {
	// TODO: answer here
	return
}
