import  json
import base64
import threading
import sys


class IDGenerator:
    def __init__(self):
        self.id = sys.maxsize  
        self.lock = threading.Lock()

    def get_next_id(self):
        with self.lock:
            self.id -= 1
            return self.id
        

class Message:
    def __init__(self, obj_reference, method_id, request_bytes, t, id_generator):
        self.ObjReference = obj_reference
        self.MethodID = method_id
        self.Args = base64.b64encode(request_bytes).decode('utf-8')
        self.T = t
        self.ID = id_generator.get_next_id()  
        self.StatusCode = 0

    def to_bytes(self):
        message = {
            "ObjReference": self.ObjReference,
            "MethodID": self.MethodID,
            "Args": self.Args,
            "T": self.T,
            "ID": self.ID,
            "StatusCode": self.StatusCode
        }
        
        json_message = json.dumps(message)
        return json_message.encode('utf-8')
