/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"kafkapipe/pipe"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var from, to string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafkapipe",
	Short: "kafkapipe reads message from one kafka, then writes them to another kafka",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		runPipe()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kafkapipe.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVar(&from, "from", "from.toml", "kafka config that read from")
	rootCmd.Flags().StringVar(&to, "to", "to.toml", "kafka config that wrote to")
}

func runPipe() {
	var err error

	var consumerMap kafka.ConfigMap
	viperFrom := viper.New()
	viperFrom.SetConfigFile(from)
	err = viperFrom.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viperFrom.UnmarshalKey("consumerConfig", &consumerMap)
	if err != nil {
		log.Fatal("consuerConfig unmarshal error", err)
	}

	topicFrom := viperFrom.GetString("topicFrom.name")
	if len(topicFrom) == 0 {
		log.Fatal("topicFrom length 0")
	}

	var producerMap kafka.ConfigMap
	viperTo := viper.New()
	viperTo.SetConfigFile(to)
	err = viperTo.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viperTo.UnmarshalKey("producerConfig", &producerMap)
	if err != nil {
		log.Fatal("producerConfig unmarshal error")
	}

	topicTo := viperTo.GetString("topicTo.name")
	if len(topicTo) == 0 {
		log.Fatal("topicTo length 0")
	}

	pipe.Start(&consumerMap, topicFrom, &producerMap, topicTo)
}
