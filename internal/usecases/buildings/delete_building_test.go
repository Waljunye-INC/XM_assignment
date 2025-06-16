package buildings

import (
	"context"
	"errors"
	"testing"

	"OMS_assignment/internal/usecases/buildings/mocks"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type DeleteBuildingSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mock_buildings.MockbuildingsRepository
	uc   *buildingsUsecase
	ctx  context.Context
}

func (s *DeleteBuildingSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_buildings.NewMockbuildingsRepository(s.ctrl)
	s.uc = NewBuildingsUsecase(s.repo)
	s.ctx = context.Background()
}

func (s *DeleteBuildingSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *DeleteBuildingSuite) TestOK() {
	s.repo.EXPECT().DeleteBuilding(s.ctx, int64(22)).Return(nil).Times(1)
	err := s.uc.buildingsRepository.DeleteBuilding(s.ctx, 22)
	s.NoError(err)
}

func (s *DeleteBuildingSuite) TestError() {
	s.repo.EXPECT().DeleteBuilding(s.ctx, int64(23)).Return(errors.New("fail")).Times(1)
	err := s.uc.buildingsRepository.DeleteBuilding(s.ctx, 23)
	s.Error(err)
}

func TestDeleteBuildingSuite(t *testing.T) {
	suite.Run(t, new(DeleteBuildingSuite))
}
