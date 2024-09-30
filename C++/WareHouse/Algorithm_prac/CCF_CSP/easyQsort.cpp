#include<stdio.h>
#include<stdlib.h>

int compare(const void* a, const void* b){
	return (*(int*)b - *(int*)a);
}

int main() {

    int num1, num2, num3;

    printf("输入三个整数：\n");

    scanf_s("%d,%d,%d", &num1, &num2, &num3);
    int a[3];
    a[0] = num1;
    a[1] = num2;
    a[2] = num3;
    int* p = a;

    qsort(p, 3, sizeof(*p), compare);

    printf("大到小：");
    for (int m = 0; m < 2; m++) {
        printf("%d,", *(p+m));
    }
    printf("%d", *(p + 2));
    printf("\n");

    return 0;
}
