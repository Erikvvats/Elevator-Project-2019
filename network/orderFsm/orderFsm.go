package orderFsm

type OrderState int

const (
	confirmed   OrderState = 1
	unconfirmed                = -1
	inactive                = 0
)

type OrderUpdate struct {
	orderState OrderState
	IDlist[] int
}
