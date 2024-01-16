package cmd

import (
	"github.com/rwirdemann/marketsync/application/usecases"
	"github.com/rwirdemann/marketsync/ports/out"
	"github.com/spf13/cobra"
)

func NewUploadCmd(catalog out.Catalog) *cobra.Command {
	return &cobra.Command{
		Use:   "upload",
		Short: "upload inventory",
		RunE: func(cmd *cobra.Command, args []string) error {
			usecases.Upload(catalog)
			return nil
		},
	}
}
