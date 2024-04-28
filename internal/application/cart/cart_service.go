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

func (s *Service) LoadCart(ctx context.Context, clientID string) (*Cart, error) {
	cart, err := s.CartRepository.Get(clientID)
	if err != nil {
		if !errors.Is(err, domainerrors.ErrRecordNotFound) {
			return nil, err
		}

		cart, err = s.CartRepository.Create(clientID)
		if err != nil {
			return nil, err
		}
	}

	return cart, nil

}

func (s *Service) AddProduct(ctx context.Context, product *Product) error {
	cart, err := s.LoadCart(ctx, product.ClientID)
	if err != nil {
		return err
	}

	// TODO verificar se produto já tá no carrinho/colocar índice de unicidade
	return s.CartProductsRepository.Create(ctx, cart.ID, product)
}

func (s *Service) RemoveProduct(ctx context.Context, clientID string, productID string) error {
	cart, err := s.LoadCart(ctx, clientID)
	if err != nil {
		return err
	}

	products, err := s.CartProductsRepository.GetByCartID(ctx, cart.ID)
	if err != nil {
		return err
	}

	for _, product := range products {
		if product.ProductID == productID {
			return s.CartProductsRepository.DeleteByProductID(ctx, cart.ID, productID)
		}
	}

	return nil
}

func (s *Service) EditProduct(ctx context.Context, product *Product) error {
	cart, err := s.LoadCart(ctx, product.ClientID)
	if err != nil {
		return err
	}

	return s.CartProductsRepository.UpdateProductByProductID(ctx, cart.ID, product.ProductID, product)
}
