package core

import (
	"math/rand"
	"sync"
)

type Brain struct {
	organsCount       int
	neuronsCount      int
	musclesCount      int
	neuronGroupsCount int

	Organs       []Organ
	NeuronGroups [][]Neuron
	Muscles      []Muscle

	visited       map[SignalReciever]bool
	unreachedOuts []*Muscle
}

func (b *Brain) LoadSignals(signals ...[]float64) {
	if b.organsCount != len(signals) {
		return
	}

	for i := 0; i < b.organsCount; i++ {
		b.Organs[i].tempStore = signals[i]
	}
}

func (b *Brain) ProcessSignals() {
	for i := 0; i < b.organsCount; i++ {
		b.Organs[i].ProcessSignals()
	}

	var wg sync.WaitGroup
	for i := 0; i < b.neuronGroupsCount; i++ {
		wg.Add(1)

		go func(neuronsGroup int) {
			for j := 0; j < len(b.NeuronGroups[neuronsGroup]); j++ {
				b.NeuronGroups[neuronsGroup][j].Process()
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < b.musclesCount; i++ {
		// fmt.Println(b.Muscles[i].MuscleMemory)
	}
}

func (b *Brain) IsAllInputsToAllOutputs() bool {
	b.visited = make(map[SignalReciever]bool)

	for i := 0; i < b.musclesCount; i++ {
		b.unreachedOuts = append(b.unreachedOuts, &b.Muscles[i])
	}

	for i := 0; i < b.organsCount; i++ {
		b.dfsTraversal(&b.Organs[i])
	}

	defer func() {
		b.unreachedOuts = make([]*Muscle, 0)
	}()

	return len(b.unreachedOuts) == 0
}

func (b *Brain) dfsTraversal(sr SignalReciever) {
	if b.visited[sr] {
		return
	}

	b.visited[sr] = true

	if sr.Type().EqualTo(NewBioTypeMuscle()) {
		for i, muscle := range b.unreachedOuts {
			if muscle == sr.(*Muscle) {
				b.unreachedOuts[i] = b.unreachedOuts[len(b.unreachedOuts)-1]
				b.unreachedOuts = b.unreachedOuts[:len(b.unreachedOuts)-1]
				break
			}
		}
	}

	for _, conn := range sr.GetAllConnections() {
		b.dfsTraversal(conn)
	}
}

func (b *Brain) GenerateNeuronConnections() {
	for i := 0; i < b.neuronGroupsCount; i++ {
		for j := 0; j < len(b.NeuronGroups[i]); j++ {
			b.NeuronGroups[i][j].ConnectTo(&b.NeuronGroups[i][rand.Intn(len(b.NeuronGroups[i]))])
		}
	}
}

func NewBrain(organShapes, muscleShapes []int, neuronsCount, neuronGroupsCount int) *Brain {
	if neuronGroupsCount > neuronsCount {
		panic("There are more neuron groups than neurons!")
	}

	brain := Brain{
		organsCount:       len(organShapes),
		musclesCount:      len(muscleShapes),
		neuronsCount:      neuronsCount,
		neuronGroupsCount: neuronGroupsCount,

		Organs:       make([]Organ, len(organShapes)),
		Muscles:      make([]Muscle, len(muscleShapes)),
		NeuronGroups: make([][]Neuron, neuronGroupsCount),

		visited:       map[SignalReciever]bool{},
		unreachedOuts: []*Muscle{},
	}

	// Create neurons in groups
	neuronsInGroup := neuronsCount / neuronGroupsCount
	neuronsInGroupRemainder := neuronsCount % neuronGroupsCount
	for i := 0; i < neuronGroupsCount; i++ {
		if i == neuronGroupsCount-1 {
			brain.NeuronGroups[i] = make([]Neuron, neuronsInGroup+neuronsInGroupRemainder)
			continue
		}

		brain.NeuronGroups[i] = make([]Neuron, neuronsInGroup)
	}

	for i := 0; i < neuronGroupsCount; i++ {
		if i == neuronGroupsCount-1 {
			for j := 0; j < neuronsInGroup+neuronsInGroupRemainder; j++ {
				brain.NeuronGroups[i][j] = NewNeuron()
			}
		}
		for j := 0; j < neuronsInGroup; j++ {
			brain.NeuronGroups[i][j] = NewNeuron()
		}
	}

	// Create organs
	for i := 0; i < brain.organsCount; i++ {
		brain.Organs[i] = NewOrgan(organShapes[i])
	}

	// Create muscles
	for i := 0; i < brain.musclesCount; i++ {
		brain.Muscles[i] = NewMuscle(muscleShapes[i])
	}

	// Connect organs and muscles to neurons
	for i := 0; i < neuronGroupsCount; i++ {
		for j := 0; j < brain.organsCount; j++ {
			for x := 0; x < organShapes[j]; x++ {
				brain.Organs[j].ConnectTo(&brain.NeuronGroups[i][rand.Intn(len(brain.NeuronGroups[i]))])
			}
		}
	}
	for i := 0; i < neuronGroupsCount; i++ {
		for j := 0; j < brain.musclesCount; j++ {
			for x := 0; x < muscleShapes[j]; x++ {
				brain.NeuronGroups[i][rand.Intn(len(brain.NeuronGroups[i]))].ConnectTo(&brain.Muscles[j])
			}
		}
	}

	for !brain.IsAllInputsToAllOutputs() {
		brain.GenerateNeuronConnections()
	}

	return &brain
}
