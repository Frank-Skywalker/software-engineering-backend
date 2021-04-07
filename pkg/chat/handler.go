package chat

import (
	"github.com/AsterNighT/software-engineering-backend/api"
	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
}

// @Summary New a chat
// @Description
// @Tags Chat
// @Produce json
// @Param chatDetail body Chat true "patient ID, doctor ID and other chat details"
// @Success 200 {object} api.ReturnedData{data=Chat}
// @Router /patient/{patientID}/chat [POST]
func (h *ChatHandler) NewChat(c echo.Context) error {
	// ...
	c.Logger().Debug("hello world")
	return c.JSON(200, api.Return("NewChat", nil))
}

// @Summary Delete a chat
// @Description
// @Tags Chat
// @Produce json
// @Param chatID path uint true "chat ID"
// @Success 200 {object} api.ReturnedData{}
// @Router /patient/{patientID}/chat/{chatID} [DELETE]
func (h *ChatHandler) DeleteChatByChatID(c echo.Context) error {
	// ...
	c.Logger().Debug("hello world")
	return c.JSON(200, api.Return("DeleteChatByChatID", nil))
}

// @Summary Get a message by message ID
// @Description
// @Tags Chat
// @Produce json
// @Param messageID path uint true "message ID"
// @Success 200 {array} api.ReturnedData{data=Message}
// @Router /patient/{patientID}/chat/{chatID}/messsage/{messageID} [GET]
func (h *ChatHandler) GetMessageByMessageID(c echo.Context) error {
	// ...
	c.Logger().Debug("hello world")
	return c.JSON(200, api.Return("GetMessageByMessageID", nil))
}

// @Summary Get the lastest message
// @Description
// @Tags Chat
// @Produce json
// @Param ID path uint true "chat ID"
// @Success 200 {object} api.ReturnedData{data=Message}
// @Router /patient/{patientID}/chat/{chatID} [GET]
func (h *ChatHandler) GetLastMessage(c echo.Context) error {
	// ...
	c.Logger().Debug("hello world")
	return c.JSON(200, api.Return("GetLastMessage", nil))
}

// @Summary New a Message
// @Description
// @Tags Chat
// @Produce json
// @Param messageDetail body Message true "chat ID, type and other message details"
// @Success 200 {object} api.ReturnedData{data=Message}
// @Router /patient/{patientID}/chat/{chatID}/message [POST]
func (h *ChatHandler) NewMessage(c echo.Context) error {
	// ...
	c.Logger().Debug("hello world")
	return c.JSON(200, api.Return("NewMessage", nil))
}

// @Summary Delete a message
// @Description Can be viewed as recall a message
// @Tags Chat
// @Produce json
// @Param messageID path uint true "message ID"
// @Success 200 {object} api.ReturnedData{}
// @Router /patient/{patientID}/chat/{chatID}/message/{messageID} [DELETE]
func (h *ChatHandler) DeleteMessage(c echo.Context) error {
	// ...
	c.Logger().Debug("hello world")
	return c.JSON(200, api.Return("DeleteMessage", nil))
}

// @Summary Get a message ID list by chat ID
// @Description
// @Tags Chat
// @Produce json
// @Param chatID path uint true "chat ID"
// @Success 200 {array} api.ReturnedData{data=[]uint}
// @Router /patient/{patientID}/chat/{chatID} [GET]
func (h *ChatHandler) GetMessagesByChatID(c echo.Context) error {
	// ...
	c.Logger().Debug("hello world")
	return c.JSON(200, api.Return("GetMessagesByChatID", nil))
}

// @Summary Get catagorys by keyword id
// @Description
// @Tags Chat
// @Produce json
// @Param keywordID path uint true "keyword ID"
// @Success 200 {object} api.ReturnedData{data=[]uint}
// @Router /keyword/{keywordID}  [GET]
func (h *ChatHandler) GetCatagorysByKeywordID(c echo.Context) error {
	// ...
	c.Logger().Debug("hello world")
	return c.JSON(200, api.Return("GetCatagorysByKeywordID", nil))
}

// @Summary Get questions by catagory id
// @Description
// @Tags Chat
// @Produce json
// @Param catagoryID path uint true "catagory ID"
// @Success 200 {object} api.ReturnedData{data=[]string}
// @Router /keyword/{keywordID}/catagory/{catagoryID}  [GET]
func (h *ChatHandler) GetQuestiosByCatagoryID(c echo.Context) error {
	// ...
	c.Logger().Debug("hello world")
	return c.JSON(200, api.Return("GetQuestionsByCatagoryID", nil))
}
