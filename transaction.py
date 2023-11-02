import json

from kafka import KafkaConsumer
from kafka import KafkaProducer


ORDER_KAFKA_TOPIC = "order_details"
ORDER_CONFIRMED_KAFKA_TOPIC = "order_confirmed"

consumer = KafkaConsumer(
    ORDER_KAFKA_TOPIC,
    bootstrap_servers=["localhost:9092", "localhost:9093" , "localhost:9094"],
    value_deserializer=lambda v: json.dumps(v).encode("utf-8"),
    auto_offset_reset='earliest'
)
producer = KafkaProducer(bootstrap_servers=["localhost:9092", "localhost:9093" , "localhost:9094"])


print("Gonna start listening")
while True:
    records = consumer.poll(60 * 1000)  # timeout in millis , here set to 1 min

    for message in records.items():

        print("Ongoing transaction..")
        consumed_message = message
        print(consumed_message)
        user_id = consumed_message["user_id"]
        total_cost = consumed_message["total_cost"]
        data = {
            "customer_id": user_id,
            "customer_email": f"{user_id}@gmail.com",
            "total_cost": total_cost
        }
        print("Successful transaction..")
        producer.send(ORDER_CONFIRMED_KAFKA_TOPIC, json.dumps(data).encode("utf-8"))