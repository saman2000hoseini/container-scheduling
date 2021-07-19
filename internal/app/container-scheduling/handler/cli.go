package handler

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/model"
	"github.com/sirupsen/logrus"

	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/config"
	"github.com/sony/sonyflake"
)

type CLIJobHandler struct {
	cfg  config.Config
	jobs chan model.Job
}

func CLINewJobHandler(cfg config.Config, jobs chan model.Job) *CLIJobHandler {
	return &CLIJobHandler{cfg: cfg, jobs: jobs}
}

func (h *CLIJobHandler) UserRequest(request string) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()

	re, _ := regexp.Compile("(<.*?>)")
	res := re.FindAllString(string(request), -1)
	if res == nil {
		logrus.Info("Invalid Request!")
		return
	}

	dest := res[len(res)-1][1 : len(res[len(res)-1])-1]

	for _, i := range res[:len(res)-1] {
		i := strings.Split(i, ",")
		job := model.Job{Id: id, Operation: i[0][1:], Source: i[1][:len(i[1])-1], Destination: dest}
		h.jobs <- job
	}

	logrus.Info("Your request with the id: " + strconv.FormatUint(id, 10) + " delivered to scheduler!")
}
