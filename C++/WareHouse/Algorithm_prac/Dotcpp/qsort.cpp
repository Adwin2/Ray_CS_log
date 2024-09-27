//#include<stdio.h>
//#include<stdlib.h>
//
//
//int compare(const void * a, const void * b)
//{
//	return (*(int*)b - *(int*)a);
//}
//int main()
//{
//	int i, j;
//	int a[15] = { 0 };
//	int* p = a;
//	printf("please input a[15]:\n");
//	for (int m = 0; m < 15; m++) {
//		scanf_s("%d", &a[m]);
//	}
//	p = a;
//	qsort(p, 15, sizeof(*p), compare);
//
//	for (i = 0; i < 15; i++) {
//		printf("%d", *(p + i));
//		printf("\n");
//	}
//	
//	return 0;
//}