package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HistoryController struct {
	DB *gorm.DB
}

func NewHistoryController(DB *gorm.DB) HistoryController {
	return HistoryController{DB}
}

func (hc *HistoryController) GetAllHistoryRecords(ctx *gin.Context) {}

func (hc *HistoryController) GetAllSubscriptionRecords(ctx *gin.Context) {}

func (hc *HistoryController) GetAllRentalRecords(ctx *gin.Context) {}

func (hc *HistoryController) GetRentalRecordByID(ctx *gin.Context) {}
func (hc *HistoryController) GetSubscriptionByID(ctx *gin.Context) {}
