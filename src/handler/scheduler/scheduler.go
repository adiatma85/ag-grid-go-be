package scheduler

import (
	"sync"
)

var (
	once = &sync.Once{}
)

type Interface interface {
	Run()
	TriggerScheduler(name string) error
}

type scheduler struct {
	// cron       *gocron.Scheduler
	// conf       config.SchedulerConfig
	// metaconf   config.ApplicationMeta
	// log        log.Interface
	// auth       auth.Interface
	// uc         *usecase.Usecases
	// template   email.TemplateInterface
	// instrument instrument.Interface
}

type InitParam struct {
	// Conf     config.SchedulerConfig
	// MetaConf config.ApplicationMeta
	// Log      log.Interface
	// Auth     auth.Interface
	// Uc       *usecase.Usecases
	// Tmpl     email.TemplateInterface
	// Instr    instrument.Interface
}
