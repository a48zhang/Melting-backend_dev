import time
import threading
import websocket
import json

url = "ws://localhost:24769/ws"
data = {"service": "echo", "message": "is everything ok?"}


def on_open(wsapp):
    print("on_open")

    def send_message():
        wsapp.send(json.dumps(data))
        return

    threading.Thread(target=send_message).start()


def on_message(wsapp, message):
    print("on_message:", message)


def on_close(wsapp):
    print("on_close")


def on_data(wsapp, frame_data, frame_opcode, frame_fin):
    print("on_data:", frame_data)


wsapp = websocket.WebSocketApp(url,
                               on_open=on_open,
                               on_message=on_message,
                               on_data=on_data)
wsapp.run_forever()
