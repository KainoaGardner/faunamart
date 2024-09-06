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
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [NAME] [TICKET]",
	Short: "add [NAME] [TICKET] Add ticket",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {

		db := database.Open()

		name := args[0]
		ticketNum := 0
		if len(args) > 1 {
			num, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal(err)
			}
			ticketNum = num

		}

		addTicket(db, name, ticketNum)

	},
}

func addTicket(db *sql.DB, name string, ticketNum int) {
	statment, err := db.Prepare("INSERT INTO tickets (ticket,name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statment.Exec(name, ticketNum)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ticket Added", ticketNum, name)

}
