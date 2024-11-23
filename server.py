from flask import Flask, request, jsonify
import logging
from flask_socketio import SocketIO, send
import json

app = Flask(__name__)
socketio = SocketIO(app)

log = logging.getLogger('werkzeug')
log.setLevel(logging.ERROR)

@app.route('/square', methods=['GET'])
def calculate_square():
    number = request.args.get('number')
    result = int(number) ** 2
    return jsonify({'squared_value': number}), 200

@socketio.on('message')
def handle_message(data):
    json_data = json.loads(data)
    number = json_data.get('number')
    id = json_data.get('id')
    result = number ** 2
    send(json.dumps({'id': id, 'squared_value': number}))

if __name__ == '__main__':
    socketio.run(app, host='0.0.0.0', port=200)






