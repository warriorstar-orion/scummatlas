package script

import "fmt"

type Var struct {
	val int
}

func (v Var) String() string {
	if v.val&0x4000 > 0 {
		return fmt.Sprintf("Local[%d]", v.val-0x4000)
	} else if v.val&0x8000 > 0 {
		return fmt.Sprintf("Bit[%d]", v.val-0x8000)
	} else {
		name := varNames[byte(v.val)]
		if name == "" {
			name = fmt.Sprintf("Var[%d]", v.val)
		}
		return name
	}
}
