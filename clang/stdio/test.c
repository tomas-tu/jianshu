#include <stdio.h>

int main()
{
    char str[10];
    fscanf(stdin, "%s", str);

    fprintf(stdout, "%s", "hi,");

    fprintf(stderr, "%s", "clang! ");

    return 0;
}

//注意输出的结果是 clang! hi,

