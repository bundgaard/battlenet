package gamedata

import (
	"fmt"
	"os"
	"testing"

	"github.com/bundgaard/battlenet/http"
)

func TestGameData(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()
	c := http.New("YOUROWNTOKENHERE", "YOUROWNTOKENHERE", http.Europe)
	gd := NewWith(c)
	if err := gd.Login(); err != nil {
		t.Error(err)
	}
	if _, err := gd.SeasonIndex(); err != nil {
		t.Error(err)
	}

	if _, err := gd.Season(24); err != nil {
		t.Error(err)
	}

	if _, err := gd.EraIndex(); err != nil {
		t.Error(err)
	}

	if _, err := gd.Era(14); err != nil {
		t.Error(err)
	}

	for i := 1; i < 15; i++ {

		if _, err := gd.EraLeaderBoard(i, "rift-team-2"); err != nil {
			t.Log(err)
		}
	}

}
