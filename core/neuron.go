package core

type Neuron struct {
	biotype   BioType
	Dendrites map[int]*Synapse
	Axon      Axon

	Activity          float64
	MembranePotential float64
	Threshold         float64
	WellBeingRate     *float64 // or SatisfactionMeter instead of NeurotransmitterLevel, so the agent will have being rated and depends on the rate the agent gets FoodMeter changes
}

// Конечно, давайте рассмотрим, как все эти факторы могут влиять на работу нейрона в контексте обучения и адаптации:

// Обратная связь:

// Пример 1: Когда нейрон активируется и вызывает активацию соседних нейронов, синаптические связи могут усилиться, улучшая передачу сигнала между ними.
// Пример 2: Отрицательная обратная связь может снизить силу синапсов и подавить активацию нейрона при получении определенного типа сигнала.
// Нейромодуляция:

// Пример 1: При высоком уровне нейромедиатора серотонина, синаптическая пластичность может увеличиться, что способствует более эффективному обучению.
// Пример 2: Влияние ацетилхолина на активацию нейронов может увеличить или уменьшить восприимчивость синапсов в зависимости от текущего контекста.
// Формирование новых связей:

// Пример 1: При повторном активировании одних и тех же нейронов при обработке новой информации могут формироваться новые синапсы, усиливающие связь между ними.
// Пример 2: В ответ на обучающий опыт, нейроны могут вырастать новые дендриты, что позволяет им установить дополнительные синаптические связи.
// Долгосрочная пластичность:

// Пример 1: В процессе обучения, активация конкретных генов может изменить структуру нейрона, что влияет на его функцию на долгий срок.
// Пример 2: Длительное воздействие стресса может привести к изменениям в генетической экспрессии, влияя на работу нейрона и синапсов.
// Структурная пластичность:

// Пример 1: После обучения нейрон может создать новые аксоны, образуя связи с ранее несвязанными нейронами, что расширяет его сетевые возможности.
// Пример 2: Изменение геометрии нейрона может повысить его чувствительность к определенным типам входных сигналов.
// Сетевая динамика:

// Пример 1: Обучение может привести к синхронизации активации нейронов в определенных паттернах, что улучшает их совместную работу.
// Пример 2: Адаптация сетевой динамики может позволить нейронам эффективнее реагировать на разнообразные комбинации входных сигналов.

func (n *Neuron) Fire() {
	// fmt.Println("Fire")

	// signal := n.MembranePotential / float64(len(n.Axon.Terminal))

	const signal float64 = 10

	for i := 0; i < len(n.Axon.Terminal); i++ {
		n.Axon.Terminal[i].Synapse.RecieveSignal(
			signal+n.Activity,
			n.Axon.Terminal[i].port,
		)
	}

	n.MembranePotential = 0
}

func (n *Neuron) AddInputConnection(port int) {
	n.Dendrites[port] = NewSynapse()
}

func (n *Neuron) AddOutputConnection(sr SignalReciever, port int) {
	n.Axon.Terminal = append(n.Axon.Terminal, BioAddr{
		Synapse: nil,
		port:    port,
	})
}

func (n *Neuron) GetAllConnections() []SignalReciever {
	conns := make([]SignalReciever, len(n.Axon.Terminal))
	for i := 0; i < len(conns); i++ {
		conns[i] = n.Axon.Terminal[i].Synapse
	}

	return conns
}

func (n *Neuron) ConnectTo(synapse SignalReciever) {
	connPort := synapse.GetFreePort()

	n.Axon.Terminal = append(n.Axon.Terminal, BioAddr{
		Synapse: synapse,
		port:    connPort,
	})

	synapse.AddInputConnection(connPort)
}

func (n *Neuron) Process() {
	// TEST DUNGEROUS TODO DELETE OR CHANGE OR WHATEVER IT IS NOT COMPLETE AAAAAAAA
	for index, _ := range n.Dendrites {
		n.Dendrites[index].Update(-1)
	}

	if n.MembranePotential+*n.WellBeingRate > n.Threshold-n.Activity {
		n.Activity += 2
		n.Fire()
	} else {
		n.Activity -= 1
	}

	if n.Activity > 6 {
		n.Activity = 6
	}

	if n.Activity < 6 {
		n.Activity = -6
	}
}

func (n *Neuron) RecieveSignal(signal float64, dendritePort int) {
	// fmt.Println("Neuron recieved signal", signal)
	signal = signal + n.Dendrites[dendritePort].Strength

	n.Dendrites[dendritePort].Update(2)

	n.MembranePotential += signal
}

func (n *Neuron) Type() BioType {
	return n.biotype
}

func NewNeuron(wellBeingRate *float64) Neuron {
	return Neuron{
		biotype:           NewBioTypeNeuron(),
		Dendrites:         make(map[int]*Synapse),
		MembranePotential: 0,
		Threshold:         16,
		Axon:              NewAxon(),
		WellBeingRate:     wellBeingRate,
	}
}

func (n *Neuron) GetFreePort() int {
	for i := 0; ; i++ {
		if _, exist := n.Dendrites[i]; !exist {
			return i
		}
	}
}

type Synapse struct {
	Strength float64
}

func (s *Synapse) Update(bias float64) {

	s.Strength += bias

	if s.Strength > 5 {
		s.Strength = 5
	}

	if s.Strength < 5 {
		s.Strength = -5
	}
}

func NewSynapse() *Synapse {
	return &Synapse{
		Strength: 0,
	}
}

type Axon struct {
	Bias     float64
	Terminal []BioAddr
}

func NewAxon() Axon {
	return Axon{
		Bias:     3,
		Terminal: []BioAddr{},
	}
}

type SignalReciever interface {
	AddInputConnection(int)
	GetFreePort() int
	RecieveSignal(signal float64, dendritePort int)
	Type() BioType
	GetAllConnections() []SignalReciever
}

type BioType [2]bool // can store up to 4 types, we need only 3 - organ, neuron, muscle

func NewBioTypeOrgan() BioType {
	return BioType{false, false}
}

func NewBioTypeNeuron() BioType {
	return BioType{false, true}
}

func NewBioTypeMuscle() BioType {
	return BioType{true, false}
}

func (bt BioType) EqualTo(other BioType) bool {
	if bt[0] == other[0] && bt[1] == other[1] {
		return true
	}

	return false
}
