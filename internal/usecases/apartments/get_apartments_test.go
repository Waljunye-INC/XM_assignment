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

type GetApartmentsSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock_apartments.MockapartmentsRepository
	uc   *apartmentsUsecase
	ctx  context.Context
}

func (s *GetApartmentsSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_apartments.NewMockapartmentsRepository(s.ctrl)
	s.uc = NewApartmentsUsecase(s.repo)
	s.ctx = context.Background()
}

func (s *GetApartmentsSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *GetApartmentsSuite) TestGetApartments_OK() {
	want := []domain.Apartment{{ID: 1, Number: "88"}}
	s.repo.EXPECT().GetApartments(s.ctx).Return(want, nil).Times(1)
	got, err := s.uc.GetApartments(s.ctx)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *GetApartmentsSuite) TestGetApartments_Error() {
	s.repo.EXPECT().GetApartments(s.ctx).Return(nil, errors.New("db")).Times(1)
	got, err := s.uc.GetApartments(s.ctx)
	s.Error(err)
	s.Empty(got)
}

func TestGetApartmentsSuite(t *testing.T) {
	suite.Run(t, new(GetApartmentsSuite))
}
