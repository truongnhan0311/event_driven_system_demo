from kafka import KafkaConsumer

consumer = KafkaConsumer('order_details', bootstrap_servers=["localhost:9092", "localhost:9093" , "localhost:9094"])

print("Gonna start listening")
while True:
    for message in consumer:
        print("Here is a message..")
        print (message)