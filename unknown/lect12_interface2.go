package main 

import "fmt"

type USB interface {
	Name() string
	Connecter 	// nested 
}

type Connecter interface {
	Connect()
}


type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

// USB interface
func Disconnect0(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected0:", pc.name)	
		return	
	}
	fmt.Println("Unknown device 0.")
}

// empty interface
func Disconnect1(usb interface{}) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected1:", pc.name)	
		return	
	}
	fmt.Println("Unknown device 1.")
}

func Disconnect2(usb interface{}) {
	switch val := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected2:", val.name)
	default:
		fmt.Println("Unknown device 2.")
	}
}

type TVConnecter struct {
	name string
}

func (tv TVConnecter) Connect() {
	fmt.Println("Connect:", tv.name)
}


func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	fmt.Println(a.Name())
	a.Connect()
	Disconnect0(a)

	fmt.Println()
	a.Connect()
	Disconnect1(a)

	fmt.Println()
	a.Connect()
	Disconnect2(a)

	// interface convert from USB to Connecter
	fmt.Println()
	pc := PhoneConnecter{"PhoneConnecter"}	// USB
	var cc0 Connecter
	cc0 = Connecter(pc)
	cc0.Connect()

	fmt.Println()
	tv := TVConnecter{"TVConnecter"}
	var cc1 Connecter
	cc1 = Connecter(tv)
	cc1.Connect() 

}



