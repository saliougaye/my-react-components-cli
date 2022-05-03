package helpers

import (
	"fmt"
	"os"
)

func CheckError(err error) {

	if err != nil {
		// TODO color with red
		fmt.Printf("❌\nerror %v\n", err)
		os.Exit(1)
	}

}
