package commands

import (
	"bufio"
	"fmt"
	"os"
	"simple-cli/expenses"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// GetInput obtiene el input de la consola de forma interactiva. Devuelve el string ingresado y un posible error si ocurre.
func GetInput() (string, error) {
	fmt.Println("-> ")
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	str = strings.Replace(str, "\r\n", "", 1)
	return str, nil
}

// ShowInConsole imprime la lista de precios introducidos en la consola, así como los detalles adicionales como el total, máximo, mínimo y promedio.
func ShowInConsole(expensesList []float32) {
	fmt.Println(contentString(expensesList))
	fmt.Println("========================")
}

// contentString convierte la lista de precios en un string legible para ser mostrado en consola o exportado a otro formato.
func contentString(expensesList []float32) string {
	builder := strings.Builder{}
	max, min, avg, total := expensesDetails(expensesList)

	fmt.Println("")
	for i, expense := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))
		if i == len(expensesList)-1 {
			fmt.Println("========================")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", max))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", min))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", avg))
		}
	}
	return builder.String()
}

// expensesDetails calcula el valor máximo, mínimo, promedio y total de la lista de gastos. Devuelve estos valores para ser mostrados o utilizados.
func expensesDetails(expensesList []float32) (max, min, average, total float32) {
	if len(expensesList) == 0 {
		return
	}
	min = expenses.Min(expensesList...)
	max = expenses.Max(expensesList...)
	total = expenses.Sum(expensesList...)
	average = expenses.Average(expensesList...)
	return
}

// Export permite exportar los datos a txt
func Export(fileName string, list []float32) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = w.WriteString(contentString(list))
	if err != nil {
		return err
	}

	return w.Flush()
}
