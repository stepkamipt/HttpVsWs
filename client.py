import time
import requests
import websocket
import json
import socketio

# Server address
SERVER = 'localhost:200' # localhost, port 200

def calculate_squares_http(requests_count):
    start_time = time.time()

    session = requests.Session()
    request_url = f"http://{SERVER}/square"
    for number in range(requests_count):
        response = session.get(f"{request_url}?number={number}")

    end_time = time.time()
    total_time = end_time - start_time
    print(f"HTTP Total time: {total_time:.2f} seconds for {requests_count} requests")

def calculate_squares_ws(requests_count):
    start_time = time.time()
    responses_count = 0  

    sio = socketio.Client()

    @sio.event
    def connect():
        for value in range(requests_count):
            data = {"number": value, "id": value}
            sio.send(json.dumps(data))

    @sio.event
    def message(data):
        nonlocal responses_count 
        nonlocal start_time

        responses_count = responses_count + 1
        if responses_count < requests_count:
            return
        
        end_time = time.time()
        total_time = end_time - start_time
        print(f"WS Total time: {total_time:.2f} seconds for {requests_count} requests")
        sio.disconnect()

    @sio.event
    def disconnect():
        print("WebSocket connection closed.")

    sio.connect(f"ws://{SERVER}")
    sio.wait()

if __name__ == "__main__":
    calculate_squares_http( 10000 )
    calculate_squares_ws( 10000 )



