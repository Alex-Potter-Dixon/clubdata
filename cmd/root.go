package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(helloCmd)
}

var rootCmd = &cobra.Command{
	Use:   "parser",
	Short: "Clubcard data parser",
	Long: `A easy eay to consume and analyse your club card data
		   via a CLI`,


}

var helloCmd = &cobra.Command {
	Use: "hello",
	Short: "Command that says hello world",
	Long: "Prints hellow world, amazing I know..",
	Run: func(cmd *cobra.Command, args []string){
		printHello(cmd, args)
	},

}

func printHello(cmd *cobra.Command, args []string){

	fmt.Println("Hello world!")
}


// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
