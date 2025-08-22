/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/xen0bit/ipom/pkg/ris"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ipom",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		asn, _ := cmd.Flags().GetString("asn")
		if asn == "" {
			log.Fatal("")
		}
		expand, _ := cmd.Flags().GetBool("expand")
		separator, _ := cmd.Flags().GetString("separator")
		filePath := "riswhoisv4.txt"

		// Check if the file exists
		_, err := os.Stat(filePath)

		if errors.Is(err, os.ErrNotExist) {
			log.Fatalf("File '%s' does not exist.\n", filePath)
		} else if err != nil {
			// Handle other potential errors (e.g., permissions)
			log.Fatalf("Error checking file '%s': %v\n", filePath, err)
		} else {
			rwrs, err := ris.LoadV4()
			if err != nil {
				log.Fatal(err)
			}
			var outstring string
			for _, prefix := range rwrs.GetRecords(asn) {
				if expand {
					addr := prefix.Addr()
					// Calculate the number of addresses in the prefix
					// (2^(32-bits) for IPv4, 2^(128-bits) for IPv6)
					var numAddresses uint64
					if addr.Is4() {
						numAddresses = 1 << (32 - prefix.Bits())
					} else if addr.Is6() {
						numAddresses = 1 << (128 - prefix.Bits())
					} else {
						log.Fatal("Unsupported address family")
					}

					for i := uint64(0); i < numAddresses; i++ {
						outstring += addr.String() + separator
						addr = addr.Next()
					}
				} else {
					outstring += prefix.String() + separator
				}
			}
			//Trim trailing separator
			outstring = outstring[:len(outstring)-1]
			fmt.Println(outstring)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ipom.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("asn", "a", "", "Target ASN Number")
	rootCmd.Flags().BoolP("expand", "e", false, "Expand the CIDR prefix into all individual IPs in range")
	rootCmd.Flags().StringP("separator", "s", "\n", "Expand the CIDR prefix into all individual IPs in range")
}
