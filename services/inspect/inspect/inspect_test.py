import unittest
import inspect as ins

class InspectTest(unittest.TestCase):
    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_can_fetch(self):
        self.assertEqual(ins.inspect("http://www.musi-cal.com/"), True)
    
    def test_cannot_fetch(self):
        self.assertEqual(ins.inspect("http://www.musi-cal.com/cgi-bin/search?city=San+Francisco"), False)

if __name__ == "__main__":
    unittest.main()
