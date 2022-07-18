package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mcuadros/go-version"
	"github.com/sivchari/hashicorp"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "list-remote",
	Short: "Display versions of all available terraform versions",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(context.Background(), ""); err != nil {
			log.Fatalf("Failed to list-remote err = %s\n", err.Error())
		}
	},
}

var c *hashicorp.Client

func init() {
	rootCmd.AddCommand(runCmd)

	c = hashicorp.New()
}

var list []string

func run(ctx context.Context, after string) error {
	l, af, err := listRemote(ctx, after)
	if err != nil {
		return err
	}
	list = append(list, l...)
	if l != nil {
		run(ctx, af)
	}
	if after != "" {
		return nil
	}
	version.Sort(list)
	out := make([]string, 0, len(list))
	for i := len(list) - 1; i >= 0; i-- {
		out = append(out, list[i])
	}
	fmt.Fprint(os.Stdout, strings.Join(out, "\n"))
	return nil
}

func listRemote(ctx context.Context, after string) ([]string, string, error) {
	ls, err := c.ListReleases(ctx, "terraform", &hashicorp.ListReleasesParam{
		After: after,
		Limit: 20,
	})
	if err != nil {
		return nil, "", err
	}
	rs := make([]string, len(ls.Releases))
	for i, l := range ls.Releases {
		rs[i] = l.Version
	}
	if len(ls.Releases) == 0 {
		return nil, "", nil
	}
	return rs, ls.Releases[len(ls.Releases)-1].TimestampCreated.Format(time.RFC3339), nil
}
