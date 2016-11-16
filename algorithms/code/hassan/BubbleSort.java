/**
 * Java implementation of Bubble Sort
 * 
 * @author Hassan Abid
 * @version 1.0
 */
public class BubbleSort
{

    /**
     * BubbleSort an input list of integers
     * 
     * @param inputList an array of integers
     * 
     * Time complexity : O(N^2)
     */
    public static void doBubbleSort(int[] inputList)
    {
         boolean isSorted = false;
        int numberOfSwaps = 0;
        int lastunSorted = inputList.length -1;
        
       while (!isSorted) {
 
            isSorted = true;
            for (int j = 0; j < lastunSorted; j++) {
            if (inputList[j] > inputList[j + 1]) {
                swap(inputList, j, j + 1);
                numberOfSwaps++;
                isSorted = false;
            }
        }
        lastunSorted--;
       }
  
       System.out.println("Array is sorted in " + numberOfSwaps + " swaps.");
       System.out.println("First Element: " + inputList[0]);
       System.out.println("Last Element: " +  inputList[inputList.length - 1]);
    }
    
    public static void main(String[] args) {
   
        int[] inputList = {2, 3, 1, 7, 8 , 10 , 15 , 11, 19};
        BubbleSort.doBubbleSort(inputList);
      

    }
    
    public static void swap(int[] array, int m, int n) {
    
        int temp = array[m];
        array[m] = array[n];
        array[n] = temp;
    }
}