package main

import (
  "./elevio"
  "./network/bcast"
  "./network/localip"
  "./network/orderFsm"
  "./network/peers"
  "fmt"
  "flag"
  "os"
  //"time"
  //"./fsm"
)

/*
type Msg struct {
  ID string
  Button elevio.ButtonEvent

}
*/

func main(){
  var id string
  flag.StringVar(&id, "id", "", "id of this peer")
  var driver_port string
  flag.StringVar(&driver_port, "driver_port", "", "port for elevator hw/sim")
  flag.Parse()

  // ... or alternatively, we can use the local IP address.
  // (But since we can run multiple programs on the same PC, we also append the
  //  process ID)
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
  //var d elevio.MotorDirection = elevio.MD_Up

  drv_buttons := make(chan elevio.ButtonEvent)
  //drv_floors  := make(chan int)
  //drv_obstr   := make(chan bool)
  //drv_stop    := make(chan bool)

  go elevio.PollButtons(drv_buttons)
  //go elevio.PollFloorSensor(drv_floors)
  //go elevio.PollObstructionSwitch(drv_obstr)
  //go elevio.PollStopButton(drv_stop)


  //curr_floor := <- drv_floors
  /*if (curr_floor == -1){
    fsm.Fsm_onInitBetweenFloors()
  }*/
  /*
  d = elevio.MD_Up
  elevio.SetMotorDirection(d)

  for{
    select{
    case button := <- drv_buttons:
      fmt.Printf("%+v\n", button)
      elevio.SetButtonLamp(button.Button, button.Floor, true)
      //elevio.SetMotorDirection(d)
    case floor := <- drv_floors:
      fmt.Println(floor);
      if(floor == 0){
        d = elevio.MD_Up
      }
      if(floor == 3){
        d = elevio.MD_Down
      }
      elevio.SetFloorIndicator(floor)
      elevio.SetMotorDirection(d)

    case stop := <- drv_stop:
          fmt.Printf("%+v\n", stop)
          for f := 0; f < numFloors; f++ {
              for b := elevio.ButtonType(0); b < 3; b++ {
                  elevio.SetButtonLamp(b, f, false)
                  d = elevio.MD_Stop
              }
          }
    }

  }*/


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
  orderMatrixTx := make(chan [2][4]orderFsm.OrderUpdate)


  //  start multiple transmitters/receivers on the same port.
  go bcast.Transmitter(10252, orderMatrix)
  go bcast.Receiver(10252, helloRx)

  // The example message. We just send one of these every second.

  //orderMatrix := make(chan [2][4]orderFsm.OrderUpdate())

  orderMatrixTx <- dasuidoi

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
