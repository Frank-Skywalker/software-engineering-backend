package cases

import (
	"github.com/AsterNighT/software-engineering-backend/api"
	"github.com/labstack/echo/v4"
)

type CaseHandler struct {
}

type MedicineHandler struct {
}

// @Summary Get all cases
// @Description
// @Tags Case
// @Produce json
// @Param department query string false "department name" nil
// @Param patient query string false "patient name" nil
// @Param doctor query string false "doctor name" nil
// @Param before query uint false "a timestamp marking end time" nil
// @Param after query uint false "a timestamp marking start time" nil
// @Param q query string false "full-text search" nil
// @Success 200 {object} api.ReturnedData{data=[]Case}
// @Router /cases [GET]
func (h *CaseHandler) GetAllCases(c echo.Context) error {
	// ...
	c.Logger().Debug("GetAllCases")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary Get the last case
// @Description
// @Tags Case
// @Produce json
// @Param patientID path uint true "patient ID"
// @Success 200 {object} api.ReturnedData{data=Case}
// @Router /patient/{patientID}/case [GET]
func (h *CaseHandler) GetLastCaseByPatientID(c echo.Context) error {
	// ...
	c.Logger().Debug("GetLastCase")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary Get cases by patient ID
// @Description
// @Tags Case
// @Produce json
// @Param patientID path uint true "patient ID"
// @Param department query string false "department name" nil
// @Param doctor query string false "doctor name" nil
// @Param before query uint false "a timestamp marking end time" nil
// @Param after query uint false "a timestamp marking start time" nil
// @Param q query string false "full-text search" nil
// @Success 200 {object} api.ReturnedData{data=[]Case}
// @Router /patient/{patientID}/cases [GET]
func (h *CaseHandler) GetCasesByPatientID(c echo.Context) error {
	// ...
	c.Logger().Debug("GetCasesByPatientID")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary New a case
// @Description
// @Tags Case
// @Produce json
// @Param caseDetail body Case true "patient ID, doctor ID, department name and other case details"
// @Success 200 {object} api.ReturnedData{}
// @Router /patient/{patientID}/case [POST]
func (h *CaseHandler) NewCase(c echo.Context) error {
	// ...
	c.Logger().Debug("NewCase")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary Delete a case
// @Description
// @Tags Case
// @Produce json
// @Param caseID path uint true "case ID"
// @Success 200 {object} api.ReturnedData{}
// @Router /patient/{patientID}/case/{caseID} [DELETE]
func (h *CaseHandler) DeleteCaseByCaseID(c echo.Context) error {
	// ...
	c.Logger().Debug("DeleteCaseByCaseID")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary Get previous cases
// @Description
// @Tags Case
// @Produce json
// @Param caseID path uint true "current case ID"
// @Success 200 {object} api.ReturnedData{data=[]Case}
// @Router /patient/{patientID}/case/{caseID} [GET]
func (h *CaseHandler) GetPreviousCases(c echo.Context) error {
	// ...
	c.Logger().Debug("GetPreviousCases")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary Update case
// @Description
// @Tags Case
// @Produce json
// @Param caseDetail body Case true "case ID and updated details"
// @Success 200 {object} api.ReturnedData{}
// @Router /patient/{patientID}/case/{caseID} [PUT]
func (h *CaseHandler) UpdateCase(c echo.Context) error {
	// ...
	c.Logger().Debug("UpdateCase")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary New a prescrition
// @Description
// @Tags Case
// @Produce json
// @Param prescriptionDetail body Prescription true "case ID and prescription details"
// @Success 200 {object} api.ReturnedData{}
// @Router /patient/{patientID}/case/{caseID}/prescription [POST]
func (h *CaseHandler) NewPrescription(c echo.Context) error {
	// ...
	c.Logger().Debug("NewPrescription")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary Delete a prescrition
// @Description
// @Tags Case
// @Produce json
// @Param prescriptionID path uint true "prescription ID"
// @Success 200 {object} api.ReturnedData{}
// @Router /patient/{patientID}/case/{caseID}/prescription/{prescriptionID} [DELETE]
func (h *CaseHandler) DeletePrescription(c echo.Context) error {
	// ...
	c.Logger().Debug("DeletePrescription")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary Update a prescrition
// @Description
// @Tags Case
// @Produce json
// @Param prescriptionDetails body Prescription true "prescription ID and updated details"
// @Success 200 {object} api.ReturnedData{}
// @Router /patient/{patientID}/case/{caseID}/prescription/{prescriptionID} [PUT]
func (h *CaseHandler) UpdatePrescription(c echo.Context) error {
	// ...
	c.Logger().Debug("UpdatePrescription")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary Get prescription by prescription id
// @Description
// @Tags Case
// @Produce json
// @Param prescriptionID path uint true "prescription ID"
// @Success 200 {object} api.ReturnedData{data=Prescription}
// @Router /patient/{patientID}/case/{caseID}/prescription/{prescriptionID}  [GET]
func (h *CaseHandler) GetPrescriptionByPrescriptionID(c echo.Context) error {
	// ...
	c.Logger().Debug("GetPrescriptionByPrescriptionID")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary Get prescriptions by case id
// @Description
// @Tags Case
// @Produce json
// @Param caseID path uint true "case ID"
// @Success 200 {object} api.ReturnedData{data=[]Prescription}
// @Router /patient/{patientID}/case/{caseID}/prescription  [GET]
func (h *CaseHandler) GetPrescriptionByCaseID(c echo.Context) error {
	// ...
	c.Logger().Debug("GetPrescriptionByCaseID")
	return c.JSON(200, api.Return("ok", nil))
}

// @Summary Query medicine by key word
// @Description
// @Tags Medicine
// @Produce json
// @Param q query string true "full-text search"
// @Success 200 {object} api.ReturnedData{data=[]Medicine}
// @Router /medicine [GET]
func (h *MedicineHandler) GetMedicines(c echo.Context) error {
	// ...
	c.Logger().Debug("QueryMedicine")
	return c.JSON(200, api.Return("ok", nil))
}
