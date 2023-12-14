package delivery

import "golang-test/usecase"

type delivery struct {
	usecase usecase.Usecase
}

func NewDelivery(usecase usecase.Usecase) *delivery {
	return &delivery{
		usecase: usecase,
	}

}
