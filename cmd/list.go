package cmd

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/KainoaGardner/faunamart/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&sort, "sort", "s", "", "Sort type [ticket,name,date]")
	listCmd.Flags().BoolVarP(&desc, "desc", "d", false, "Sort type [ticket,name,date]")
}

var sort string = "date"
var desc bool
var name string = ""

var listCmd = &cobra.Command{
	Use:   "list [NAME]",
	Short: "list [NAME] List all active tickets",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db := database.Open()

		if len(args) != 0 {
			name = args[0]

		}
		listAll(db)

	},
}

func listAll(db *sql.DB) {
	var ticket database.Ticket

	query := "SELECT * FROM tickets"
	if name != "" {
		query += fmt.Sprintf(" WHERE name ='%s'", name)

	}

	query += " ORDER BY ?"
	if desc {

		query += " DESC"
	}

	statment, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := statment.Query(sort)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ticket.ID, &ticket.Ticket, &ticket.Name, &ticket.Date)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %d %s  %s\n", ticket.ID, ticket.Ticket, ticket.Name, ticket.Date)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)

	}

}
