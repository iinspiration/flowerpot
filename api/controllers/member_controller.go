package controllers

import (
	"flowerpot/models"
	"flowerpot/models/requests"
	"flowerpot/models/responses"
	"flowerpot/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gitlab.com/1hopin/go-module/constants"
	"gitlab.com/1hopin/go-module/utils"
)

type MemberController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
	ActiveList(c *gin.Context)
}

type memberController struct {
	memberRepository repositories.MemberRepository
}

func InitMemberController(memberRepository repositories.MemberRepository) MemberController {
	return &memberController{
		memberRepository,
	}
}

func (controller *memberController) Create(c *gin.Context) {
	var controllerName = "Create"
	utils.LoggerInfo(controllerName, "start", nil, nil, false)
	var req requests.Member
	err := c.Bind(&req)

	insertModel := models.Member{
		Name:         req.Name,
		Lastname:     req.Lastname,
		MobileNumber: req.MobileNumber,
		Email:        req.Email,
		Address:      req.Address,
		TaxDetail:    req.TaxDetail,
		Level:        req.Level,
		Status:       req.Status,
	}

	fmt.Println(insertModel)

	member, _ := controller.memberRepository.Insert(insertModel)

	if err != nil {
		utils.LoggerError(controllerName, "validate", req, nil, err, false)
		utils.RespondWithJSON(c, http.StatusBadRequest, constants.InvalidRequest, err.Error(), nil)
		return
	}

	rtn := responses.Member{
		Id:           member.ID,
		Name:         member.Name,
		Lastname:     member.Lastname,
		MobileNumber: member.MobileNumber,
		Email:        member.Email,
		Address:      member.Address,
		TaxDetail:    member.TaxDetail,
		Level:        member.Level,
		Status:       member.Status,
		RegisteredAt: member.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	utils.LoggerInfo(controllerName, "success", nil, rtn, false)
	utils.RespondWithJSON(c, http.StatusOK, constants.Success, constants.Success, rtn)
}

func (controller *memberController) List(c *gin.Context) {
	var controllerName = "GetList"
	utils.LoggerInfo(controllerName, "start", nil, nil, false)

	members, _ := controller.memberRepository.List()

	res := make([]responses.Member, 0)
	if len(members) > 0 {
		for _, member := range members {
			member := responses.Member{
				Id:           member.ID,
				Name:         member.Name,
				Lastname:     member.Lastname,
				MobileNumber: member.MobileNumber,
				Email:        member.Email,
				Address:      member.Address,
				TaxDetail:    member.TaxDetail,
				Level:        member.Level,
				Status:       member.Status,
				RegisteredAt: member.CreatedAt.Format("2006-01-02 15:04:05"),
			}
			res = append(res, member)
		}
	}

	utils.LoggerInfo(controllerName, "success", nil, res, false)
	utils.RespondWithJSON(c, http.StatusOK, constants.Success, constants.Success, res)
}

func (controller *memberController) ActiveList(c *gin.Context) {

	var controllerName = "GetList"
	utils.LoggerInfo(controllerName, "start", nil, nil, false)

	members, _ := controller.memberRepository.ActiveList()

	res := make([]responses.Member, 0)
	if len(members) > 0 {
		for _, member := range members {
			member := responses.Member{
				Id:           member.ID,
				Name:         member.Name,
				Lastname:     member.Lastname,
				MobileNumber: member.MobileNumber,
				Email:        member.Email,
				Address:      member.Address,
				TaxDetail:    member.TaxDetail,
				Level:        member.Level,
				Status:       member.Status,
				RegisteredAt: member.CreatedAt.Format("2006-01-02 15:04:05"),
			}
			res = append(res, member)
		}
	}

	utils.LoggerInfo(controllerName, "success", nil, res, false)
	utils.RespondWithJSON(c, http.StatusOK, constants.Success, constants.Success, res)
}

func (controller *memberController) Get(c *gin.Context) {
	var controllerName = "GetList"
	utils.LoggerInfo(controllerName, "start", nil, nil, false)

	id, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)

	member, getError := controller.memberRepository.Get(uint(id))

	if getError != nil {
		utils.LoggerError(controllerName, getError.Error(), id, nil, getError, false)
		utils.RespondWithJSON(c, http.StatusBadRequest, constants.InvalidRequest, getError.Error(), nil)
		return
	}

	rtn := responses.Member{
		Id:           member.ID,
		Name:         member.Name,
		Lastname:     member.Lastname,
		MobileNumber: member.MobileNumber,
		Email:        member.Email,
		Address:      member.Address,
		TaxDetail:    member.TaxDetail,
		Level:        member.Level,
		Status:       member.Status,
		RegisteredAt: member.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	utils.LoggerInfo(controllerName, "success", nil, rtn, false)
	utils.RespondWithJSON(c, http.StatusOK, constants.Success, constants.Success, rtn)
}

func (controller *memberController) Update(c *gin.Context) {
	var controllerName = "Update"
	utils.LoggerInfo(controllerName, "start", nil, nil, false)
	var req requests.Member
	c.Bind(&req)

	if err := validator.New().Struct(req); err != nil {
		utils.LoggerError(controllerName, "validate", req, nil, err, false)
		utils.RespondWithJSON(c, http.StatusBadRequest, constants.InvalidRequest, err.Error(), nil)
		return
	}

	id, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)

	updateModel := models.Member{
		Name:         req.Name,
		Lastname:     req.Lastname,
		MobileNumber: req.MobileNumber,
		Email:        req.Email,
		Address:      req.Address,
		TaxDetail:    req.TaxDetail,
		Level:        req.Level,
		Status:       req.Status,
	}

	err := controller.memberRepository.Update(uint(id), updateModel)
	if err != nil {
		utils.LoggerError(controllerName, err.Error(), id, nil, err, false)
		utils.RespondWithJSON(c, http.StatusBadRequest, constants.InvalidRequest, err.Error(), nil)
		return
	}

	rtn := responses.Member{
		Id:           uint(id),
		Name:         updateModel.Name,
		Lastname:     updateModel.Lastname,
		MobileNumber: updateModel.MobileNumber,
		Email:        updateModel.Email,
		Address:      updateModel.Address,
		TaxDetail:    updateModel.TaxDetail,
		Level:        updateModel.Level,
		Status:       updateModel.Status,
	}

	utils.LoggerInfo(controllerName, "success", rtn, nil, false)
	utils.RespondWithJSON(c, http.StatusOK, constants.Success, constants.Success, rtn)
}

func (controller *memberController) Delete(c *gin.Context) {
	var controllerName = "Delete"
	utils.LoggerInfo(controllerName, "start", nil, nil, false)

	id, _ := strconv.ParseInt(c.Params.ByName("id"), 0, 64)

	err := controller.memberRepository.Delete(uint(id))
	if err != nil {
		utils.LoggerError(controllerName, err.Error(), id, nil, err, false)
		utils.RespondWithJSON(c, http.StatusBadRequest, constants.InvalidRequest, err.Error(), nil)
		return
	}

	utils.LoggerInfo(controllerName, "success", id, nil, false)
	utils.RespondWithJSON(c, http.StatusOK, constants.Success, constants.Success, id)
}
