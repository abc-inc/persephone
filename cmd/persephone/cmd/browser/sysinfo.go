// Copyright 2022 The persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"strconv"
	"strings"

	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/event"
	"github.com/abc-inc/persephone/graph"
	. "github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
)

type DBInfo struct {
	Name    string
	Address string
	Role    string
	Status  string
	Default bool
	Error   string
}

func (d DBInfo) String() string {
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
		console.SetFormatter(DBInfo{}, func(i interface{}) (string, error) {
			db := i.(DBInfo)
			return strings.Join([]string{db.Name, db.Address, db.Role, db.Status,
				strconv.FormatBool(db.Default), db.Error}, sep), nil
		})
	})
}

func SysInfo() {
	console.Write(ListDBs())
}

func ListDBs() []DBInfo {
	t := graph.NewTypedTemplate[DBInfo](graph.GetConn())
	dbs, _ := MustTuple(t.Query("SHOW DATABASES", nil, func(rec *neo4j.Record) DBInfo {
		return DBInfo{
			Name:    MustOk(rec.Get("name")).(string),
			Address: MustOk(rec.Get("address")).(string),
			Role:    MustOk(rec.Get("role")).(string),
			Status:  MustOk(rec.Get("requestedStatus")).(string),
			Default: MustOk(rec.Get("default")).(bool),
			Error:   MustOk(rec.Get("error")).(string),
		}
	}))
	return dbs
}
