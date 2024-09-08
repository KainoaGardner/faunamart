package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/KainoaGardner/faunamart/database"
	"github.com/KainoaGardner/faunamart/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&times, "times", "t", "1", "Amount of times to add ticket")
}

var times string

var addCmd = &cobra.Command{
	Use:   "add [NAME] [TICKET]",
	Short: "add [NAME] [TICKET] Add ticket",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {

		db := database.Open()
		var ticketNum int
		times, err := strconv.Atoi(times)
		if err != nil {
			log.Fatal(err)
		}

		name := args[0]

		for i := 0; i < times; i++ {

			if len(args) > 1 {
				num, err := strconv.Atoi(args[1])
				if err != nil {
					log.Fatal(err)
				}
				ticketNum = num

			} else {
				var err error
				ticketNum, err = utils.RandomTicket(3)
				if err != nil {
					log.Fatal(err)
				}

			}

			addTicket(db, name, ticketNum)
		}

	},
}

func addTicket(db *sql.DB, name string, ticketNum int) {
	statment, err := db.Prepare("INSERT INTO tickets (ticket,name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statment.Exec(ticketNum, name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ticket Added", ticketNum, name)

}
