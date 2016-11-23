var quicksort = function(arr) {
    
    var sorted_array = arr.slice()
    
    function swap(target, source) {
        var tmp = sorted_array[target];
        sorted_array[target] = sorted_array[source];
        sorted_array[source] = tmp;
    }
    
    function partition(start, end){
        var pivot = sorted_array[start];
        var tmpIdx = start+1;
        for(var i = start+1; i <= end; i++){
            if(sorted_array[i] < pivot){
                swap(i, tmpIdx);
                tmpIdx++;
            }
        }
        swap(start, tmpIdx-1);
        return tmpIdx-1;
    }
    
    function _quicksort(start, end) {
        if(end > start){
            var pivotIndex = partition(start, end);
            _quicksort(start, pivotIndex -1)
            _quicksort(pivotIndex+1, end)
        }
    }
    
    _quicksort(0, sorted_array.length -1);
    return sorted_array; 
}

function processData(input) {
    //Enter your code here
    var input_arr = input.split('\n');
    var len = input_arr[0];
    var arr = input_arr[1].split(' ').map(function(i) { return parseInt(i) });
    console.log(quicksort(arr))
    
} 

process.stdin.resume();
process.stdin.setEncoding("ascii");
_input = "";
process.stdin.on("data", function (input) {
    _input += input;
});

process.stdin.on("end", function () {
   processData(_input);
});

