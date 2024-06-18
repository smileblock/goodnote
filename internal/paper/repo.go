package paper

import "goodnode/internal/correlation"

type Repo interface {
	Add(input *AddInput) correlation.Id
}
