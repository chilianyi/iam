// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package im

import (
	"time"

	"github.com/golang/protobuf/ptypes"

	"openpitrix.io/iam/pkg/pb/im"
)

type DBUser struct {
	UserId      string    `db:"user_id"`
	GroupId     string    `db:"group_id"`
	RoleId      string    `db:"role_id"`
	UserName    string    `db:"user_name"`
	Position    string    `db:"position"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	Password    string    `db:"password"`
	OldPassword string    `db:"old_password"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	Owner       string    `db:"owner"`
	OwnerPath   string    `db:"owner_path"`
	CreateTime  time.Time `db:"create_time"`
	StatusTime  time.Time `db:"status_time"`
	UpdateTime  time.Time `db:"update_time"`
}

func pbUserToDB(p *pbim.User) *DBUser {
	if p == nil {
		return new(DBUser)
	}
	var q = &DBUser{
		//UserId:      p.UserId,
		//GroupId:     p.GroupId,
		//RoleId:      p.RoleId,
		//UserName:    p.UserName,
		//Position:    p.Position,
		//Email:       p.Email,
		//PhoneNumber: p.PhoneNumber,
		//Password:    p.Password,
		//OldPassword: p.OldPassword,
		//Description: p.Description,
		//Status:      p.Status,
		//Owner:       p.Owner,
		//OwnerPath:   p.OwnerPath,
	}

	q.CreateTime, _ = ptypes.Timestamp(p.CreateTime)
	q.UpdateTime, _ = ptypes.Timestamp(p.UpdateTime)
	q.StatusTime, _ = ptypes.Timestamp(p.StatusTime)

	return q
}

func (p *DBUser) ToPb() *pbim.User {
	if p == nil {
		return new(pbim.User)
	}
	var q = &pbim.User{
		//UserId:      p.UserId,
		//GroupId:     p.GroupId,
		//RoleId:      p.RoleId,
		//UserName:    p.UserName,
		//Position:    p.Position,
		//Email:       p.Email,
		//PhoneNumber: p.PhoneNumber,
		//Password:    p.Password,
		//OldPassword: p.OldPassword,
		//Description: p.Description,
		//Status:      p.Status,
		//Owner:       p.Owner,
		//OwnerPath:   p.OwnerPath,
	}

	q.CreateTime, _ = ptypes.TimestampProto(p.CreateTime)
	q.UpdateTime, _ = ptypes.TimestampProto(p.UpdateTime)
	q.StatusTime, _ = ptypes.TimestampProto(p.StatusTime)

	return q
}

func (p *DBUser) ValidateForInsert() error {
	return nil
}
func (p *DBUser) ValidateForUpdate() error {
	return nil
}
