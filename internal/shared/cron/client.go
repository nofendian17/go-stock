package cron

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type CronClient interface {
	AddJob(spec string, cmd func()) (cron.EntryID, error)
	RemoveJob(cron.EntryID)
	Start()
	Stop() context.Context
	Entries() []cron.Entry
}

type cronClient struct {
	c *cron.Cron
}

func NewCronClient(tz string) CronClient {
	location, err := time.LoadLocation(tz)
	if err != nil {
		log.Printf("Failed to load location %s: %v. Defaulting to UTC.", tz, err)
		location = time.UTC
	}
	log.Printf("Initializing cron client with timezone: %s", location)

	return &cronClient{
		c: cron.New(cron.WithLocation(location)),
	}
}

func (cc *cronClient) AddJob(spec string, cmd func()) (cron.EntryID, error) {
	id, err := cc.c.AddFunc(spec, cmd)
	if err != nil {
		log.Printf("Error adding cron job with spec %q: %v", spec, err)
	} else {
		log.Printf("Added cron job: spec=%q, id=%d", spec, id)
	}
	return id, err
}

func (cc *cronClient) RemoveJob(id cron.EntryID) {
	log.Printf("Removing cron job with id=%d", id)
	cc.c.Remove(id)
}

func (cc *cronClient) Start() {
	log.Println("Starting cron scheduler")
	cc.c.Start()
}

func (cc *cronClient) Stop() context.Context {
	log.Println("Stopping cron scheduler")
	return cc.c.Stop()
}

func (cc *cronClient) Entries() []cron.Entry {
	entries := cc.c.Entries()
	log.Printf("Fetching %d cron entries", len(entries))
	return entries
}
