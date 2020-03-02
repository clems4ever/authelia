package commands

import (
	"fmt"
	"github.com/authelia/authelia/internal/authentication"

	"github.com/spf13/cobra"
)

var HashPasswordCmd = &cobra.Command{
	Use:   "hash-password [password]",
	Short: "Hash a password to be used in file-based users database",
	Run: func(cobraCmd *cobra.Command, args []string) {
		var err error
		var hash string
		sha512, _ := cobraCmd.Flags().GetBool("sha512")
		iterations, _ := cobraCmd.Flags().GetInt("iterations")
		salt, _ := cobraCmd.Flags().GetString("salt")
		saltLength, _ := cobraCmd.Flags().GetInt("salt-length")
		memory, _ := cobraCmd.Flags().GetInt("memory")
		parallelism, _ := cobraCmd.Flags().GetInt("parallelism")

		if sha512 {
			hash, err = authentication.HashPassword(args[0], salt, authentication.HashingAlgorithmSHA512, iterations, memory, parallelism, saltLength)
		} else {
			hash, err = authentication.HashPassword(args[0], salt, authentication.HashingAlgorithmArgon2id, iterations, memory, parallelism, saltLength)
		}

		if err != nil {
			fmt.Println(fmt.Sprintf("Error occured during hashing: %s", err))
		} else {
			fmt.Println(fmt.Sprintf("Password hash: %s", hash))
		}
	},
	Args: cobra.MinimumNArgs(1),
}
