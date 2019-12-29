package cmd

import (
	"errors"
	"fmt"

	utils "pumba/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "use add or list to add config or view config",
	Long: `
		pumba config add
		pumba config view
	`,
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the current configuration",
	Long:  `Use list to check the existing configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		keys := viper.AllKeys()
		for _, value := range keys {
			fmt.Printf("%s : %s\n", value, viper.GetString(value))
		}
	},
}
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new config the current configuration",
	Long: `I can allow you to configure your application again
	
Set --secretAccessKey: Aws Secret keys
Set --accessKeyId: Aws Access Keys
Set --pem__filepath: Location where all the pemFile is stored
Set --default_region: AWS default region`,
	Run: func(cmd *cobra.Command, args []string) {
		err := setConfig()
		utils.PanicWithError(err)

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(listCmd, addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// take user input and set Configuration
func setConfig() error {
	return errors.New("a error")
}
