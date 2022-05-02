package helpers

import (
	"time"

	"github.com/briandowns/spinner"
)

// TODO customize colors

func Loading(prefix, finalMSG string) *spinner.Spinner {

	s := spinner.New(spinner.CharSets[40], 100*time.Millisecond)
	s.Prefix = prefix
	s.FinalMSG = finalMSG
	return s
}
