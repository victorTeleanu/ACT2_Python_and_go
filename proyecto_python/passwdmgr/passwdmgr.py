# passwdmgr/passwdmgr.py

class Entry:
    """Entrada de servicio, usuario y contraseña."""
    def __init__(self, service: str, username: str, password: str):
        self.service = service
        self.username = username
        self.password = password

    def to_dict(self):
        return {"service": self.service, "username": self.username, "password": self.password}

# Diccionario para guardar contraseñas
passwords = {}

# Añadir o actualizar un servicio
def add(service, username, password):
    passwords[service] = {"username": username, "password": password}

# Obtener credenciales
def get(service):
    return passwords.get(service, "Servicio no encontrado")

# Eliminar un servicio
def delete(service):
    if service in passwords:
        del passwords[service]
    else:
        print("Servicio no encontrado")

# Listar todos los servicios
def list_services():
    return list(passwords.keys())