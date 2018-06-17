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

// Button struct based on Label
type Button struct {
	Label // Embedding(delegation)
}

// NewButton to new a Button
func NewButton(x, y int, text string) Button {
	return Button{Label{Widget{x, y}, text, x}}
}

// Paint to paint Button
func (button Button) Paint() {
	// override
	fmt.Printf("[%p] = Button.Paint(%q)\n", &button, button.Text)
}

// Click func
func (button Button) Click() {
	fmt.Printf("[%p] = Button.Click()\n", &button)
}

// ListBox for list of Box
type ListBox struct {
	Widget
	Texts []string
	Index int
}

// Paint to paint ListBox
func (lb ListBox) Paint() {
	fmt.Printf("[%p] = ListBox.Paint(%q)\n", &lb, lb.Texts)
}

// Click func
func (lb ListBox) Click() {
	fmt.Printf("[%p] = ListBox.Click()\n", &lb)
}

// Painter interface
type Painter interface {
	Paint()
}

// Clicker interface
type Clicker interface {
	Click()
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

	btn1 := Button{Label{Widget{10, 70}, "OK", 10}}
	btn2 := NewButton(50, 70, "Canel")
	lbox := ListBox{Widget{10, 40}, []string{"AL", "AK", "AZ", "AR"}, 0}

	fmt.Println("\n--- for struct ---")
	for _, painter := range []Painter{lab, lbox, btn1, btn2} {
		painter.Paint()
	}
	// [0xc42007a2a0] = Label.Paint("State")
	// [0xc42007a2d0] = ListBox.Paint(["AL" "AK" "AZ" "AR"])
	// [0xc42007a300] = Button.Paint("OK")
	// [0xc42007a330] = Button.Paint("Canel")

	fmt.Println("\n--- for interface ---")
	for _, wid := range []interface{}{lab, lbox, btn1, btn2} {
		if clicker, ok := wid.(Clicker); ok {
			clicker.Click()
		}
	}
	// [0xc42007a420] = ListBox.Click()
	// [0xc42007a450] = Button.Click()
	// [0xc42007a480] = Button.Click()

}
