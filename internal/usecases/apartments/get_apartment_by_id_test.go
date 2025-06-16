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

type GetApartmentByIDSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock_apartments.MockapartmentsRepository
	uc   *apartmentsUsecase
	ctx  context.Context
}

func (s *GetApartmentByIDSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_apartments.NewMockapartmentsRepository(s.ctrl)
	s.uc = NewApartmentsUsecase(s.repo)
	s.ctx = context.Background()
}

func (s *GetApartmentByIDSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *GetApartmentByIDSuite) TestOK() {
	want := domain.Apartment{ID: 92}
	s.repo.EXPECT().GetApartmentByID(s.ctx, int64(92)).Return(want, nil).Times(1)
	got, err := s.uc.GetApartmentByID(s.ctx, 92)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *GetApartmentByIDSuite) TestError() {
	s.repo.EXPECT().GetApartmentByID(s.ctx, int64(51)).Return(domain.Apartment{}, errors.New("not found")).Times(1)
	got, err := s.uc.GetApartmentByID(s.ctx, 51)
	s.Error(err)
	s.Equal(domain.Apartment{}, got)
}

func TestGetApartmentByIDSuite(t *testing.T) {
	suite.Run(t, new(GetApartmentByIDSuite))
}
