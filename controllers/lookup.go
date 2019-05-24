package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/vin/core"
)

type LookupController struct {
	control.APIController
}

func NewLookupCtrl(ctrlMap *control.ControllerMap) *LookupController {
	result := &LookupController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title Validate and Deserialize
// @Description Gets the details of a VIN after validation
// @Success 200 {[]core.Profile} []core.Portfolio]
// @router /:vin [get]
func (req *LookupController) Get() {
	vin := req.Ctx.Input.Param(":vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	results, err := core.BuildInfo(vin)

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, results)
}

// @Title Create & Save
// @Description Creates the details of a VIN after validation
// @Success 200 {husk.Key} husk.Key
// @router / [post]
func (req *LookupController) Post() {
	var obj core.VIN
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &obj)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := obj.Create()

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}
