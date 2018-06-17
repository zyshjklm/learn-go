package main

import "fmt"

// Widget struct
type Widget struct {
	X, Y int
}

// Label struct
type Label struct {
	Widget        // Embedding(delegation)
	Text   string // Aggregation
	X      int    // Override
}

// Paint func
func (label Label) Paint() {
	fmt.Printf("[%p] = Label.Paint(%q)\n", &label, label.Text)
}

func main() {
	lab := Label{Widget{10, 10}, "State", 100}
	fmt.Printf("X=%d, Y=%d, Text=%s, Widget.X=%d, Widget.Y=%d\n",
		lab.X, lab.Y, lab.Text, lab.Widget.X, lab.Widget.Y)
	// X=100, Y=10, Text=State, Widget.X=10, Widget.Y=10

	fmt.Println()
	fmt.Printf("%+v\n%v\n", lab, lab)
	// {Widget:{X:10 Y:10} Text:State X:100}
	// {{10 10} State 100}
}
