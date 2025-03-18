package handlers

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/user/go-invoice-service/db"
    "github.com/user/go-invoice-service/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func CreateInvoice(c *gin.Context) {
    var invoice models.Invoice
    if err := c.ShouldBindJSON(&invoice); err != nil {
        c.AbortWithStatus(400)
        return
    }
    collection := db.GetDatabase().Collection("invoices")
    result, err := collection.InsertOne(context.TODO(), invoice)
    if err != nil {
        c.AbortWithStatus(500)
        return
    }
    invoice.ID = result.InsertedID.(string)
    c.JSON(201, invoice)
}
