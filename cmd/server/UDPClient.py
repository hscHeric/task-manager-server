import socket
from datetime import datetime, timezone
from Message import *
import time

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

