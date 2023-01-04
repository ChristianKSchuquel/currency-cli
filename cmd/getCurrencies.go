/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// getCurrenciesCmd represents the getCurrencies command
var getCurrenciesCmd = &cobra.Command{
	Use:   "getCurrencies",
	Short: "Get all currencies",
	Long:  `Get all currency names.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies.json")
		if err != nil {
			fmt.Print(err)
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err)
			return
		}

		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(getCurrenciesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCurrenciesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCurrenciesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
