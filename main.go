package main

import (
	"cmdtrans/cmd"
	"cmdtrans/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	transCmd := newTransCmd()
	if err := transCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var sourceLanguage string
var targetLanguage string
var showLanguageCodes bool

func newTransCmd() *cobra.Command {
	transCmd := &cobra.Command{
		Use:   "fy",
		Short: "fy Command-line Interface v0.1.0",
	}
	transCmd.Flags().StringVarP(&sourceLanguage, "source", "s", "auto", "Choose what language you want to translate from, default as auto.")
	transCmd.Flags().StringVarP(&targetLanguage, "target", "t", "zh", "Choose what language you want to translate to, default as zh.")
	transCmd.Flags().BoolVarP(&showLanguageCodes, "langcodes", "l", false, "Show the language codes suppored.")
	transCmd.Run = func(ccmd *cobra.Command, args []string) {
		if showLanguageCodes {
			cmd.ShowLanguageCodes()
			return
		}
		config.ParamsObj = &config.Params{
			SourceLang: sourceLanguage,
			TargetLang: targetLanguage,
		}
		if len(args) == 0 && !showLanguageCodes {
			transCmd.Help()
			return
		}
		cmd.DealTrans(args)
	}
	return transCmd
}
