# passwdmgr/passwdmgr_test.py
import unittest
from .passwdmgr import add, get, delete, list_services

class Test(unittest.TestCase):

    def test_basic_operations(self):
        add("web", "prueba", "prueba_pass")
        add("email", "prueba2", "prueba2_pass")
        creds = get("web")
        self.assertEqual(creds["username"], "prueba")
        self.assertEqual(creds["password"], "prueba_pass")
        services = list_services()
        self.assertIn("web", services)
        self.assertIn("email", services)
        delete("web")
        services_after = list_services()
        self.assertNotIn("web", services_after)

    def test_add_and_get(self):
        add("single", "user1", "pass1")
        creds = get("single")
        self.assertEqual(creds["username"], "user1")
        self.assertEqual(creds["password"], "pass1")

    def test_fail_intentional(self):
        add("fail", "user_fail", "pass_fail")
        creds = get("fail")
        self.assertEqual(creds["username"], "incorrect_user")  # test intencionalmente falla

if __name__ == "__main__":
    unittest.main()