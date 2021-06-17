package main

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	xflagsCmd.Flags().StringP("arch", "a", "", "Sets the BuildArch flag value")
	xflagsCmd.Flags().StringP("build", "b", "0", "Sets the BuildNumber flag value")
	xflagsCmd.Flags().StringP("extra", "e", "", "Sets the BuildStateExtra flag value")
}

var xflagsCmd = &cobra.Command{
	Use:   "xflags",
	Run:   runXFlags,
	Short: "Generate X LDFlags for building Authelia",
}

func runXFlags(cobraCmd *cobra.Command, _ []string) {
	arch, err := cobraCmd.Flags().GetString("arch")
	if err != nil {
		log.Fatal(err)
	}

	build, err := cobraCmd.Flags().GetString("build")
	if err != nil {
		log.Fatal(err)
	}

	extra, err := cobraCmd.Flags().GetString("extra")
	if err != nil {
		log.Fatal(err)
	}

	flags, err := getXFlags(arch, "", build, extra)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Join(flags, " "))
}