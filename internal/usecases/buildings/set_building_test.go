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

type SetBuildingSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock_buildings.MockbuildingsRepository
	uc   *buildingsUsecase
	ctx  context.Context
}

func (s *SetBuildingSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_buildings.NewMockbuildingsRepository(s.ctrl)
	s.uc = NewBuildingsUsecase(s.repo)
	s.ctx = context.Background()
}

func (s *SetBuildingSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *SetBuildingSuite) TestOK() {
	b := domain.Building{ID: 111, Name: "Дом", Address: "Улица, 1"}
	s.repo.EXPECT().SetBuilding(s.ctx, b).Return(nil).Times(1)
	err := s.uc.buildingsRepository.SetBuilding(s.ctx, b)
	s.NoError(err)
}

func (s *SetBuildingSuite) TestError() {
	b := domain.Building{ID: 111, Name: "Дом", Address: "Улица, 1"}
	s.repo.EXPECT().SetBuilding(s.ctx, b).Return(errors.New("fail")).Times(1)
	err := s.uc.buildingsRepository.SetBuilding(s.ctx, b)
	s.Error(err)
}

func TestSetBuildingSuite(t *testing.T) {
	suite.Run(t, new(SetBuildingSuite))
}
