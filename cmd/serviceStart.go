package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	serviceCmd.AddCommand(serviceStartCmd)
}

// serviceStartCmd represents the serviceStart command
var serviceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an installed bitfan service",
	Run: func(cmd *cobra.Command, args []string) {
		s := getService(nil)
		err := s.Start()
		if err != nil {
			log.Printf("start service error : %s", err)
		} else {
			log.Println("start signal sent to service bitfan")
		}

	},
}
