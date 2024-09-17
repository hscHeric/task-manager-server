from UDPClient import *
import json
import base64
from Message import *

class Proxy:
    def __init__(self, hostname, port):
        self.client = UDPClient(hostname, port)
   
    def sendMessage(self,msg):
        self.client.send_request(msg)
        
        response_bytes = self.client.receive_response()
        return response_bytes
    
    def do_operation(self, obj_reference, method_id, params):
        id = IDGenerator()
        msg = Message(obj_reference, method_id, params, 0, id)
        message = msg.to_bytes()
        
        return self.sendMessage(message)
    
   
    def InsertTask(self, task):
        return self.do_operation("Task","InsertTask",task)     

    def GetTaskByID(self, task_id):
        return self.do_operation("Task","GetTaskByID", task_id)

    def DeleteTask(self, task_id):
        return self.do_operation("Task","DeleteTask", task_id)

    def GetAllTasks(self):
        id = IDGenerator()
        request_bytes = b'' 
        message = Message("Task", "GetAllTasks", request_bytes, 0, id)
        message_bytes = message.to_bytes()
        return self.sendMessage(message_bytes)
        
        
        

    def close(self):
        self.client.close()

