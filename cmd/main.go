package main

import (
	"XM_assignment/internal/contract/oapi/authcontract"
	"XM_assignment/internal/events"
	authrepository "XM_assignment/internal/repositories/auth"
	"XM_assignment/internal/usecases/auth"
	"database/sql"
	"flag"
	"os"
	"os/signal"

	"XM_assignment/cmd/config"
	"XM_assignment/internal/contract/oapi"
	"XM_assignment/internal/contract/oapi/companiescontract"
	companiesrepository "XM_assignment/internal/repositories/companies"
	"XM_assignment/internal/usecases/companies"
	"XM_assignment/libs/application"
	listeners2 "XM_assignment/libs/listeners"
	"XM_assignment/utils"

	_ "github.com/go-sql-driver/mysql"
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

	dbConn, err := sql.Open("mysql", utils.MysqlDSN(
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
	eventReciever := events.NewEventReciever()
	eventsProducer := events.NewProducer(eventReciever, cfg.KafkaTopic(), cfg.Brokers())

	companiesRepository := companiesrepository.NewRepository(db)
	authRepository := authrepository.NewRepository(db)

	companiesUseCase := companies.NewUseCase(companiesRepository, eventReciever)
	authUsecase := auth.NewUseCase(cfg.JWTKey(), authRepository)

	companiesContract := companiescontract.NewContract(companiesUseCase)
	authContract := authcontract.NewContract(authUsecase)
	openAPI := oapi.New(nil, cfg.JWTKey(), companiesContract, authContract)
	listeners := map[int]listeners2.PortListener{
		cfg.PublicApiPort(): openAPI,
	}

	backgroundWorkers := []listeners2.BackgroundWorker{
		eventsProducer,
	}

	result := application.New(listeners, backgroundWorkers)
	return result, nil
}
