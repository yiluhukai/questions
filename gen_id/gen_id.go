package gen_id

import (
	"fmt"
	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

// setting的machineId应该是一个函数

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// Init functions is uesd init settings of sonyflake and created a instance of Sonyflake
func Init(machineID uint16) (err error) {
	sonyMachineID = machineID
	//configure Sonyflake
	st := sonyflake.Settings{}
	st.MachineID = getMachineID
	sonyFlake = sonyflake.NewSonyflake(st)
	return
}

// GetID can return a new id and err
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}
	return sonyFlake.NextID()
}
