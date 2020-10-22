package main

import (
	"github.com/jinzhu/gorm"
	"github.com/viki-org/go-utils/healthcheck"
	"github.com/viki-org/go-utils/queue"
)

type gormChecker struct {
	db *gorm.DB
}

type queryResponse struct {
	Response string
}

// NewGormChecker is a constructor for the Custom Gorm Checker
func newGormChecker(db *gorm.DB) healthcheck.Checker {
	return gormChecker{db}
}

type vikiQueueChecker struct {
	queueClient queue.MessageQueueClient
}

func newVikiQueueChecker(q queue.MessageQueueClient) healthcheck.Checker {
	return vikiQueueChecker{q}
}

// Check is a custom checker for GORM
func (g gormChecker) Check() healthcheck.Health {
	resp := queryResponse{}
	h := healthcheck.NewHealth()

	if g.db == nil {
		h.Down().AddInfo("error", "Empty resource")
		return h
	}

	output := g.db.Raw("SELECT 1 as response").Scan(&resp)
	if output.Error != nil {
		h.Down().AddInfo("error", output.Error.Error())
		return h
	}

	output = g.db.Raw("SELECT VERSION() as response").Scan(&resp)
	if output.Error != nil {
		h.Down().AddInfo("error", output.Error.Error())
		return h
	}

	h.Up().AddInfo("version", resp.Response)
	return h
}

// Check is a custom checker for GORM
func (q vikiQueueChecker) Check() healthcheck.Health {

	const (
		statusString = "status"
		okString     = "ok"
	)

	h := healthcheck.NewHealth()

	if q.queueClient == nil {
		h.Down().AddInfo("error", "Empty resource")
		return h
	}

	err := q.queueClient.Publish(queue.CreateMessage, "healthcheck", "1-healthcheck", "healthcheck", "")
	if err != nil {
		h.Down().AddInfo("error", err.Error())
		return h
	}

	err = q.queueClient.Publish(queue.DeleteMessage, "healthcheck", "1-healthcheck", "healthcheck", "")
	if err != nil {
		h.Down().AddInfo("error", err.Error())
		return h
	}

	h.Up().AddInfo(statusString, okString)
	return h
}
