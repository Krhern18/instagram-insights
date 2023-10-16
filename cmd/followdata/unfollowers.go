package followdata

import (
	"github.com/cecobask/instagram-insights/pkg/instagram"
	"github.com/cecobask/instagram-insights/pkg/instagram/followdata"
	"github.com/spf13/cobra"
)

const CommandNameUnfollowers = "unfollowers"

func NewUnfollowersCommand() *cobra.Command {
	return &cobra.Command{
		Use:   CommandNameUnfollowers,
		Short: "Retrieve a list of users who are not following you back",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := instagram.NewOptions(instagram.OutputTable)
			return followdata.NewHandler().Unfollowers(opts)
		},
	}
}
