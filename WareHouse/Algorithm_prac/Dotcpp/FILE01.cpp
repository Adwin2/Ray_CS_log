#define _CRT_SECURE_NO_WARNINGS 1   //enale scanf() on VS2022

#include <stdio.h>
#include <stdlib.h>


int compare(const void* a, const void* b)
{
	return (*(int*)a - *(int*)b);
}

int main() {
	//定义部分
	int data[10];
	char outputfilename[200];
	char filename[200];

	//实现部分
	printf("请输入文件a路径\n");
	scanf("%s",filename);
	FILE* InputFile = fopen(filename, "r");
	if (InputFile == NULL) {
		printf("打开文件失败\n");
		return 1;
	}
	for (int i = 0; i < 10; i++) {
		fscanf(InputFile, "%d", &data[i]);
	}
	fclose(InputFile);
	//qsort
	qsort(data, 10, sizeof(int), compare);

	printf("请输入文件b路径");
	scanf_s("%s", outputfilename);
	FILE* OutFile = fopen(outputfilename, "w");
	if (OutFile == NULL) {
		printf("打开文件失败\n");
		return 1;
	}
	for (int m = 0; m < 10; m++) {
		fprintf(OutFile, "%d ", data[m]);
	}
	fclose(OutFile);

	return 0;



}