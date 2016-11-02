def mergesort(arr):
    size = len(arr)

    if size == 1:
        return arr

    l_arr = mergesort(arr[:size // 2])
    r_arr = mergesort(arr[size // 2:])

    l_size = len(l_arr)
    r_size = len(r_arr)

    i, j = 0, 0

    merged = []

    while i < l_size and j < r_size:
        if l_arr[i] < r_arr[j]:
            merged.append(l_arr[i])
            i += 1
        else:
            merged.append(r_arr[j])
            j += 1

    merged.extend(l_arr[i:])
    merged.extend(r_arr[j:])

    return merged
