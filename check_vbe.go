package done

import "github.com/pkg/errors"

type Vbe struct {
	V bool
	E error
	*Vce[bool]
}

func VBE(val bool, err error) *Vbe {
	return &Vbe{
		V:   val,
		E:   err,
		Vce: VCE[bool](val, err),
	}
}

func (a *Vbe) TRUE() {
	a.assert1T1(a.Done())
}

func (a *Vbe) FALSE() {
	a.assert0F0(a.Done())
}

func (a *Vbe) Yes() {
	a.assert1T1(a.Done())
}

func (a *Vbe) No() {
	a.assert0F0(a.Done())
}

func (a *Vbe) OK() {
	a.assert1T1(a.Done())
}

func (a *Vbe) Not() {
	a.assert0F0(a.Done())
}

func (a *Vbe) assert1T1(v bool) {
	if !v {
		panic(errors.New("SHOULD BE TRUE BUT IS FALSE"))
	}
}

func (a *Vbe) assert0F0(v bool) {
	if v {
		panic(errors.New("SHOULD BE FALSE BUT IS TRUE"))
	}
}
