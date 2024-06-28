package script

import "fmt"

type Var struct {
	val  int
	more *Var
}

func NewVar(val int, more ...Var) Var {
	if len(more) == 1 {
		return Var{val, &more[0]}
	}
	return Var{val, nil}
}

func (v Var) String() string {
	more := ""
	if v.more != nil {
		more = fmt.Sprintf(" + %v", v.more)
	}
	if v.val&0x4000 > 0 {
		return fmt.Sprintf("Local[%d%v]", v.val-0x4000, more)
	} else if v.val&0x8000 > 0 {
		return fmt.Sprintf("Bit[%d%v]", v.val-0x8000, more)
	} else {
		name := varNames[byte(v.val)]
		if name == "" {
			name = fmt.Sprintf("Var[%d%v]", v.val, more)
		}
		return name
	}
}
