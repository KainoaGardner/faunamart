package cmd

import (
	"database/sql"
	"fmt"
	"github.com/KainoaGardner/faunamart/database"
	"github.com/KainoaGardner/faunamart/utils"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(winCmd)
}

var winCmd = &cobra.Command{
	Use:   "win",
	Short: "Select a winning ticket",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.Open()

		winId, err := utils.RandomTicket(3)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The winning ticket is %d\n", winId)
		showWinningTicket(db, winId)

	},
}

func showWinningTicket(db *sql.DB, winId int) {
	var ticket database.Ticket
	var found bool

	query := "SELECT * FROM tickets WHERE ticket = ?"
	statment, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := statment.Query(winId)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		found = true
		err := rows.Scan(&ticket.ID, &ticket.Ticket, &ticket.Name, &ticket.Date)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d: %d %s  %s\n", ticket.ID, ticket.Ticket, ticket.Name, ticket.Date)
	}
	if !found {
		fmt.Printf("No winners for %d\n", winId)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)

	}

}
