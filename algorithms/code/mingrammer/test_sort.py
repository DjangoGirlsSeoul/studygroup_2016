import unittest

from heapsort import heapsort
from mergesort import mergesort


class TestSort(unittest.TestCase):
    def test_heapsort(self):
        arr = [9, 3, 2, 5, 8, 4, 7, 10, 6, 1]

        arr = heapsort(arr)

        self.assertEqual(arr, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10])

    def test_mergesort(self):
        arr = [10, 1, 4, 3, 8, 2, 7, 9, 5, 6]

        arr = mergesort(arr)

        self.assertEqual(arr, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10])


if __name__ == '__main__':
    unittest.main()
