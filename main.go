package main

import (
	"flag"
	"simple-cli/commands"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// main Inicializa el programa
func main() {
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000Z"
	var expenses []float32
	var export string

	flag.StringVar(&export, "export", "-", "Export details to .txt")
	flag.Parse()

	for {
		input, err := commands.GetInput()
		if err != nil {
			log.Panic().Msg(err.Error())
		}
		if input == "cls" {
			break
		}
		expense, err := strconv.ParseFloat(input, 32)
		if err != nil {
			log.Error().AnErr("No se puede convertir el valor a numero", err).Msg("Error al Parseo")
			continue
		}
		expenses = append(expenses, float32(expense))
	}
	if export == "" {
		commands.ShowInConsole(expenses)
	} else {
		commands.Export(export, expenses)
	}
	commands.ShowInConsole(expenses)
	log.Info().Msg("Terminal cerrada")

}
