package main

import (
  "./elevio"
  "fmt"
  //"./fsm"
)

func main(){
  numFloors := 4

  elevio.Init("localhost:15657", numFloors)
  var d elevio.MotorDirection = elevio.MD_Up

  drv_buttons := make(chan elevio.ButtonEvent)
  drv_floors  := make(chan int)
  drv_obstr   := make(chan bool)
  drv_stop    := make(chan bool)

  go elevio.PollButtons(drv_buttons)
  go elevio.PollFloorSensor(drv_floors)
  go elevio.PollObstructionSwitch(drv_obstr)
  go elevio.PollStopButton(drv_stop)


  //curr_floor := <- drv_floors
  /*if (curr_floor == -1){
    fsm.Fsm_onInitBetweenFloors()
  }*/
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

  }
}
