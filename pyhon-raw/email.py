import json

from kafka import KafkaConsumer


ORDER_CONFIRMED_KAFKA_TOPIC = "order_confirmed"


consumer = KafkaConsumer(
    ORDER_CONFIRMED_KAFKA_TOPIC,
    bootstrap_servers=["localhost:9092", "localhost:9093" , "localhost:9094"]
)

emails_sent_so_far = set()
print("Gonna start listening")
while True:
    print("hello")
    for message in consumer:
        print(message)
        consumed_message = json.loads(message.value.decode())
        customer_email = consumed_message["customer_email"]
        print(f"Sending email to {customer_email} ")
        emails_sent_so_far.add(customer_email)
        print(f"So far emails sent to {len(emails_sent_so_far)} unique emails")