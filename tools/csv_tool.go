package tools

import (
	"context"
	"encoding/csv"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

func DataCsv() mcp.Tool {

	newTool := mcp.NewTool("dataByCsv",
		mcp.WithDescription("Busca de dados através de CSV"),
	)

	return newTool

}

func DataCsvHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	f, err := os.Open("E:\\GoLangProjects\\teste-mcp\\dados.csv")

	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	value := recordsToString(records)

	return mcp.NewToolResultText(value), nil

}

func recordsToString(records [][]string) string {
	var lines []string
	for _, row := range records {
		line := strings.Join(row, ",") // Junta os campos de cada linha com vírgula
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n") // Junta as linhas com quebra de linha
}
