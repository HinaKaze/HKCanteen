package common

type PrivilegeSignal int64

const (
	PrivilegeOrderCreate PrivilegeSignal = 0x01
	PrivilegeChargeMoney PrivilegeSignal = 0x02
)
