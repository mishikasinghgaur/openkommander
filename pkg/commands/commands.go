package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RegisterCommands(rootCmd *cobra.Command) {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Connect to a Kafka cluster",
		Run: func(cmd *cobra.Command, args []string) {
			loginCommand()
		},
	}

	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "End the current session",
		Run: func(cmd *cobra.Command, args []string) {
			logoutCommand()
		},
	}

	sessionCmd := &cobra.Command{
		Use:   "session",
		Short: "Display current session information",
		Run: func(cmd *cobra.Command, args []string) {
			sessionInfoCommand()
		},
	}

	metadataCmd := &cobra.Command{
		Use:   "metadata",
		Short: "Display cluster information",
		Run: func(cmd *cobra.Command, args []string) {
			metadataCommand()
		},
	}

	brokersCmd := &cobra.Command{
		Use:   "brokers",
		Short: "Broker management commands",
	}

	listBrokersCmd := &cobra.Command{
		Use:   "list",
		Short: "List all brokers",
		Run: func(cmd *cobra.Command, args []string) {
			brokersListCommand()
		},
	}

	brokersCmd.AddCommand(listBrokersCmd)

	rootCmd.AddCommand(brokersCmd)

	serverCmd := restServerCommands()

	topicsCmd := topicsCommands()

	rootCmd.AddCommand(
		loginCmd,
		logoutCmd,
		sessionCmd,
		metadataCmd,
		topicsCmd,
		serverCmd,
		restCmd,
	)
}

func restServerCommands() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "REST server commands",
	}

	startServerCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the REST server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Starting REST server...\n")
		},
	}

	stopServerCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop the REST server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Stopping REST server...\n")
		},
	}

	serverCmd.AddCommand(startServerCmd, stopServerCmd)

	return serverCmd
}

func topicsCommands() *cobra.Command {
	topicsCmd := &cobra.Command{
		Use:   "topics",
		Short: "Topic management commands",
	}

	createTopicCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new topic",
		Run: func(cmd *cobra.Command, args []string) {
			createTopicCommand()
		},
	}

	deleteTopicCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a topic",
		Run: func(cmd *cobra.Command, args []string) {
			deleteTopicCommand()
		},
	}

	listTopicsCmd := &cobra.Command{
		Use:   "list",
		Short: "List all topics",
		Run: func(cmd *cobra.Command, args []string) {
			listTopicsCommand()
		},
	}

	topicsCmd.AddCommand(createTopicCmd, deleteTopicCmd, listTopicsCmd)

	return topicsCmd
}
