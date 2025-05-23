// Package database is the root command for civo database.
package database

import (
	"errors"

	"github.com/spf13/cobra"
)

var firewallID, networkID, size, updatedName, software, softwareVersion string
var nodes int

// DBCmd is the root command for the db subcommand
var DBCmd = &cobra.Command{
	Use:     "database",
	Aliases: []string{"db", "databases"},
	Short:   "Manage Civo Database ",
	Long:    `Create, update, delete, and list Civo Databases.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		if err != nil {
			return err
		}
		return errors.New("a valid subcommand is required")
	},
}

func init() {
	DBCmd.AddCommand(dbListCmd)
	DBCmd.AddCommand(dbCreateCmd)
	DBCmd.AddCommand(dbUpdateCmd)
	DBCmd.AddCommand(dbShowCmd)
	DBCmd.AddCommand(dbDeleteCmd)
	DBCmd.AddCommand(dbCredentialCmd)
	DBCmd.AddCommand(dbSizeCmd)
	DBCmd.AddCommand(dbEngineCmd)
	DBCmd.AddCommand(dbBackupCmd)
	DBCmd.AddCommand(dbRestoreCmd)
	DBCmd.AddCommand(dbVersionListCmd)

	dbCredentialCmd.Flags().BoolVarP(&connectionString, "connection-string", "c", false, "show the connection string for the database")

	dbCreateCmd.Flags().IntVarP(&nodes, "nodes", "", 1, "the number of nodes for the database")
	dbCreateCmd.Flags().StringVarP(&firewallID, "firewall", "", "", "the firewall to use for the database")
	dbCreateCmd.Flags().StringVarP(&networkID, "network", "n", "", "the network to use for the database")
	dbCreateCmd.Flags().StringVarP(&rulesFirewall, "firewall-rules", "u", "", "the firewall rules to use for the database")
	dbCreateCmd.Flags().StringVarP(&size, "size", "s", "g3.db.small", "the size of the database. You can list available DB sizes by `civo size list -s database`")
	dbCreateCmd.Flags().StringVarP(&software, "software", "m", "MySQL", "the software to use for the database. One of: MySQL, PostgreSQL. Please make sure you use the correct capitalisation.")
	dbCreateCmd.Flags().StringVarP(&softwareVersion, "version", "v", "", "the version of the software to use for the database.")
	dbCreateCmd.Flags().BoolVarP(&waitDatabase, "wait", "w", false, "a simple flag (e.g. --wait) that will cause the CLI to spin and wait for the database to be ACTIVE")

	dbUpdateCmd.Flags().IntVarP(&nodes, "nodes", "", 0, "the number of nodes for the database")
	dbUpdateCmd.Flags().StringVarP(&firewallID, "firewall", "", "", "the firewall to use for the database")
	dbUpdateCmd.Flags().StringVarP(&updatedName, "name", "n", "", "the new name for the database")

	dbRestoreCmd.Flags().StringVarP(&backup, "backup", "b", "", "the backup name which you can restore database")
	dbRestoreCmd.Flags().StringVarP(&restoreName, "name", "n", "", "name of the restore")
	dbRestoreCmd.MarkFlagRequired("backup")
	dbRestoreCmd.MarkFlagRequired("name")
}
