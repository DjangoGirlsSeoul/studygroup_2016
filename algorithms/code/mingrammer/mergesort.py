def mergesort(arr):
    size = len(arr)

    if size == 1:
        return arr

    l_arr = mergesort(arr[:int(size/2)])
    r_arr = mergesort(arr[int(size/2):])

    l_size = len(l_arr)
    r_size = len(r_arr)

    i, j = 0, 0

    merged = []

    for k in range(size):
        if i == l_size:
            merged.append(r_arr[j])
            j += 1
        elif j == r_size:
            merged.append(l_arr[i])
            i += 1
        else:
            if l_arr[i] <= r_arr[j]:
                merged.append(l_arr[i])
                i += 1
            else:
                merged.append(r_arr[j])
                j += 1

    return merged
