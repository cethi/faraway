package proof_of_work

import (
	"github.com/umahmood/hashcash"
)

type HashCashPOW struct {
}

func (h *HashCashPOW) Compute(payload string) (string, error) {
	hash, computationError := h.compute(payload)

	for computationError != nil {
		if computationError == hashcash.ErrSolutionFail {
			return "", computationError
		}
		hash, computationError = h.Compute(payload)
	}
	return hash, computationError
}

func (h *HashCashPOW) compute(payload string) (string, error) {
	hc, err := hashcash.New(
		&hashcash.Resource{
			Data: payload,
		},
		nil,
	)
	if err != nil {
		return "", err
	}

	solution, err := hc.Compute()
	if err != nil {
		return "", err
	}
	return solution, nil
}

func (h *HashCashPOW) Validate(task string, solution string) (bool, error) {
	hc, err := hashcash.New(
		&hashcash.Resource{
			ValidatorFunc: func(decoded string) bool {
				return decoded == task
			},
		},
		nil,
	)
	if err != nil {
		return false, err
	}

	return hc.Verify(solution)
}
