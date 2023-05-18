package cmd

import (
	"customer-data-migration/cmd/advisor"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "customer-data-migration",
	Short: "customer-data-migration",
	Long:  `customer-data-migration`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.AddCommand(advisor.AdvisorCmd)
}
