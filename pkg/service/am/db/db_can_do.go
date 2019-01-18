// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package db

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"openpitrix.io/iam/pkg/internal/funcutil"
	pbam "openpitrix.io/iam/pkg/pb/am"
	pbim "openpitrix.io/iam/pkg/pb/im"
	"openpitrix.io/logger"
)

func (p *Database) CanDo(ctx context.Context, req *pbam.CanDoRequest) (*pbam.CanDoResponse, error) {
	logger.Infof(ctx, funcutil.CallerName(1))

	type DBCanDo struct {
		Url       string
		UrlMethod string
	}

	var query = sqlCanDo
	var rows = []DBCanDo{}

	err := p.DB.Raw(query, req.UserId, req.Url, req.UrlMethod).Scan(&rows).Error
	if err != nil {
		logger.Warnf(ctx, "%v", query)
		logger.Warnf(ctx, "%+v", err)
		return nil, err
	}
	if len(rows) == 0 {
		err := status.Errorf(codes.PermissionDenied, "disable")
		logger.Warnf(ctx, "%+v", err)
		return nil, err
	}

	// get owner path from IM server
	ownerPath, err := p.GetOwnerPathByUserId(ctx, req.UserId)
	if len(rows) == 0 {
		logger.Warnf(ctx, "%+v", err)
		return nil, err
	}

	// get access path

	_ = ownerPath

	// 1. get role list by user_id
	// 2. get action list by role_id
	// 3. check action rule

	// try get OwnerPath from user_path

	reply := &pbam.CanDoResponse{
		UserId:     req.UserId,
		OwnerPath:  ownerPath,
		AccessPath: "AccessPath-todo",
	}

	return reply, nil
}

func (p *Database) GetOwnerPathByUserId(ctx context.Context, userId string) (string, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", p.cfg.ImHost, p.cfg.ImPort), grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pbim.NewAccountManagerClient(conn)
	reply, err := client.GetGroupsByUserId(ctx, &pbim.UserId{Uid: userId})
	if err != nil {
		return "", err
	}
	fmt.Println(reply.GetValue())

	// no group
	if len(reply.Value) == 0 {
		return "", nil
	}

	// take firest group
	ownerPath := reply.Value[0].GroupPath

	// OK
	return ownerPath, nil
}