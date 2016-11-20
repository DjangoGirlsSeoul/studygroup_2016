void insertsorting(int* arr)
{
	int size = sizeof(arr)/sizeof(int);
	int temp =0;

	for(int i=1; i<size; i++)
	{
		temp = arr[i];

		for(int j=i; j>0; j--)
		{
			if(arr[j-1]>temp)
			{
				arr[j] = arr[j-1];
				arr[j-1] = temp;
			}
			else
				break;
			print(" %d ", arr[j]);	
		}
		printf("\n");
	}
}