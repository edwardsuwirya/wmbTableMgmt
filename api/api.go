package api

import (
	"database/sql"
	"github.com/edwardsuwirya/wmbTableMgmt/config"
	"github.com/edwardsuwirya/wmbTableMgmt/delivery"
	"github.com/edwardsuwirya/wmbTableMgmt/entity"
	"github.com/edwardsuwirya/wmbTableMgmt/manager"
	"log"
)

type Server interface {
	Run()
}

type server struct {
	config  *config.Config
	infra   manager.Infra
	usecase manager.UseCaseManager
}

func NewApiServer() Server {
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUseCaseManger(repo)
	return &server{
		config:  appConfig,
		infra:   infra,
		usecase: usecase,
	}
}

func (s *server) Run() {
	if !(s.config.RunMigration == "Y" || s.config.RunMigration == "y") {
		db, _ := s.infra.SqlDb().DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(db)
		delivery.NewServer(s.config.RouterEngine, s.usecase)
		err := s.config.RouterEngine.Run(s.config.ApiBaseUrl)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		db := s.infra.SqlDb()
		err := db.AutoMigrate(&entity.CustomerTable{}, &entity.CustomerTableTransaction{})
		db.Unscoped().Where("id like ?", "%%").Delete(entity.CustomerTable{})
		db.Model(&entity.CustomerTable{}).Save([]entity.CustomerTable{
			{
				ID: "A01",
			},
			{
				ID: "A02",
			},
			{
				ID: "A03",
			},
		})

		if err != nil {
			log.Fatalln(err)
		}
	}
}
