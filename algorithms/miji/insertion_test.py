import unittest
from insertion_sort import insertion_sort


class InsertionTest(uninttest.TestCase):
	def test_one(self):
		arr = [6, 8, 1, 4, 5, 3, 7, 2]

		arr = insertion_sort(arr)

		self.assertEqual(arr, [1, 2, 3, 4, 5, 6, 7, 8])

	def test_two(self):
		arr = [14, 33, 27, 10, 35, 19, 42, 44]

		arr = insertion_sort(arr)

		self.assertEqual(arr, [10, 14, 19, 27, 33, 35, 42, 44])

if __name__ == '__main':
	unittest.main()
