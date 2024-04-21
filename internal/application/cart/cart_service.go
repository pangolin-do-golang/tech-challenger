package cart

import (
	"context"
	"errors"
	"github.com/pangolin-do-golang/tech-challenge/internal/domainerrors"
)

type Service struct {
	CartRepository         ICartRepository
	CartProductsRepository ICartProductRepository
}

func NewService(cartRepository ICartRepository, cartProductsRepository ICartProductRepository) IService {
	return &Service{
		CartRepository:         cartRepository,
		CartProductsRepository: cartProductsRepository,
	}
}

func (s *Service) LoadCart(ctx context.Context) (*Cart, error) {
	clientID := ctx.Value("client_id").(string)
	cart, err := s.CartRepository.Get(ctx, clientID)
	if err != nil {
		if !errors.Is(err, domainerrors.ErrRecordNotFound) {
			return nil, err
		}

		cart, err = s.CartRepository.Create(ctx)
		if err != nil {
			return nil, err
		}
	}

	return cart, nil

}

func (s *Service) AddProduct(ctx context.Context, product *Product) error {
	cart, err := s.LoadCart(ctx)
	if err != nil {
		return err
	}

	return s.CartProductsRepository.Create(ctx, cart.ID, product)
}
