package buildings

import (
	"context"
	"errors"
	"testing"

	"OMS_assignment/internal/domain"
	"OMS_assignment/internal/usecases/buildings/mocks"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type GetBuildingByIDSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock_buildings.MockbuildingsRepository
	uc   *buildingsUsecase
	ctx  context.Context
}

func (s *GetBuildingByIDSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_buildings.NewMockbuildingsRepository(s.ctrl)
	s.uc = NewBuildingsUsecase(s.repo)
	s.ctx = context.Background()
}

func (s *GetBuildingByIDSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *GetBuildingByIDSuite) TestOK() {
	want := domain.Building{ID: 5, Name: "Beta", Address: "Address"}
	s.repo.EXPECT().GetBuildingByID(s.ctx, int64(5)).Return(want, nil).Times(1)
	got, err := s.uc.buildingsRepository.GetBuildingByID(s.ctx, 5)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *GetBuildingByIDSuite) TestError() {
	s.repo.EXPECT().GetBuildingByID(s.ctx, int64(7)).Return(domain.Building{}, errors.New("not found")).Times(1)
	got, err := s.uc.buildingsRepository.GetBuildingByID(s.ctx, 7)
	s.Error(err)
	s.Equal(domain.Building{}, got)
}

func TestGetBuildingByIDSuite(t *testing.T) {
	suite.Run(t, new(GetBuildingByIDSuite))
}
