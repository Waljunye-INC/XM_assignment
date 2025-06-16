package apartments

import (
	"OMS_assignment/internal/usecases/apartments/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type DeleteApartmentSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock_apartments.MockapartmentsRepository
	uc   *apartmentsUsecase
	ctx  context.Context
}

func (s *DeleteApartmentSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_apartments.NewMockapartmentsRepository(s.ctrl)
	s.uc = NewApartmentsUsecase(s.repo)
	s.ctx = context.Background()
}

func (s *DeleteApartmentSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *DeleteApartmentSuite) TestOK() {
	s.repo.EXPECT().DeleteApartment(s.ctx, int64(31)).Return(nil).Times(1)
	err := s.uc.DeleteApartment(s.ctx, 31)
	s.NoError(err)
}

func (s *DeleteApartmentSuite) TestError() {
	s.repo.EXPECT().DeleteApartment(s.ctx, int64(45)).Return(errors.New("fail")).Times(1)
	err := s.uc.DeleteApartment(s.ctx, 45)
	s.Error(err)
}

func TestDeleteApartmentSuite(t *testing.T) {
	suite.Run(t, new(DeleteApartmentSuite))
}
