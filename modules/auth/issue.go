// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"
	"reflect"

	"github.com/codegangsta/martini"

	"github.com/gogits/binding"

	"github.com/gogits/gogs/modules/base"
	"github.com/gogits/gogs/modules/log"
)

type CreateIssueForm struct {
	IssueName   string `form:"name" binding:"Required;MaxSize(50)"`
	RepoId      int64  `form:"repoid" binding:"Required"`
	MilestoneId int64  `form:"milestoneid" binding:"Required"`
	AssigneeId  int64  `form:"assigneeid"`
	Labels      string `form:"labels"`
	Content     string `form:"content"`
}

func (f *CreateIssueForm) Name(field string) string {
	names := map[string]string{
		"IssueName":   "Issue name",
		"RepoId":      "Repository ID",
		"MilestoneId": "Milestone ID",
	}
	return names[field]
}

func (f *CreateIssueForm) Validate(errors *binding.Errors, req *http.Request, context martini.Context) {
	if req.Method == "GET" || errors.Count() == 0 {
		return
	}

	data := context.Get(reflect.TypeOf(base.TmplData{})).Interface().(base.TmplData)
	data["HasError"] = true
	AssignForm(f, data)

	if len(errors.Overall) > 0 {
		for _, err := range errors.Overall {
			log.Error("CreateIssueForm.Validate: %v", err)
		}
		return
	}

	validate(errors, data, f)
}
