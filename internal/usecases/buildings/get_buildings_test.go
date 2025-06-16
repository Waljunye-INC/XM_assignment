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

type GetBuildingsSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock_buildings.MockbuildingsRepository
	uc   *buildingsUsecase
	ctx  context.Context
}

func (s *GetBuildingsSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_buildings.NewMockbuildingsRepository(s.ctrl)
	s.uc = NewBuildingsUsecase(s.repo)
	s.ctx = context.Background()
}

func (s *GetBuildingsSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *GetBuildingsSuite) TestOK() {
	want := []domain.Building{{ID: 1, Name: "Alpha", Address: "AAA"}}
	s.repo.EXPECT().GetBuildings(s.ctx).Return(want, nil).Times(1)
	got, err := s.uc.buildingsRepository.GetBuildings(s.ctx)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *GetBuildingsSuite) TestError() {
	s.repo.EXPECT().GetBuildings(s.ctx).Return(nil, errors.New("fail")).Times(1)
	got, err := s.uc.buildingsRepository.GetBuildings(s.ctx)
	s.Error(err)
	s.Nil(got)
}

func TestGetBuildingsSuite(t *testing.T) {
	suite.Run(t, new(GetBuildingsSuite))
}
