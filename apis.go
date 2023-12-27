package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetMemos(c *gin.Context) {
	var memos []Memo
	db.Find(&memos)
	c.JSON(http.StatusOK, memos)
}

func GetMemo(c *gin.Context) {
	id := c.Param("id")
	var memo Memo
	if err := db.First(&memo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Memo not found"})
		return
	}
	c.JSON(http.StatusOK, memo)
}

func CreateMemo(c *gin.Context) {
	var memo Memo
	if err := c.ShouldBindJSON(&memo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&memo)
	c.JSON(http.StatusCreated, memo)
}

func UpdateMemo(c *gin.Context) {
	id := c.Param("id")
	var memo Memo
	if err := db.First(&memo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Memo not found"})
		return
	}

	var updatedMemo Memo
	if err := c.ShouldBindJSON(&updatedMemo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	memo.Title = updatedMemo.Title
	memo.Content = updatedMemo.Content
	db.Save(&memo)

	c.JSON(http.StatusOK, memo)
}

func DeleteMemo(c *gin.Context) {
	id := c.Param("id")
	var memo Memo
	if err := db.First(&memo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Memo not found"})
		return
	}

	db.Delete(&memo)
	c.JSON(http.StatusOK, gin.H{"message": "Memo deleted successfully"})
}
