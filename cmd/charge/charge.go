package charge

import (
	"fmt"

	"github.com/LeuenbergerP/ln-atm-cli/cmd/charge"
	"github.com/spf13/cobra"
)

var chargeCmd = &cobra.Command{
	Use:     "charge",
	Aliases: []string{"rev"},
	Short:   "Creates a lightning invoice",
	Long:    `Creates a lightning invoice and returns the QR code and the invoice itself`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res := charge.Charge(args[0], args[1])
		fmt.Println(res.InvoiceQrCode)
		fmt.Println(res.Invoice)
	},
}

func init() {
	rootCmd.AddCommand(chargeCmd)
}
