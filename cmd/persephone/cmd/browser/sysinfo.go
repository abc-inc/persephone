package cmd

import (
	"strconv"
	"strings"

	"github.com/abc-inc/persephone/event"
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	. "github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
)

type dbInfo struct {
	Name    string
	Address string
	Role    string
	Status  string
	Default bool
	Error   string
}

func (d dbInfo) String() string {
	return d.Address + "/" + d.Name
}

var SysinfoCmd = &cobra.Command{
	Use:   ":sysinfo",
	Short: "Print system information",
	Run:   func(cmd *cobra.Command, args []string) { SysInfo() },
}

func init() {
	event.Subscribe(event.FormatEvent{}, func(e event.FormatEvent) {
		sep := e.Sep
		format.SetFormatter(dbInfo{}, func(i interface{}) (string, error) {
			db := i.(dbInfo)
			return strings.Join([]string{db.Name, db.Address, db.Role, db.Status,
				strconv.FormatBool(db.Default), db.Error}, sep), nil
		})
	})
}

func SysInfo() {
	t := graph.NewTypedTemplate[dbInfo](graph.GetConn())
	dbs, _ := MustTuple(t.Query("SHOW DATABASES", nil, func(rec *neo4j.Record) dbInfo {
		return dbInfo{
			Name:    MustOk(rec.Get("name")).(string),
			Address: MustOk(rec.Get("address")).(string),
			Role:    MustOk(rec.Get("role")).(string),
			Status:  MustOk(rec.Get("requestedStatus")).(string),
			Default: MustOk(rec.Get("default")).(bool),
			Error:   MustOk(rec.Get("error")).(string),
		}
	}))

	format.Writeln(dbs)
}