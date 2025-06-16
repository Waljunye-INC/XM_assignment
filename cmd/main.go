package main

import (
	"OMS_assignment/internal/contract/oapi"
	apartments3 "OMS_assignment/internal/contract/oapi/apartments"
	buildings3 "OMS_assignment/internal/contract/oapi/buildings"
	"OMS_assignment/internal/repository/apartments"
	"OMS_assignment/internal/repository/buildings"
	apartments2 "OMS_assignment/internal/usecases/apartments"
	buildings2 "OMS_assignment/internal/usecases/buildings"
	"OMS_assignment/libs/application"
	listeners2 "OMS_assignment/libs/listeners"
	"database/sql"
	"flag"
	"os"
	"os/signal"

	"OMS_assignment/cmd/config"
	"OMS_assignment/utils"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type app interface {
	Run(stopChan chan os.Signal)
}

func main() {
	cfgFile := flag.String("cfg-file", "", "define .env file")
	flag.Parse()

	cfg, err := config.LoadFromEnv(cfgFile)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	dbConn, err := sql.Open("postgres", utils.PGDSN(
		cfg.DBName(),
		cfg.DBHost(),
		cfg.DBPort(),
		cfg.DBUsername(),
		cfg.DBPassword(),
	))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer dbConn.Close()

	if err = run(cfg, dbConn); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func run(cfg *config.Config, db *sql.DB) error {
	app, err := build(cfg, db)
	if err != nil {
		return err
	}

	log.Info().Msg("build finished, starting app")
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	app.Run(stopChan)
	return nil
}

func build(cfg *config.Config, db *sql.DB) (app, error) {
	buildingsRepository := buildings.NewBuildingsRepository(db)
	apartmentsRepository := apartments.NewApartmentsRepository(db)

	buildingsUsecase := buildings2.NewBuildingsUsecase(buildingsRepository)
	apartmentsUsecase := apartments2.NewApartmentsUsecase(apartmentsRepository)

	buildingsContract := buildings3.NewBuildingsContract(buildingsUsecase)
	apartmentsContract := apartments3.NewApartmentsContract(apartmentsUsecase)

	openAPI := oapi.New(nil, buildingsContract, apartmentsContract)
	listeners := map[int]listeners2.PortListener{
		cfg.PublicApiPort(): openAPI,
	}

	result := application.New(listeners)
	return result, nil
}
