/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/xen0bit/ipom/pkg/ris"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Downloading RISWhois IPv4 Dataset...")
		rwv4, err := ris.RISWhoisV4()
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile("riswhoisv4.txt", []byte(rwv4), 0755); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Downloading RISWhois IPv6 Dataset...")
		rwv6, err := ris.RISWhoisV6()
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile("riswhoisv6.txt", []byte(rwv6), 0755); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// refreshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// refreshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
