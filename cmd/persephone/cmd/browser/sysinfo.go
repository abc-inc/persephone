// Copyright 2022 The Persephone authors
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

	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/internal"
	"github.com/abc-inc/roland/graph"
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

func NewCmdSysInfo(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ":sysinfo",
		Short: "Print system information",
		Run:   func(cmd *cobra.Command, args []string) { SysInfo() },
	}

	console.OnFormatChange(func(i console.FormatInfo) {
		sep := i.Sep
		console.SetFormatter(DBInfo{}, func(i any) (string, error) {
			db := i.(DBInfo)
			return strings.Join([]string{db.Name, db.Address, db.Role, db.Status,
				strconv.FormatBool(db.Default), db.Error}, sep), nil
		})
	})

	return cmd
}

func SysInfo() {
	console.Write(ListDBs())
}

func ListDBs() []DBInfo {
	t := graph.NewTemplate[DBInfo](graph.GetConn())
	r := graph.Request{Query: "SHOW DATABASES"}
	dbs, _ := internal.MustTuple(t.Query(r, func(rec *neo4j.Record) DBInfo {
		return DBInfo{
			Name:    internal.MustOk(rec.Get("name")).(string),
			Address: internal.MustOk(rec.Get("address")).(string),
			Role:    internal.MustOk(rec.Get("role")).(string),
			Status:  internal.MustOk(rec.Get("requestedStatus")).(string),
			Default: internal.MustOk(rec.Get("default")).(bool),
			Error:   internal.MustOk(rec.Get("error")).(string),
		}
	}))
	return dbs
}
