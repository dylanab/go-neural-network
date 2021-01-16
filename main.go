package main;

import "fmt";
import . "NNFS/neuron";

func main() {

	testneuron := &Neuron{
		id: "testneuron",
		bias: 2
	};

	StartingInput(1, 0.2, testneuron);
	StartingInput(2, 0.8, testneuron);
	StartingInput(3, -0.5, testneuron);
	
	CalculateOutput(testneuron);

	fmt.Printf(testneuron.output);
}