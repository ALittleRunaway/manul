/*
Copyright © 2022 Maria Petrova marycool674@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"manul/converter"
	"os"
)

var (
	cfgFile    string
	isInvert   bool
	width      int
	height     int
	outputFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "manul [options...] <filename>",
	Example: "manul -w 50 -o ~/myimageascii.txt myimage.png",
	Version: "1.0.0",
	Short:   "ASCII-image converter",
	Long: `Manul is a CLI tool for converting your images to ASCII-symbols.
Supported image formats: PNG, JPG, GPEG.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			converter.ConvertToAscii(args)
		} else {
			fmt.Println("try 'manul --help' or 'manul -h' for more information")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolVarP(&isInvert, "invert", "i", false, "Invert image colors")
	rootCmd.Flags().IntVarP(&width, "width", "w", 0, "Width in pixels. Default is 300. "+
		"If both width and height are provided, height is ignored")
	rootCmd.Flags().IntVarP(&height, "height", "n", 0, "Height in pixels. Default is "+
		"200. If both width and height are provided, height is ignored")
	rootCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "ascii.txt", "Output file "+
		"name. Default is 'ascii.txt'")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".manul")
	}

	viper.Set("invert", isInvert)
	viper.Set("width", width)
	viper.Set("height", height)
	viper.Set("outputFile", outputFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
