// Copyright 2019 HenryYee.
//
// Licensed under the AGPL, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.gnu.org/licenses/agpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"google.golang.org/grpc"
	"time"
)

type mysql struct {
	Host     string
	User     string
	Password string
	Db       string
	Port     string
}

type general struct {
	SecretKey string
	Host      string
	Hours     time.Duration
	GrpcAddr  string
}

type DbInfo struct {
	Host     string
	User     string
	Password string
	Port     string
	Db       string
}

type Config struct {
	General general
	Mysql   mysql
}

var C Config

var JWT = ""

var Host = ""

var Grpc = ""

var Conn *grpc.ClientConn

var GloPer CoreGlobalConfiguration

var GloLdap Ldap

var GloOther Other

var GloMessage Message
