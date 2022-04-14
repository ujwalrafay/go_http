/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// httpServerCmd represents the httpServer command
var httpServerCmd = &cobra.Command{
	Use:   "httpServer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		students()
		fmt.Println("httpServer called")
	},
}

func init() {
	rootCmd.AddCommand(httpServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func students() {
	http.HandleFunc("/", student)
	http.ListenAndServe(":8010", nil)
}

func student(w http.ResponseWriter, req *http.Request) {
	fmt.Println("called student")
	switch req.Method {
	case "GET":
		fmt.Fprintln(w, "got the get method")
	case "POST":
		fmt.Fprintln(w, "got the post method")
	case "PUT":
		fmt.Fprintln(w, "got the put method")
	}

}
