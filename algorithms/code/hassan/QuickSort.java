
/**
 * QuickSort algorithms
 * 
 * @Hassan Abid
 * @version 1.0
 */
public class QuickSort
{

    
    public static void main(int[] arg) 
    {
        
        int[] inputList = {2, 3, 1, 7, 8 , 10 , 15 , 11, 5};
        QuickSort.quicksort(inputList);
        
        System.out.println("sorted array: " + inputList.toString());
       System.out.println("sorted array [array.length - 1]: " + inputList[inputList.length - 1]);
    }

    /**
     * An example of a method - replace this comment with your own
     * 
     * @param  y   a sample parameter for a method
     * @return     the sum of x and y 
     */
    public static void quicksort(int[] array)
    {
        
        quicksort(array, 0, array.length - 1);
        
    }
    
    public static void quicksort(int[] array, int left, int right) 
    {
    
        if (left >= right) {
            return;
        }
        
        int pivot = array[ (left + right) / 2];
        int index = partition(array, left, right, pivot);
        quicksort(array, left, index - 1);
        quicksort(array, index, right);

    
    }
    
    public static int partition(int[] array, int left, int right, int pivot) 
    {
    
       System.out.println("right : " + right);

        while (left <= right) {
        
            while(array[left] < pivot) {
                left++;
            }
            
            while(array[right] > pivot) {
            
                right--;
            }
            
            if(left <= right) {
            
                swap(array, left, right);
                left++;
                right--;
            }
            
        }
        
        
        return left;
    
    
    }
    
     public static void swap(int[] array, int m, int n) 
     {
    
        int temp = array[m];
        array[m] = array[n];
        array[n] = temp;
    }
    
}
