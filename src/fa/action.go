package fa

import (
	"io"
)

type Action interface {
	GetName() string
	GetId() int
	String() string
	Fprint(w io.Writer)
}

type ActionList struct {
	actions []Action
}

func NewActionList() *ActionList {
	return &ActionList{}
}

func (this *ActionList) Empty() bool {
	return len(this.actions) == 0
}

func (this *ActionList) String() string {
	return ""
}

func (this *ActionList) Fprint(w io.Writer) {
}
