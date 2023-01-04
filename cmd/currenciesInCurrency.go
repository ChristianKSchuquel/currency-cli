
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var currenciesInCurrencyCmd = &cobra.Command{
	Use:   "currenciesInCurrency",
	Short: "Get the currency list in another currency value",
	Long:  `Get all currencies values in euro.`,
	Run: func(cmd *cobra.Command, args []string) {
		currency := args[0]
		url := fmt.Sprintf("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/%v.json", currency)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Print("error while requesting data, make sure the currency name is valid")
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
	rootCmd.AddCommand(currenciesInCurrencyCmd)
}
