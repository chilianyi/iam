// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package db

import (
	"crypto/rand"
	"regexp"
	"strings"
	"unicode"

	"github.com/golang/protobuf/ptypes/timestamp"

	"openpitrix.io/iam/pkg/internal/base58"
)

func isValidSearchWord(name string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)
	return re.MatchString(name)
}
func isValidSortKey(name string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)
	return re.MatchString(name)
}

func isValidGids(ids ...string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)
	for _, id := range ids {
		if !re.MatchString(id) {
			return false
		}
	}
	return true
}
func isValidUids(ids ...string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)
	for _, id := range ids {
		if !re.MatchString(id) {
			return false
		}
	}
	return true
}
func isValidNames(ids ...string) bool {
	return true
}
func isValidStatus(ids ...string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)
	for _, id := range ids {
		if !re.MatchString(id) {
			return false
		}
	}
	return true
}

func isValidEmails(emails ...string) bool {
	for _, v := range emails {
		if v == "" {
			return false
		}
		if idx := strings.IndexByte(v, '@'); idx <= 0 || idx >= len(v) {
			return false
		}
		if strings.Count(v, "@") != 1 {
			return false
		}
		for _, c := range v {
			if unicode.IsSpace(c) {
				return false
			}
		}
	}
	return true
}

func isValidPhoneNumbers(phoneNumbers ...string) bool {
	var re = regexp.MustCompile(`^[0-9]+$`)
	for _, id := range phoneNumbers {
		if !re.MatchString(id) {
			return false
		}
	}
	return true
}

func isZeroTimestamp(x *timestamp.Timestamp) bool {
	if x == nil {
		return true
	}
	if x.Seconds == 0 && x.Nanos == 0 {
		return true
	}
	return false
}

func genGid() string {
	buf := make([]byte, 8)
	rand.Read(buf)
	s := string(base58.EncodeBase58(buf))
	return "gid-" + s[:8]
}

func genUid() string {
	buf := make([]byte, 8)
	rand.Read(buf)
	s := string(base58.EncodeBase58(buf))
	return "uid-" + s[:8]
}
func genRoleId() string {
	buf := make([]byte, 8)
	rand.Read(buf)
	s := string(base58.EncodeBase58(buf))
	return "role-" + s[:8]
}
func genXid() string {
	buf := make([]byte, 8)
	rand.Read(buf)
	s := string(base58.EncodeBase58(buf))
	return "xid-" + s[:8]
}
func genId(prefix string, maxLen int) string {
	if prefix == "" {
		prefix = "xid-"
	}
	if maxLen <= 0 {
		maxLen = 12
	}

	if maxLen <= len(prefix) {
		maxLen += len(prefix)
	}

	buf := make([]byte, maxLen-len(prefix))
	rand.Read(buf)
	s := string(base58.EncodeBase58(buf))
	return prefix + s[:len(buf)]
}
