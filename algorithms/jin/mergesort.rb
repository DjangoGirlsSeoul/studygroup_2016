def mergesort(array)
    # base case
    return array if array.length < 2

    mid = array.length / 2
    # three dots(...) doesnt take the last value
    sorted_left = mergesort(array[0...mid])
    # two dots(..) does take the last value
    sorted_right = mergesort(array[mid..array.length])
    
    merge(sorted_left, sorted_right)
end

def merge(left, right)
    # ruby is magic! (empty? is same as array.length == 0)
    if left.empty?
        right
    elsif right.empty?
        left
    elsif left.first < right.first
        [left.first] + merge(left[1..left.length], right)
    else
        [right.first] + merge(left,right[1..right.length])
    end
end
