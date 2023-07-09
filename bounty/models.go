package bounty

type (
	Bounty struct {
		ID     string
		Name   string
		Bounty float32
	}

	CreateBountyRequest struct {
		Name   string  `json:"name" validate:"required"`
		Bounty float32 `json:"bounty" validate:"required"`
	}

	CreateBountyResponse struct {
		ID     string  `json:"id"`
		Name   string  `json:"name"`
		Bounty float32 `json:"bounty"`
	}
)
