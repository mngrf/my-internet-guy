package core

type BioPort struct {
	Storage     map[BioAddr][2]float64 //potential, weights
	Connections []BioConn
}

type BioDepot struct {
	Storage map[BioAddr]float64
}

type BioSource struct {
	Connections []BioConn
}

type BioConn struct {
	storage *map[BioAddr]float64
	addr    BioAddr
}

type BioAddr int16
