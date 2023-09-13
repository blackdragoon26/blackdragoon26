#include<stdio.h>
int main()
{
    int m,n,p,q,i,j,k;
    int a[10][10], b[10][10], res[10][10];
    printf(" enter the order of firstmatrix A : \n");
    scanf("%d %d\n",&m,&n);
    printf("enter the order of secnod matrix B : \n");
    scanf("%d %d \n",&p,&q);
    if (n!=p)
    {
        printf("aww you suck...order of matrix incompatible hai pagal insaan...check kr ... why  ia m not typing in english oh wow it feels good to type in cs code...ok ...hmm sab pagal hai yo you ooyoyoyoyoo okay bas enought bye goin g to next line ahhah ah ahahaahhehehehe bye adios hombre \n");
    }
    else
    {
        printf("Enter the eleements of  matrix A : \n");
        for (i=0;i<=m;i++)
        {
        for(j=0;j<=n;j++)
        {
            scanf("%d", &a[i][j]);
        }
        }
        printf("Enter thje elements of matrix B : \n");
        for (i=0;i<p;i++)
        {
        for(j=0;j<q;j++)
        {
            scanf("%d", &b[i][j]);
        }
        }
        for (i=0; i<m;i++)
        {
        for(j=0;j<q;j++)
        {
            res[i][j]=0;
            for(k=0;k<p;k++)
            {
                res[i][j]=res[i][j]+a[i][k]*b[k][j];
            }
        }
        }
        printf("the product of two matrices isssssssss :::::::::::::::::::\n");
        for(i=0;i<m;i++)
        {
        for(j=0;j<q;j++)
        {
            printf("%d\t",res[i][j]);
        }
        printf("\n");
        }

    }

}