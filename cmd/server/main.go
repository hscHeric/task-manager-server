import socket
import json
import base64
from datetime import datetime, timezone

class UDPClient:
    def __init__(self, hostname, port, max_retries=3):
        self.server_address = (hostname, port)
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        self.socket.settimeout(None)  # Remove o timeout
        self.max_retries = max_retries

    def send_request(self, request_bytes):
        # Envia a requisição para o servidor
        self.socket.sendto(request_bytes, self.server_address)

    def receive_response(self):
        retries = 0
        while retries < self.max_retries:
            try:
                packet, _ = self.socket.recvfrom(4096)
                return json.loads(packet.decode())
            except socket.timeout:
                retries += 1
                print(f"Timeout! Tentativa {retries} de {self.max_retries}. Reenviando a mensagem...")

        print("Erro: Número máximo de tentativas alcançado. Falha na comunicação com o servidor.")
        return None

    def close(self):
        self.socket.close()

def create_request(title, description, date=None):
    # Cria a estrutura da requisição, com data opcional.
    if date is None:
        date = datetime.now(timezone.utc).isoformat()  # Formato RFC3339 com fuso horário
    
    request = {
        "date": date,
        "title": title,
        "description": description
    }
    
    # Serializa a requisição para JSON e converte para bytes
    request_json = json.dumps(request)
    request_bytes = request_json.encode('utf-8')

    return request_bytes

def create_message(obj_reference, method_id, request_bytes, t, id_value):
    # Cria a mensagem
    message = {
        "ObjReference": obj_reference,
        "MethodID": method_id,
        "Args": base64.b64encode(request_bytes).decode('utf-8'),  # Codifica bytes em base64 para envio em JSON
        "T": t,
        "ID": id_value,
        "StatusCode": 0
    }
    
    # Serializa a mensagem inteira em JSON
    json_message = json.dumps(message)
    
    return json_message.encode('utf-8')

def main():
    client = UDPClient('localhost', 12345)

    # Cria uma requisição para enviar
    request = create_request("Estudar Python", "Revisar conceitos de sockets e JSON")

    # Cria uma mensagem com a requisição convertida em bytes
    message = create_message("Task", "InsertTask", request, 1, 1001)

    # Envia a mensagem para o servidor
    client.send_request(message)

    # Recebe a resposta do servidor
    response = client.receive_response()
    if response:
        print("Resposta do Servidor:")
        
        # Desserializa a resposta JSON
        response_json = json.loads(json.dumps(response))
        
        # Converte o campo base64 'Args' de volta para bytes
        args_base64 = response_json.get('Args', '')
        args_bytes = base64.b64decode(args_base64)
        
        # Desserializa os bytes para um dicionário Python
        request_data = json.loads(args_bytes.decode('utf-8'))
        
        # Imprime o conteúdo da resposta
        print("ObjReference:", response_json.get('ObjReference'))
        print("MethodID:", response_json.get('MethodID'))
        print("Args:", request_data)  # Exibe os dados da requisição
        print("T:", response_json.get('T'))
        print("ID:", response_json.get('ID'))
        print("StatusCode:", response_json.get('StatusCode'))

    # Fecha o cliente
    client.close()

if __name__ == "__main__":
    main()

