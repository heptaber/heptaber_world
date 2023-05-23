package cron

import (
	"heptaber/auth/domain/repository"
	"log"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var c *cron.Cron

func init() {
	c = cron.New(cron.WithSeconds())
}

func SetUpDeleteAllExpiredVerificationCodesJob(db *gorm.DB) {
	c.AddFunc("0 30 3 * * *", func() {
		log.Print("started running cron job deleting expired verification codes")
		vcr := repository.NewVerificationCodeRepository(db)
		err := vcr.DeleteAllExpired()
		if err != nil {
			log.Print("error while deleting all expired verification codes: ", err)
		}
	})
	c.Start()
}
