package script_archive_drop_reports

import (
	_ "net/http/pprof"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func run(ctx *cli.Context, deps CommandDeps, dateStr string) error {
	log.Info().Str("date", dateStr).Msg("running script")

	var err error

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return errors.Wrap(err, "failed to parse date")
	}

	if err = deps.DropReportArchiveService.ArchiveByDate(ctx.Context, date); err != nil {
		return errors.Wrap(err, "failed to run archiveDropReports")
	}

	log.Info().Msg("script finished")

	return nil
}