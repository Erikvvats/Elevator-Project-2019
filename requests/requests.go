package requests

import (
  ".././elevio"
  ".././elevator"
  "fmt"
)

//static???
func requests_above(e Elevator) int{
  for f := e.floor+1; f < _numFloors; f++ {
    for btn := 0; btn < _numButtons; btn++ {
      if e.requests[f][btn] {
        return 1
      }
    }
  }
  return 0
}

func requests_below(e Elevator) int{
  for f := 0; f < e.floor; f++ {
    for btn := 0; btn < _numButtons; btn++ {
      if e.requests[f][btn] {
        return 1
      }
    }
  }
  return 0
}


func request_chooseDir(e Elevator) MotorDirection {
  switch (e.motorDir) {
  case MD_Up:
    if res:= requests_above(e){
      return MD_Up
    } else if res:= requests_below(e){
      return MD_Down
    }
    return MD_Stop

  case MD_Down:
    

  case MD_Stop:
    if res:= requests_below(e){
      return MD_Down
    } else if res:= requests_above(e){
      return MD_Up
    }
    return MD_Stop
  }



}


func request_shouldStop(e Elevator) int {
  
}

func request_clearAtCurrentFloor(e Elevator) Elevator {
  
}