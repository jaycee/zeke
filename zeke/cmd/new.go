package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Make a new zettelkasten note",
	Run:   newZk,
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func newZk(cmd *cobra.Command, args []string) {
	now := time.Now()
	filename := fmt.Sprintf("%v.md", now.Format("20060102150405"))
	contents := renderTemplate(now)

	f, err := os.Create(filename)
	if err != nil {
		handleErr(err, "Couldn't create new file")
	}
	defer f.Close()
	f.WriteString(contents)
	f.Sync()

	fmt.Println(filename)
}

func handleErr(err error, msg string) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%v: %v\n", msg, err))
	os.Exit(2)
}

func renderTemplate(now time.Time) string {
	output := []string{
		"---",
		"title:",
		fmt.Sprintf("date: %v", now.Format("2006-01-02 15:04")),
		"tags:",
		"backlinks:",
		"---",
		"\n",
		"From SOURCE:",
		"\n",
		"  > QUOTE",
	}
	return strings.Join(output, "\n")
}
