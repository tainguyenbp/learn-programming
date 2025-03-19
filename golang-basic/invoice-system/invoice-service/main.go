package main

import (
    "github.com/gin-gonic/gin"
    "github.com/user/go-invoice-service/db"
    "github.com/user/go-invoice-service/handlers"
)

func main() {
    err := db.Connect()
    if err != nil {
        panic(err)
    }
    defer db.Close()
    router := gin.Default()
    router.Post("/invoices", handlers.CreateInvoice)
    router.Get("/invoices", handlers.GetInvoices)
    router.Get("/invoices/:id", handlers.GetInvoiceById)
    router.Get("/invoices/range", handlers.GetInvoicesByRange)
    router.Run(":3000")
}
