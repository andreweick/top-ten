package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/urfave/cli/v2"

	"top-ten/internal/lists"
)

func main() {
	app := &cli.App{
		Name:  "top-ten",
		Usage: "Display random David Letterman Top 10 lists",
		Commands: []*cli.Command{
			{
				Name:  "random",
				Usage: "Display a random Top 10 list",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "ascii",
						Usage: "Display output using ASCII characters only (no colors)",
					},
				},
				Action: func(c *cli.Context) error {
					return showRandomList(c.Context, c.Bool("ascii"))
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func showRandomList(ctx context.Context, ascii bool) error {
	// Set ASCII mode if requested
	if ascii {
		lipgloss.SetColorProfile(termenv.Ascii)
	}
	
	service, err := lists.NewService(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize service: %w", err)
	}

	list, err := service.GetRandomList()
	if err != nil {
		return fmt.Errorf("failed to get random list: %w", err)
	}

	printList(list)
	return nil
}

func printList(list lists.TopTenList) {
	// Style definitions
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Align(lipgloss.Center).
		Padding(0, 2).
		Margin(1, 0)

	dateStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#626262")).
		Align(lipgloss.Center).
		Margin(0, 0, 1, 0)

	numberStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF6B6B")).
		Bold(true)

	// Create the title box
	titleBox := titleStyle.Render(list.Title)
	width := lipgloss.Width(titleBox)

	// Style the date to match the width of the title box
	styledDate := dateStyle.Width(width).Render(list.Date)

	// Print the styled title and date
	fmt.Printf("\n%s\n", titleBox)
	fmt.Printf("%s\n\n", styledDate)

	// Print the list items with styled numbers
	for _, item := range list.Items {
		number := strings.Split(item, ".")[0]
		content := strings.Join(strings.Split(item, ".")[1:], ".")
		content = strings.TrimSpace(content)

		// Right-align number within 2 character width
		formattedNumber := fmt.Sprintf("%2s.", number)
		styledNumber := numberStyle.Render(formattedNumber)
		fmt.Printf("  %s %s\n", styledNumber, content)
	}

	fmt.Println()
}
