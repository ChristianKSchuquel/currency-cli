package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var compareCurrenciesCmd = &cobra.Command{
	Use:   "compareCurrencies",
	Short: "Compare 2 currencies",
	Long: `Gets 2 currencies and gets the value of the first in the second. 
	For example: getCurrencies brl usd
`,
	Run: func(cmd *cobra.Command, args []string) {
		currency1 := args[0]
		currency2 := args[1]
		url := fmt.Sprintf("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/%v/%v.json", currency1, currency2)

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
	rootCmd.AddCommand(compareCurrenciesCmd)
}
