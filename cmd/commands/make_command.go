package commands

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/wesleysnt/finance-api/cmd/commands/stubs"
)

var makeMigrationCmd = &cobra.Command{
	Use:   "make:migration <file_name>",
	Short: "Create a migration file",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		if fileName == "" {
			cmd.Help()
			return
		}

		time := time.Now().Format("20060102150405")

		tableName, mType, valid := parseMigrationName(fileName)

		CreateUp(fileName, time, tableName, mType, valid)
		CreateDown(fileName, time, tableName, mType, valid)

		color.Greenln("Migration file created")
	},
}

func init() {
	gobaseCommand.AddCommand(makeMigrationCmd)
}

// make migration
func CreateUp(fileName, time, table, mType string, stub bool) {

	fileName1 := fmt.Sprintf("database/migrations/%v_%v.up.sql", time, fileName)

	f, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err)
		color.Redln(err.Error())
		return
	}
	if stub {
		if mType == "create" {
			_, err = f.WriteString(stubs.PostgresqlStubs{}.CreateUp(fileName, table))
			if err != nil {
				fmt.Println(err)
				color.Redln(err.Error())
				f.Close()
				return
			}
		} else if mType == "update" {
			_, err = f.WriteString(stubs.PostgresqlStubs{}.UpdateUp(fileName, table))
			if err != nil {
				fmt.Println(err)
				color.Redln(err.Error())
				f.Close()
				return
			}
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		color.Redln(err.Error())
		return
	}
}

func CreateDown(fileName, time, table, mType string, stub bool) {

	fileName1 := fmt.Sprintf("database/migrations/%v_%v.down.sql", time, fileName)

	f, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err)
		color.Redln(err.Error())
		return
	}

	if stub {
		if mType == "create" {
			_, err = f.WriteString(stubs.PostgresqlStubs{}.CreateDown(fileName, table))
			if err != nil {
				fmt.Println(err)
				color.Redln(err.Error())
				f.Close()
				return
			}
		} else if mType == "update" {
			_, err = f.WriteString(stubs.PostgresqlStubs{}.UpdateDown(fileName, table))
			if err != nil {
				fmt.Println(err)
				color.Redln(err.Error())
				f.Close()
				return
			}
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		color.Redln(err.Error())
		return
	}
}

func parseMigrationName(input string) (string, string, bool) {
	// Split the string by '_'
	words := strings.Split(input, "_")

	// Check if the first and last words match the criteria
	if len(words) >= 3 && (words[0] == "create" || words[0] == "update") && words[len(words)-1] == "table" {
		// Extract the table name
		tableName := strings.Join(words[1:len(words)-1], "_")
		return tableName, words[0], true
	}

	// If the criteria are not met, return an empty string and false
	return "", "", false
}
