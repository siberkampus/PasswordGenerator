package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("length", "l", 8, "Length of generated password")
	generateCmd.Flags().BoolP("digit", "d", false, "Include digits in generated password")
	generateCmd.Flags().BoolP("special", "s", false, "Include special characters in generated password")
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random passwords",
	Long: `Generate randon passwords with special options. For example:

PasswordGenerator -l 12 -d -s`,

	Run: GeneratePassword,
}

func GeneratePassword(cmd *cobra.Command, args []string) {
	rand.Seed(time.Now().UnixNano())
	digit, _ := cmd.Flags().GetBool("digit")
	special, _ := cmd.Flags().GetBool("special")
	length, _ := cmd.Flags().GetInt("length")

	char := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if digit {
		char += "0123456789"
	}
	if special {
		char += "-*+?_*/"
	}

	result := ""
	for i := 0; i < length; i++ {
		result += string(char[rand.Intn(len(char))])
	}
	fmt.Println("\x1b[31mGenerating Password....\x1b[0m")
	fmt.Println(result)
}
