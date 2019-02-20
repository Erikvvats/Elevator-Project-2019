package orderFsm

type OrderState int

const (
	confirmed   OrderState = 1
	unconfirmed                = -1
	inactive                = 0
	idle 									= 2
)

type OrderUpdate struct {
	orderState OrderState
	IDlist[] int
}



for{
	select{

	}
}

/*
orderMatrix := make(chan [2][4]orderFsm.OrderUpdate{}) //GREIT AT DENNE ER GLOBAL?????

func merge(ourOrderState OrderUpdate, remoteOrderState OrderUpdate, chan confirmedOrders) OrderUpdate{
	confirmedOrder OrderUpdate
	confirmedOrder = ourOrderState
	switch ourOrderState.orderState {
	case inactive:
		switch remoteOrderState.orderState {
		case inactive:
		case unconfirmed:
			//append this ID to ourOrderState.IDlist
			// HVordAN FINNE DENNE pc iD
		case confirmed:
		case init:
		}
	case unconfirmed:
		switch remoteOrderState.orderState {
		case unconfirmed:
			// Append ID to list, if all ID's are present -> order confirmed -> send to channel confirmedOrders
		case confirmed:
			//Append IDs list
			//send to channel confirmed orders
		}
	case confirmed:
		switch remoteOrderState.orderState {
		case inactive:
			ourOrderState.orderState = inactive
			//delete all ID's from table
		case confirmed:
		}
	case init:
		ourOrderState.orderState = remoteOrderState.orderState
	}

	//return ourOrderState, ready for broadcasting/share
}

func updateOrderMatrix(){
	var id string
	flag.StringVar(&id, "id", "", "id of this peer")
	var driver_port string
	flag.StringVar(&driver_port, "driver_port", "", "port for elevator hw/sim")
	flag.Parse()

	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	}
	if driver_port == "" {
		driver_port = "15657"
	}
	numFloors := 4
	elevio.Init("localhost:"+driver_port, numFloors)
	drv_buttons := make(chan elevio.ButtonEvent)
	go elevio.PollButtons(drv_buttons)

		// We make a channel for receiving updates on the id's of the peers that are
	//  alive on the network
	peerUpdateCh := make(chan peers.PeerUpdate)
	// We can disable/enable the transmitter after it has been started.
	// This could be used to signal that we are somehow "unavailable".
	peerTxEnable := make(chan bool)
	go peers.Transmitter(10251, id, peerTxEnable)
	go peers.Receiver(10251, peerUpdateCh)

	// We make channels for sending and receiving our custom data types
	helloTx := make(chan Msg)
	helloRx := make(chan Msg)
	// ... and start the transmitter/receiver pair on some port
	// These functions can take any number of channels! It is also possible to
	//  start multiple transmitters/receivers on the same port.
	go bcast.Transmitter(10252, helloTx)
	go bcast.Receiver(10252, helloRx)

	// The example message. We just send one of these every second.

	orderMatrix := make(chan [2][4]orderFsm.OrderUpdate{})

	for{


		select {
			case p := <-peerUpdateCh:
				fmt.Printf("Peer update:\n")
				fmt.Printf("  Peers:    %q\n", p.Peers)
				fmt.Printf("  New:      %q\n", p.New)
				fmt.Printf("  Lost:     %q\n", p.Lost)

			case a := <-helloRx:
				fmt.Printf("Received: %#v\n", a)


			case button:= <- drv_buttons:
				if(button.ButtonType != 2){
					orderMatrix[button.ButtonType][button.Floor].orderState -> unconfirmed
				}
				orderMatrix[button.ButtonType][button.Floor].IDlist[0] -> id

				fmt.Println("Button push")
				helloTx <- OrderMsg{orderMatrix}

	}
}
}
*/
