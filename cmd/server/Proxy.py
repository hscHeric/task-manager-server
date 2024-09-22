from UDPClient import *
import json
import base64
from Message import *

class Proxy:
    def __init__(self, hostname, port,):
        self.client = UDPClient(hostname, port)
   
    def sendMessage(self, msg):
        retries = 0
        while retries < self.client.max_retries:
            self.client.send_request(msg)  # Envia a requisição

            response_bytes = self.client.receive_response()  # Espera pela resposta

            if response_bytes is not None:
                return response_bytes  

            print(f"Timeout! Tentativa {retries + 1} de {self.client.max_retries}. Reenviando a mensagem...")
            retries += 1 

        print("Erro: Número máximo de tentativas alcançado. O servidor pode ter ignorado a mensagem.")
        return None  

    
    def do_operation(self, obj_reference, method_id, params):
        id = IDGenerator()
        msg = Message(obj_reference, method_id, params, 0, id)
        message = msg.to_bytes()
        
        return self.sendMessage(message)
    
    def LostMsgTest(self, obj_reference, method_id, params):
        id = IDGenerator()
        msg = Message(obj_reference, method_id, params, 3, id)
        message = msg.to_bytes()
        
        return self.sendMessage(message)
    

    
   
    def InsertTask(self, task):
        return self.do_operation("Task","InsertTask",task)   

    def LostTaskInsetMsg(self, task):
        return self.LostMsgTest("Task","InsertTask",task)      

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

