package fsm
import (
  ".././elevio"
  "fmt"
  )

//NoT working!!
/*
func Fsm_onInitBetweenFloors() {
  var d elevio.MotorDirection = elevio.MD_Down
  elevio.SetMotorDirection(d)
  //Need to set elevator state -> dirn = D_down, behaviour = Moving
}*/


func Fsm_onRequestButtonPress(button elevio.ButtonEvent ){
  fmt.Println("Floor: ", button.Floor," Dir: ", button.Button)
}
