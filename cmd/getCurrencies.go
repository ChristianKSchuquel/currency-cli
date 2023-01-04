
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

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
}
