package usecase

import "github.com/renanmoreirasan/fc-go-clean-architecture/internal/entity"

type ListOrderInputDTO struct{}

type ListOrderOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute(input ListOrderInputDTO) (ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return ListOrderOutputDTO{}, err
	}

	ordersDTOList := []OrderOutputDTO{}

	for _, order := range orders {
		orderDTO := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		ordersDTOList = append(ordersDTOList, orderDTO)
	}

	output := ListOrderOutputDTO{
		Orders: ordersDTOList,
	}

	return output, nil
}
