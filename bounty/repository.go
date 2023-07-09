package bounty

import (
	"bounty/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/rs/zerolog/log"
)

func insertNewBounty(ctx context.Context, dto CreateBountyRequest) (Bounty, error) {
	var bounty Bounty

	query, args, _ := sq.Insert("bounties").
		Columns("name", "bounties").
		Values(dto.Name, dto.Bounty).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	log.Debug().Str("query", query).Any("args", args).Send()

	err := pgxscan.Get(ctx, db.DB, &bounty, query, args...)
	if err != nil {
		log.Error().Err(err).Type("dto", dto).Msg("error inserting new bounty")
	}

	return bounty, err
}
