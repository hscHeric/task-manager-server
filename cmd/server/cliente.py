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

def imprimir(response):
    print("\nResposta do Servidor:")
    response_json = json.loads(json.dumps(response))
    args_base64 = response_json.get('Args', '')
    args_bytes = base64.b64decode(args_base64)
    request_data = json.loads(args_bytes.decode('utf-8'))
    print("ObjReference:", response_json.get('ObjReference'))
    print("MethodID:", response_json.get('MethodID'))
    print("Args:", request_data)  # Exibe os dados da requisição
    print("T:", response_json.get('T'))
    print("ID:", response_json.get('ID'))
    print("StatusCode:", response_json.get('StatusCode'))
    

def main():
    hostname = 'localhost'
    port = 1234  

    try:
        proxy = Proxy(hostname, port)

        while True:
            print("\nEscolha uma opção:")
            print("1. Adicionar Tarefa")
            print("2. Editar Tarefa")
            print("3. Remover Tarefa")
            print("4. Listar Tarefas")
            print("5. Sair")
            opcao = input("Opção: ")

            if opcao == '1':
                titulo = input("Título: ")
                descricao = input("Descrição: ")
                data_vencimento = input("Data de Vencimento (YYYY-MM-DD): ")

                # Verifica se a data é válida
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

                # Verifica se o ID é válido
                try:
                    validar_id(task_id)
                except ValueError as e:
                    print(e)
                    continue

                # Verifica se o ID existe
                if not id_existe(task_id, tarefas):
                    print(f"Erro: Tarefa com ID {task_id} não encontrada.")
                    continue

                titulo = input("Novo Título: ")
                descricao = input("Nova Descrição: ")
                data_vencimento = input("Nova Data de Vencimento (YYYY-MM-DD): ")

                try:
                    validar_data(data_vencimento)
                except ValueError as e:
                    print(e)
                    continue

                task = {'titulo': titulo, 'descricao': descricao, 'data_vencimento': data_vencimento}
                response = proxy.GetTaskById(task_id, task)
                print(f"Tarefa {task_id} editada com sucesso.")

            elif opcao == '3':
                task_id = input("ID da Tarefa: ")

                # Verifica se o ID é válido
                try:
                    validar_id(task_id)
                except ValueError as e:
                    print(e)
                    continue

                # Verifica se o ID existe
                if not id_existe(task_id, tarefas):
                    print(f"Erro: Tarefa com ID {task_id} não encontrada.")
                    continue

                response = proxy.RemoveTask(task_id)
                print(f"Tarefa {task_id} removida com sucesso.")

            elif opcao == '4':
                tarefas = proxy.GetAllTasks()
                print("Tarefas cadastradas:")
                for tarefa in tarefas:
                    print(f"ID: {tarefa['id']}, Título: {tarefa['titulo']}, Descrição: {tarefa['descricao']}, Data de Vencimento: {tarefa['data_vencimento']}")

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



