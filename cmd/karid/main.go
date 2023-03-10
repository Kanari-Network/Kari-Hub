package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	kariapp "github.com/Kanari-Netwrok/kari/app"
)

func main() {
	setAddressPrefixes(kariapp.AccountAddressPrefix)
	rootCmd := NewRootCmd(kariapp.MakeEncodingConfig())
	if err := svrcmd.Execute(rootCmd, "Kari", kariapp.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
