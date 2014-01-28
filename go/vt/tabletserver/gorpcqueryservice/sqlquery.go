// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gorpcqueryservice

import (
	mproto "github.com/youtube/vitess/go/mysql/proto"
	"github.com/youtube/vitess/go/rpcwrap"
	rpcproto "github.com/youtube/vitess/go/rpcwrap/proto"
	"github.com/youtube/vitess/go/vt/tabletserver"
	"github.com/youtube/vitess/go/vt/tabletserver/proto"
)

type SqlQuery struct {
	server *tabletserver.SqlQuery
}

var sqlQuery *SqlQuery

func (sq *SqlQuery) GetSessionId(sessionParams *proto.SessionParams, sessionInfo *proto.SessionInfo) error {
	return sq.server.GetSessionId(sessionParams, sessionInfo)
}

func (sq *SqlQuery) Begin(context *rpcproto.Context, session *proto.Session, txInfo *proto.TransactionInfo) error {
	return sq.server.Begin(context, session, txInfo)
}

func (sq *SqlQuery) Commit(context *rpcproto.Context, session *proto.Session, noOutput *string) error {
	return sq.server.Commit(context, session, noOutput)
}

func (sq *SqlQuery) Rollback(context *rpcproto.Context, session *proto.Session, noOutput *string) error {
	return sq.server.Rollback(context, session, noOutput)
}

func (sq *SqlQuery) Execute(context *rpcproto.Context, query *proto.Query, reply *mproto.QueryResult) error {
	return sq.server.Execute(context, query, reply)
}

func (sq *SqlQuery) StreamExecute(context *rpcproto.Context, query *proto.Query, sendReply func(reply interface{}) error) error {
	return sq.server.StreamExecute(context, query, sendReply)
}

func (sq *SqlQuery) ExecuteBatch(context *rpcproto.Context, queryList *proto.QueryList, reply *proto.QueryResultList) error {
	return sq.server.ExecuteBatch(context, queryList, reply)
}

func init() {
	tabletserver.SqlQueryRegisterFunctions = append(tabletserver.SqlQueryRegisterFunctions, func(sq *tabletserver.SqlQuery) {
		sqlQuery = &SqlQuery{sq}
		rpcwrap.RegisterAuthenticated(sqlQuery)
	})
}
