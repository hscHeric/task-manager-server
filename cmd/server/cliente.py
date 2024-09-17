from Proxy import Proxy
from datetime import datetime, timezone
import json
import base64

def validar_data(data_str):
    try:
        data = datetime.strptime(data_str, '%Y-%m-%d')
        
        data = data.replace(hour=23, minute=59, second=59, microsecond=0, tzinfo=timezone.utc)
        iso_format_data = data.isoformat()

        return iso_format_data  
    except ValueError:
        raise ValueError(f"Data inválida: {data_str}. O formato correto é YYYY-MM-DD.")

def validar_id(task_id):
    if not task_id.isdigit():
        raise ValueError(f"ID inválido: {task_id}. O ID deve ser numérico.")

def id_existe(task_id, tarefas):
    return any(tarefa['id'] == int(task_id) for tarefa in tarefas)

def create_task(title, description, date):
    task = {
    "date": date,
    "title": title,
    "description": description
    }
    
    request_json = json.dumps(task)
    request_bytes = request_json.encode('utf-8')

    return request_bytes

def create_task_id(task_id):
    task_id_dict = {
        "TaskId": task_id
    }

    request_json = json.dumps(task_id_dict)  # Use o dicionário task_id_dict
    request_bytes = request_json.encode('utf-8')

    return request_bytes


def imprimirTarefas(response):
    print("Resposta do Servidor:\n")
    
    response_json = json.loads(json.dumps(response))
    
    args_base64 = response_json.get('Args', '')
    args_bytes = base64.b64decode(args_base64)
    
    request_data = json.loads(args_bytes.decode('utf-8'))
    
    tasks = request_data
    print("\nTarefas Recebidas:")
    for task in tasks:
        print(f"\nTask ID: {task['taskId']}")
        print(f"Date: {task['date']}")
        print(f"Title: {task['title']}")
        print(f"Description: {task['description']}")

def imprimir(response):
    response_json = json.loads(json.dumps(response))
    args_base64 = response_json.get('Args', '')
    args_bytes = base64.b64decode(args_base64)
    print(args_bytes)


def main():
    hostname = 'localhost'
    port = 1234  

    try:
        proxy = Proxy(hostname, port)

        while True:
            print("\nEscolha uma opção:")
            print("1. Adicionar Tarefa")
            print("2. Obter Tarefa pelo Id")
            print("3. Remover Tarefa")
            print("4. Listar Tarefas")
            print("5. Sair")
            opcao = input("Opção: ")

            if opcao == '1':
                titulo = input("Título: ")
                descricao = input("Descrição: ")
                data_vencimento = input("Data de Vencimento (YYYY-MM-DD): ")

                try:
                    data = validar_data(data_vencimento)
                except ValueError as e:
                    print(e)
                    continue
                task = create_task(titulo,descricao,data)
                response = proxy.InsertTask(task)
                if response:
                    imprimir(response)
                else:
                    print('Error ao inserir a tarefa')


            elif opcao == '2':
                task_id = input("ID da Tarefa: ")
                taskId = create_task_id(task_id)
                response = proxy.GetTaskByID(taskId)
                if response:
                    imprimir(response)
                else:
                    print("Erro ao obter tarefa")

            elif opcao == '3':
                task_id = input("ID da Tarefa: ")
                taskId = create_task_id(task_id)
                response = proxy.DeleteTask(taskId)
                if response:
                    imprimir(response)
                else:
                    print("Erro ao excluir tarefa")


            elif opcao == '4':
                response = proxy.GetAllTasks()
                if response:
                    imprimirTarefas(response)

            elif opcao == '5':
                break

            else:
                print("Opção inválida. Tente novamente.")

    except Exception as e:
        print(f"Erro: {e}")
    finally:
        proxy.close()

if __name__ == "__main__":
    main()





