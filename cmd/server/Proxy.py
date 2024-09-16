from UDPClient import *
import json
import base64
from Message import *

class Proxy:
    def __init__(self, hostname, port):
        self.client = UDPClient(hostname, port)
   
    def InsertTask(self, task):
        p = ('localhost', 12345)
        id = IDGenerator()
        obj_reference = "Task"
        method_id = "InsertTask"      
        msg = Message(obj_reference,method_id,task,0,id)
        message = msg.to_bytes()
        self.client.send_request(message)
        request = self.client.receive_response()        
        return request

    def GetTaskById(self, task_id, task):
        # Edita uma tarefa
        response = self.do_operation("Task","GetTaskById", task_id=task_id, task=task)
        return response

    def RemoveTask(self, task_id):
        # Remove uma tarefa
        response = self.do_operation("Task","RemoveTask", task_id=task_id)
        return response

    def GetAllTasks(self):
        # Lista todas as tarefas
        response = self.do_operation("Task","GetAllTasks")
        return response.get('tarefas', [])  # Retorna uma lista de tarefas

    def close(self):
        self.client.close()
