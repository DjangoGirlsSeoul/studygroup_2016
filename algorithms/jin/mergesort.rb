def mergesort(array)
    return array if array.length < 2

    mid = array.length / 2
    sorted_left = mergesort(array[0...mid])
    sorted_right = mergesort(array[mid..array.length])
    
    merge(sorted_left, sorted_right)
end

def merge(left, right)
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
