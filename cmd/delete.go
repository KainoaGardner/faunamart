package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/KainoaGardner/faunamart/database"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&all, "all", "a", false, "Delete all tickets")
}

var all bool

var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "delete [ID] delete a ticket by id",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		db := database.Open()
		var id int
		var err error

		if len(args) > 0 {

			id, err = strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
		} else if !all {
			log.Fatal("Must have an id")

		}

		if all {

			deleteAllTickets(db)
		} else {

			deleteTicket(db, id)
		}

	},
}

func deleteTicket(db *sql.DB, id int) {
	statment, err := db.Prepare("DELETE FROM tickets WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	result, err := statment.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if affected == 1 {

		fmt.Printf("id %d was deleted\n", id)
	} else {
		fmt.Printf("id %d not found\n", id)

	}

}

func deleteAllTickets(db *sql.DB) {
	statment, err := db.Prepare("DELETE FROM tickets")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statment.Exec()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted all tickets\n")

}
