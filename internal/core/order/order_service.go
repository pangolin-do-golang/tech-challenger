package order

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/cart"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/product"
	"github.com/pangolin-do-golang/tech-challenge/internal/errutil"
	"os"
)

type Service struct {
	HttpClient     resty.Client
	CartService    cart.IService
	ProductService *product.Service
}

func NewOrderService(cartService cart.IService, productService *product.Service) IOrderService {
	client := resty.New()
	client.SetBaseURL(os.Getenv("ORDER_SERVICE_URL"))
	return &Service{
		CartService:    cartService,
		ProductService: productService,
		HttpClient:     *client,
	}
}

func (s *Service) Get(id uuid.UUID) (*Order, error) {

	var order Order

	_, err := s.HttpClient.R().
		SetPathParams(map[string]string{
			"id": id.String(),
		}).
		SetResult(&order).
		Get("/orders/{id}")

	if err != nil {
		if errors.Is(err, errutil.ErrRecordNotFound) {
			return nil, errutil.NewBusinessError(err, "order not found")
		}

		return nil, err
	}

	return &order, nil
}

func (s *Service) GetAll() ([]Order, error) {
	var orders []Order

	_, err := s.HttpClient.R().
		SetResult(&orders).
		Get("/orders/{id}")

	if err != nil {

		return nil, err
	}

	return orders, nil
}

type UpdateOrderPayload struct {
	Status string `json:"status" binding:"required" example:"paid"`
}

func (s *Service) Update(order *Order) (*Order, error) {
	var result Order

	_, err := s.HttpClient.R().
		SetPathParams(map[string]string{
			"id": order.ID.String(),
		}).
		SetBody(UpdateOrderPayload{
			Status: order.Status,
		}).
		SetResult(&result).
		Patch("/orders/{id}")

	if err != nil {
		return nil, err
	}

	return &result, nil
}

type CreateOrderPayload struct {
	ClientID uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
}

func (s *Service) Create(clientID uuid.UUID) (*Order, error) {
	var result Order

	_, err := s.HttpClient.R().
		SetBody(CreateOrderPayload{
			ClientID: clientID,
		}).
		SetResult(&result).
		Post("/orders")

	if err != nil {
		return nil, err
	}

	return &result, nil
}
