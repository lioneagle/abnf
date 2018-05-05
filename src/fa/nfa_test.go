package fa

import (
	_ "fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestNfaTransitionFprint(t *testing.T) {
	transition := NewNfaTransition()
	transition.AddChar('a')
	transition.AddChar(' ')

	destState := NewNfaState()
	destState.id = 1

	transition.SetDestSate(destState)

	str := transition.String()
	test.EXPECT_EQ(t, str, "{ ' ', 'a' } --> q0001", "")

	transition.SetEpsilon()
	str = transition.String()
	test.EXPECT_EQ(t, str, "{ epsilon } --> q0001", "")
}
