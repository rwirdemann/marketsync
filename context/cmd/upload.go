package cmd

import (
	"github.com/rwirdemann/marketsync/application/usecases"
	"github.com/rwirdemann/marketsync/ports/out"
	"github.com/spf13/cobra"
)

func NewUploadCmd(catalog out.Catalog, marketplace out.Marketplace) *cobra.Command {
	return &cobra.Command{
		Use:   "upload",
		Short: "upload inventory",
		RunE: func(cmd *cobra.Command, args []string) error {
			return usecases.Upload(catalog, marketplace)
		},
	}
}
