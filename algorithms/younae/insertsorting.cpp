#include <iostream>
using namespace std;

void insertsorting(int* arr, int size);
int main()
{
	int random[5];

	printf("array initialization : " );
	for(int i=0; i<5; i++)
	{
		random[i] = rand()%100;

		printf("%d ", random[i]);
	}
	printf("\n\n");
	printf("insetsorting start \n");
	insertsorting(random, sizeof(random)/sizeof(int));

	return 0;
}
void insertsorting(int* arr, int size)
{
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

			for (int k = 0; k < size; k++)
			{
				printf("%d ", arr[k]);
			}
			printf("\n");
		}
	}
}
