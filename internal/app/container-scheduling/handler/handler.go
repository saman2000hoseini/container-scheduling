package handler

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/model"

	"github.com/labstack/echo/v4"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/config"
	"github.com/sony/sonyflake"
)

type ServerJobHandler struct {
	cfg  config.Config
	jobs chan model.Job
}

func ServerNewJobHandler(cfg config.Config, jobs chan model.Job) *ServerJobHandler {
	return &ServerJobHandler{cfg: cfg, jobs: jobs}
}

func (h *ServerJobHandler) UserRequest(c echo.Context) error {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	re, _ := regexp.Compile("(<.*?>)")
	res := re.FindAllString(string(b), -1)
	if res == nil {
		return c.NoContent(http.StatusBadRequest)
	}

	dest := res[len(res)-1][1 : len(res[len(res)-1])-1]

	for _, i := range res[:len(res)-1] {
		i := strings.Split(i, ",")
		job := model.Job{Id: id, Operation: i[0][1:], Source: i[1][:len(i[1])-1], Destination: dest}
		h.jobs <- job
	}

	return c.String(http.StatusOK, "Your request with the id: "+strconv.FormatUint(id, 10)+" delivered to scheduler\n Enter new request")
}
