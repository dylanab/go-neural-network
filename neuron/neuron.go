package neuron;

struct Connection {
	value int;
	weight float;
}

struct Neuron {
	id string;
	inputs []*connection; 
	outputs []*connection;
	bias int;
	output int;
}

/*
func ConnectNeurons(input, output *Neuron) (*Connection){
	connection := &Connection{};
	
	if(output != nil && input != nil) {
		output.outputs = append(output.outputs, connection);
		input.inputs = append(input.inputs, connection);
	} else if(output != nil && input == nil) {
		output.outputs = append(output.outputs, connection);
	} else if(output == nil && input != nil) {
		input.inputs = append(input.inputs, connection);
	}

	return connection;
}
*/

func StartingInput(value int, weight float, n *Neuron) {
	connection := &Connection{value: value, weight: weight};
	n.inputs = append(n.inputconnections, connection);
}

func CalculateOutput(n *Neuron) {
	inputcount := len(n.inputs);
	output := 0.0;
	for i := 0; i < inputcount; i++ {
		output += (n.inputs[i].value * n.inputs[i].weight) + n.bias;
	}
	n.output = output;
}




