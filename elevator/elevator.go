package elevator

import (
  ".././elevio"
  "fmt"
)

type ElevatorBehaviour int


const (
  EB_Idle ElevatorBehaviour = 0
  EB_DoorOpen ElevatorBehaviour = 1
  EB_Moving ElevatorBehaviour = 2
)

type ClearRequestVariant int
const (
  CV_All
  CV_InDirn
)

type Elevator struct{
  floor int
  motorDir MotorDirection
  requests[_numFloors][_numButtons] int
  behaviour ElevatorBehaviour
  config: struct {
    clearRequestVariant ClearRequestVariant
    doorOpenDuration_s float64
  }
}

//func elevator_uninitialized() Elevator{

//
}
