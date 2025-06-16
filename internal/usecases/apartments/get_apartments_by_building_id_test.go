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

type GetApartmentsByBuildingIDSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock_apartments.MockapartmentsRepository
	uc   *apartmentsUsecase
	ctx  context.Context
}

func (s *GetApartmentsByBuildingIDSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_apartments.NewMockapartmentsRepository(s.ctrl)
	s.uc = NewApartmentsUsecase(s.repo)
	s.ctx = context.Background()
}

func (s *GetApartmentsByBuildingIDSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *GetApartmentsByBuildingIDSuite) TestOK() {
	want := []domain.Apartment{{ID: 3, BuildingID: 10}}
	s.repo.EXPECT().GetApartmentsByBuildingID(s.ctx, int64(10)).Return(want, nil).Times(1)
	got, err := s.uc.GetApartmentsByBuildingID(s.ctx, 10)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *GetApartmentsByBuildingIDSuite) TestError() {
	s.repo.EXPECT().GetApartmentsByBuildingID(s.ctx, int64(13)).Return(nil, errors.New("fail")).Times(1)
	got, err := s.uc.GetApartmentsByBuildingID(s.ctx, 13)
	s.Error(err)
	s.Empty(got)
}

func TestGetApartmentsByBuildingIDSuite(t *testing.T) {
	suite.Run(t, new(GetApartmentsByBuildingIDSuite))
}
