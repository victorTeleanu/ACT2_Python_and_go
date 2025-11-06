# main.py
from passwdmgr.passwdmgr import add, get, delete, list_services

def main():
    add("aules", "10480107", "holamundo123")
    add("email", "alexteleanu9@gmail.com", "p@ssw0rd!")

    print("Servicios guardados:", list_services())
    print("Credenciales de aules:", get("aules"))
    print("Credenciales de email:", get("email"))

    delete("email")
    print("Servicios después de eliminar email:", list_services())

if __name__ == "__main__":
    main()
