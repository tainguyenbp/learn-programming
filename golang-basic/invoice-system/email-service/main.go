package main

import (
    "context"
    "log"
    "os"
    "time"
    "github.com/streadway/amqp"
)

func main() {
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
    msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
    if err != nil {
        log.Fatalf("Failed to consume from queue: %s", err)
    }
    go func() {
        for msg := range msgs {
            processReport(msg.Body)
        }
    }()
    select {}
}

func processReport(body []byte) {
    var report map[string]interface{}
    err := json.Unmarshal(body, &report)
    if err != nil {
        log.Fatalf("Failed to parse report data: %s", err)
    }
    sendEmail(report)
}

func sendEmail(report map[string]interface{}) {
    smtpHost := os.Getenv("SMTP_HOST")
    smtpPort := os.Getenv("SMTP_PORT")
    smtpUser := os.Getenv("SMTP_USER")
    smtpPass := os.Getenv("SMTP_PASS")
    to := "recipient@example.com"
    subject := "Daily Sales Report"
    body := fmt.Sprint(report)
    e := email.NewEmail()
    e.From = smtpUser
    e.To = []string{to}
    e.Subject = subject
    e.HTML = []byte("<h1>Daily Sales Report</h1>" + body)
    err := e.Send(smtpHost+":"+smtpPort, smtpUser, smtpPass)
    if err != nil {
        log.Fatalf("Failed to send email: %s", err)
    }
    log.Println("Email sent successfully")
}
