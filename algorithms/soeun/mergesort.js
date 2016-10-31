var merge = function(array, p, q, r) {
    var lowHalf = [];
    var highHalf = [];


    var k = p;
    var i;
    var j;
    for (i = 0; k <= q; i++, k++) {
        lowHalf[i] = array[k];
    }
    for (j = 0; k <= r; j++, k++) {
        highHalf[j] = array[k];
    }


    k = p;
    i = 0;
    j = 0;
    
    while( i < lowHalf.length && j < highHalf.length) {
        if (lowHalf[i] < highHalf[j]) {
            array[k] = lowHalf[i];
            i++;
        } else {
            array[k] = highHalf[j];
            j++;
        }
        k++;
    }
    while (i < lowHalf.length) {
        array[k] = lowHalf[i];
        k++;
        i++;
    }
    while (j < highHalf.length) {
        array[k] = highHalf[j];
        j++;
        k++;
    }    
};
// Takes in an array and recursively merge sorts it
var mergeSort = function(array, p, r) {
    if(r > p) {
        var middle = Math.floor( (p+r)/2 );
        mergeSort(array, p, middle);
        mergeSort(array, middle+1, r);
        merge(array, p, middle, r);
    }
};


var array = [14, 7, 3, 12, 9, 11, 6, 2];
var array2 = [15, 7, 3, 12, 9, 11, 6, -1];
mergeSort(array, 0, array.length-1);
mergeSort(array2, 0, array2.length-1);

console.log(array);
console.log(array2);

