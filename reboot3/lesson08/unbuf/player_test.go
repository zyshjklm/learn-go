package unbuf

import (
	"testing"
)

func TestPlayer(t *testing.T) {
	obj := NewPlayer("chen", "song")
	obj.StartPlayer()
}
