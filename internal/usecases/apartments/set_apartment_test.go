package apartments

import (
	"context"
	"errors"
	"testing"

	"OMS_assignment/internal/domain"
	"OMS_assignment/internal/usecases/apartments/mocks"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type SetApartmentSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock_apartments.MockapartmentsRepository
	uc   *apartmentsUsecase
	ctx  context.Context
}

func (s *SetApartmentSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_apartments.NewMockapartmentsRepository(s.ctrl)
	s.uc = NewApartmentsUsecase(s.repo)
	s.ctx = context.Background()
}

func (s *SetApartmentSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *SetApartmentSuite) TestOK() {
	apt := domain.Apartment{ID: 9, Number: "1A"}
	s.repo.EXPECT().SetApartment(s.ctx, apt).Return(nil).Times(1)
	err := s.uc.SetApartment(s.ctx, apt)
	s.NoError(err)
}

func (s *SetApartmentSuite) TestError() {
	apt := domain.Apartment{ID: 9, Number: "1A"}
	s.repo.EXPECT().SetApartment(s.ctx, apt).Return(errors.New("fail")).Times(1)
	err := s.uc.SetApartment(s.ctx, apt)
	s.Error(err)
}

func TestSetApartmentSuite(t *testing.T) {
	suite.Run(t, new(SetApartmentSuite))
}
