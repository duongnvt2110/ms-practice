package container

import (
	"fmt"
	"log"
	"os"

	"ms-practice/noti-service/pkg/config"
	"ms-practice/noti-service/pkg/model"
	"ms-practice/noti-service/pkg/provider/notifier"
	"ms-practice/noti-service/pkg/repository"
	"ms-practice/noti-service/pkg/usecase"
	nkafka "ms-practice/noti-service/pkg/util/kafka"
	"ms-practice/noti-service/pkg/worker"
	sharedgorm "ms-practice/pkg/db/gorm_client"
)

type Container struct {
	Cfg       *config.Config
	Repo      repository.NotificationRepository
	Usecase   usecase.NotificationUsecase
	Messaging *nkafka.NotificationMessaging
	Workers   []*worker.Worker
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	db, err := sharedgorm.NewGormClient(cfg.Mysql)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to init db", err)
		os.Exit(1)
	}

	repo := repository.NewNotificationRepository(db)
	uc := usecase.NewNotificationUsecase(repo, cfg)
	messaging := nkafka.NewNotificationMessaging(cfg.Kafka)

	emailProvider := &notifier.EmailProvider{From: cfg.SMTP.From}
	pushProvider := &notifier.PushProvider{}
	inAppProvider := &notifier.InAppProvider{}

	workers := []*worker.Worker{
		{Channel: model.ChannelEmail, Repo: repo, Provider: emailProvider, Cfg: cfg},
		{Channel: model.ChannelPush, Repo: repo, Provider: pushProvider, Cfg: cfg},
		{Channel: model.ChannelInApp, Repo: repo, Provider: inAppProvider, Cfg: cfg},
	}

	log.Printf("notification container initialized")
	return &Container{
		Cfg:       cfg,
		Repo:      repo,
		Usecase:   uc,
		Messaging: messaging,
		Workers:   workers,
	}
}
