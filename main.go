package main

import (
	"fmt"
	"os"

	filter_csv "github.com/jeffersonmss/whc-cli/csv"

	"github.com/spf13/cobra" // Librery used to create a cli application
)

func main() {
	// Open the CSV file
	sites, err := filter_csv.ReadCSV()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		panic(err)
	}

	rootCmd := FilterCLIByCategoryORState(sites)

	// Run main cli command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error running the command:", err)
		os.Exit(1)
	}
}

// Filter sites and print Name and category
func FilterCLIByCategoryORState(sites []filter_csv.WHSite) *cobra.Command {
	var category, state string

	var rootCmd = &cobra.Command{
		Use:   "whc",
		Short: "Return Name and category from sites on WHC list",
		Run: func(cmd *cobra.Command, args []string) {
			result := sites
			noHasFilter := hasnotFilter(category, state, result)
			if noHasFilter {
				return
			}

			result = filter_csv.FilterByCategoryORState(sites, category, state)
			fmt.Println("Filter by category or state:")
			for _, site := range result {
				fmt.Println("Name:", site.NameEN)
				fmt.Println("Category:", site.Category)
			}
		},
	}

	rootCmd.Flags().StringVarP(&category, "category", "c", "", "Category to filter, values: Cultural | Natural | Mixed")
	rootCmd.Flags().StringVarP(&state, "state", "s", "", "Country to filter")
	return rootCmd
}

// Print name and category from all sites
func hasnotFilter(category string, state string, result []filter_csv.WHSite) bool {
	if len(category) == 0 && len(state) == 0 {
		fmt.Println("Printing all sites")
		for _, site := range result {
			fmt.Println("Name:", site.NameEN)
			fmt.Println("Category:", site.Category)
		}
		return true
	}
	return false
}
