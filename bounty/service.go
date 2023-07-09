package bounty

import "golang.org/x/net/context"

func createBounty(ctx context.Context, dto CreateBountyRequest) (CreateBountyResponse, error) {
	var (
		res CreateBountyResponse
	)

	row, err := insertNewBounty(ctx, dto)
	if err != nil {
		return res, err
	}

	return CreateBountyResponse{
		ID:     row.ID,
		Name:   dto.Name,
		Bounty: dto.Bounty,
	}, nil
}
