package core

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// type Brain struct {
// 	organsCount  int
// 	neuronsCount int
// 	musclesCount int

// 	Organs  []Organ
// 	Neurons []Neuron
// 	Muscles []Muscle

// 	visited       map[SignalReciever]bool
// 	unreachedOuts []*Muscle
// }

// // Connects random neuron to each organ's outputs
// func (b *Brain) connectOrgansToNeurons() {
// 	neuronsCount := len(b.Neurons)

// 	for i := 0; i < len(b.Organs); i++ {
// 		for j := 0; j < b.Organs[i].Shape; j++ {
// 			b.Organs[i].ConnectTo(&b.Neurons[rand.Intn(neuronsCount)])
// 		}
// 	}
// }

// func (b *Brain) connectMusclesToNeurons() {
// 	neuronsCount := len(b.Neurons)

// 	for i := 0; i < len(b.Muscles); i++ {
// 		for j := 0; j < b.Muscles[i].Shape; j++ {
// 			b.Neurons[rand.Intn(neuronsCount)].ConnectTo(&b.Muscles[i])
// 		}
// 	}
// }

// func (b *Brain) Tick() {
// 	for i := 0; i < b.organsCount; i++ {
// 		b.Organs[i].ProcessSignals()
// 	}

// 	for i := 0; i < b.neuronsCount; i++ {
// 		b.Neurons[i].Process()
// 	}

// 	// for i := 0; i < len(b.Muscles); i++ {
// 	// 	// fmt.Println("muscle buffer: ", b.Muscles[i].MuscleMemory) // TODO: handle muscles buffers
// 	// }
// }

// func (b *Brain) Run() {
// 	startTime := time.Now()

// 	for i := 0; i < 60; i++ {
// 		b.Tick()
// 	}

// 	endTime := time.Now()
// 	elapsedTime := endTime.Sub(startTime)
// 	fmt.Println(elapsedTime)

// 	var allReached bool = true
// 	for _, muscle := range b.Muscles {
// 		for _, memCell := range muscle.MuscleMemory {
// 			if memCell == 0 {
// 				allReached = false
// 			}
// 		}
// 	}

// 	fmt.Println("All Muscles reached:", allReached)
// }

// func (b *Brain) connectNeurons() {
// 	neuronsCount := len(b.Neurons)

// 	if neuronsCount < 2 {
// 		return
// 	}

// 	for i := 0; i < neuronsCount; i++ {
// 		b.Neurons[i].ConnectTo(
// 			&b.Neurons[rand.Intn(neuronsCount)],
// 		)
// 	}
// }

// func (b *Brain) ProcessSignals(signals [][]float64) {
// 	if len(signals) != b.organsCount {
// 		panic("Shapes does not match")
// 	}

// 	for i := 0; i < len(signals); i++ {
// 		b.Organs[i].ProcessSignals()
// 	}
// }

// func (b *Brain) LoadSignals(signals ...[]float64) {
// 	if len(signals) != b.organsCount {
// 		panic("Shapes do not match!")
// 	}

// 	for i := 0; i < b.organsCount; i++ {
// 		b.Organs[i].LoadSignals(signals[i])
// 	}

// }

// func (b *Brain) IsAllInputsToAllOutputs() bool {
// 	b.visited = make(map[SignalReciever]bool)

// 	for i := 0; i < b.musclesCount; i++ {
// 		b.unreachedOuts[i] = &b.Muscles[i]
// 	}

// 	for i := 0; i < b.organsCount; i++ {
// 		b.dfsTraversal(&b.Organs[i])
// 	}

// 	defer func() {
// 		b.unreachedOuts = make([]*Muscle, b.musclesCount)
// 	}()

// 	return len(b.unreachedOuts) == 0
// }

// func (b *Brain) dfsTraversal(sr SignalReciever) {
// 	if b.visited[sr] {
// 		return
// 	}

// 	b.visited[sr] = true

// 	// If the SignalReceiver is a Muscle, mark it as reached and remove it from the unreached list.
// 	if sr.Type().EqualTo(NewBioTypeMuscle()) {
// 		for i, muscle := range b.unreachedOuts {
// 			if muscle == sr.(*Muscle) {
// 				b.unreachedOuts[i] = b.unreachedOuts[len(b.unreachedOuts)-1]
// 				b.unreachedOuts = b.unreachedOuts[:len(b.unreachedOuts)-1]
// 				break
// 			}
// 		}
// 	}

// 	for _, conn := range sr.GetAllConnections() {
// 		b.dfsTraversal(conn)
// 	}
// }

// func NewBrain(organShapes, muscleShapes []int, neuronsCount int) *Brain {
// 	organs := make([]Organ, len(organShapes))
// 	for i := 0; i < len(organShapes); i++ {
// 		organs[i] = NewOrgan(organShapes[i])
// 	}

// 	neurons := make([]Neuron, neuronsCount)
// 	for i := 0; i < neuronsCount; i++ {
// 		neurons[i] = NewNeuron()
// 	}

// 	muscles := make([]Muscle, len(muscleShapes))
// 	for i := 0; i < len(muscleShapes); i++ {
// 		muscles[i] = NewMuscle(muscleShapes[i])
// 	}

// 	brain := Brain{
// 		musclesCount:  len(muscles),
// 		visited:       make(map[SignalReciever]bool),
// 		unreachedOuts: make([]*Muscle, len(muscles)),
// 		organsCount:   len(organs),
// 		Organs:        organs,
// 		neuronsCount:  neuronsCount,
// 		Neurons:       neurons,
// 		Muscles:       muscles,
// 	}

// 	brain.connectOrgansToNeurons()
// 	brain.connectMusclesToNeurons()

// 	for {
// 		if !brain.IsAllInputsToAllOutputs() {
// 			fmt.Println("not yet") // *debug TODO: delete
// 			brain.connectNeurons()
// 			continue
// 		}
// 		fmt.Println("here we are") // *debug TODO: delete
// 		break
// 	}

// 	brain.connectNeurons()

// 	return &brain
// }
