import "github.com/robfig/cron/v3"

func main() {
    // ... code trước
    c := cron.New()
    c.AddFunc("0 12 * * *", generateReport)
    c.Start()
    router.Run(":3000")
}

func generateReport() {
    // Logic tạo báo cáo, ví dụ:
    reportData := map[string]interface{}{
        "totalSales": 1000.0,
        "numberOfInvoices": 10,
    }
    // Xuất bản qua RabbitMQ
    conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %s", err)
    }
    defer conn.Close()
    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %s", err)
    }
    defer ch.Close()
    q, err := ch.QueueDeclare("daily_report_queue", false, false, false, false, nil)
    if err != nil {
        log.Fatalf("Failed to declare a queue: %s", err)
    }
    jsonData, err := json.Marshal(reportData)
    if err != nil {
        log.Fatalf("Failed to encode report data: %s", err)
    }
    err = ch.Publish("", q.Name, false, false, amqp.Publishing{
        ContentType: "application/json",
        Body:        jsonData,
    })
    if err != nil {
        log.Fatalf("Failed to publish to queue: %s", err)
    }
    log.Println("Report published to queue")
}
