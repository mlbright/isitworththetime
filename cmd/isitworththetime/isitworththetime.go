package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mlbright/isitworththetime/pkg/timebudget"
)

func main() {
	var saved time.Duration
	var frequency uint
	var interval time.Duration
	var span time.Duration
	var users uint

	flag.DurationVar(&saved, "saved", 0, "time shaved off")
	flag.UintVar(&frequency, "frequency", 1, fmt.Sprintf("number of occurrences per time interval (default: %d)", 1))
	flag.DurationVar(&interval, "interval", timebudget.DefaultWorkDay, fmt.Sprintf("time interval, (default: %s)", timebudget.HumanizeDuration(timebudget.DefaultWorkDay)))
	flag.DurationVar(&span, "span", timebudget.DefaultSpan, fmt.Sprintf("total time over which the intervals occur, (default: %s)", timebudget.HumanizeDuration(timebudget.DefaultSpan)))
	flag.UintVar(&users, "users", timebudget.DefaultUsers, fmt.Sprintf("number of users affected (default %d)", timebudget.DefaultUsers))

	flag.Usage = func() {
		flagSet := flag.CommandLine
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		order := []string{"saved", "frequency", "interval", "span", "users"}
		for _, name := range order {
			f := flagSet.Lookup(name)
			fmt.Printf("\t--%s\t%s\n", f.Name, f.Usage)
		}
	}

	flag.Parse()

	t, err := timebudget.TimeBudget(&saved, frequency, &interval, &span, users)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(timebudget.HumanizeDuration(t))
}
