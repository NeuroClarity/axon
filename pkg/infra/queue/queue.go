package queue

import (
  "fmt"
  "encoding/json"
  "github.com/streadway/amqp"
  "github.com/NeuroClarity/axon/pkg/domain/gateway"
)

// NOTE: Rabbit Queues and Exchanges are created when the RabbitMQ host is created
func NewQueue(username, password, endpoint, port, queueName string) (gateway.Queue, error) {
  targetUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, endpoint, port)
  conn, err := amqp.Dial(targetUrl)
  if err != nil {
    return nil, errors.New("Unable to connect to RabbitMQ servers")
  }
  ch, err := conn.Channel()
  if err != nil {
    return nil, errors.New("Unable to open a connection with RabbitMQ servers")
  }
  ch.QueueDeclare(
    queueName,  // name
    true,       // durable
    false,      // auto delete
    false,      // exclusive 
    true,       // no wait 
    nil         // args
  )
  return &queue{
          channel: ch,
          queue: queueName
        }, nil
}

type queue struct {
  channel     *amqp.Channel
  queue       string
}

func (repo queue) publishData(data string) error {
  err = repo.channel.Publish(
    "",         // default: direct exchange
    repo.queue, // routing key
    false,      // mandatory
    false,      // immediate
    amqp.Publishing {
      ContentType: "text/plain",
      Body:        []byte(data),
    })

  return err
}

