  #include <stdio.h>
  #include <stdlib.h>
  #define BUF 10000000
  #define DUP 10000
  /* long abs(long val) { */
  /*   return val < 0 ? -val : val; */
  /* } */
  int main(int argc, char **args) {
    FILE *f = fopen("16b.txt", "r");

    char *list = malloc(BUF);
    int listc = 0;
    for (int b = fgetc(f); b >= '0' && b <= '9'; b = fgetc(f)) {
      list[listc++] = b - '0';
    }
    int listc1 = listc;
    for (int i = 0; i < DUP; i++) {
      for (int j = 0; j < listc; j++) {
        list[i * listc + j] = list[j];
      }
    }
    listc = DUP * listc;

    long *sum = malloc(sizeof(long) * (BUF+1));

    char *list2 = malloc(BUF);
    for (int step = 0; step < 100; step++) {
      sum[0] = 0;
      for (int i = 0; i < listc; i++) sum[i+1] = sum[i] + list[i];
      //printf("... %d\n", step);
      if (step >= 1) {
      for (int i = 0; i < listc; i++) {
      //for (int i = 800; i < 805; i++) {
      //{ int i = 801;
        long val = 0;

        //printf("#0 %2d: patc=%d, listc1=%d\n", i, patc, listc1);
        if (1) {
          int patc = 4*(i+1);
          for (int j = 0; j < listc; j += patc) {
            int ones1start = j + 1*(i+1) - 1;
            int ones1end = j + 2*(i+1) - 1;
            if (ones1end >= listc) ones1end = listc;
            if (ones1start < ones1end)
              val += sum[ones1end] - sum[ones1start];

            int ones2start = j + 3*(i+1) - 1;
            int ones2end = j + 4*(i+1) - 1;
            if (ones2end >= listc) ones2end = listc;
            if (ones2start < ones2end)
              val -= sum[ones2end] - sum[ones2start];
            /* printf("#8 i=%d, j=%d, listc=%d, +=%d..%d=%ld, -=%d..%d=%d, val=%d\n", i, j, listc, ones1start, ones1end, sum[ones1end] - sum[ones1start], ones2start, ones2end, sum[ones2end] - sum[ones2start], val); */
            /* printf("  +: %d - %d  -: %d - %d\n", sum[ones1end], sum[ones1start], sum[ones2end], sum[ones2start]); */
          }
        /* } else if (5*(i+1) - 1 >= listc) { */
        /*   int ones1start = 1*(i+1) - 1; */
        /*   int ones1end = 2*(i+1) - 2; */
        /*   if (ones1end >= listc) ones1end = listc; */
        /*   if (ones1start < listc) */
        /*     val += sum[ones1end] - sum[ones1start-1]; */

        /*   int ones2start = 3*(i+1) - 1; */
        /*   int ones2end = 4*(i+1) - 2; */
        /*   if (ones2end >= listc) ones2end = listc; */
        /*   if (ones2start < listc) */
        /*     val -= sum[ones2end] - sum[ones2start-1]; */

          //printf("#8 i=%d, listc=%d, +=%d..%d=%ld, -=%d..%d=%d, val=%d\n", i, listc, ones1start, ones1end, sum[ones1end] - sum[ones1start], ones2start, ones2end, sum[ones2end] - sum[ones2start], val);
          //printf("  +: %d - %d  -: %d - %d\n", sum[ones1end], sum[ones1start], sum[ones2end], sum[ones2start]);
        /* } else { */
        /* if (i >= listc1) { */
        /*   int ones1start = 1*(i+1) - 1; */
        /*   int ones1end = 2*(i+1) - 1; */
        /*   int ones2start = 3*(i+1) - 1; */
        /*   int ones2end = 4*(i+1) - 1; */

        /*   int jmodval = 0; */
        /*   for (int j = ones1start; j < ones1end; j++) { */
        /*     val += list[j]; */
        /*     if (j > ones1start && (ones1start % listc1) == (j % listc1)) { */
        /*       int loops = (ones1end - ones1start) / listc1; */
        /*       val += jmodval * loops; */
        /*       printf("#6 start=%d, end=%d, loops=%d, len=%d, listc1=%d, val=%d\n",  */
        /*       ones1start, j, loops, ones1end - ones1start, listc1, jmodval); */
        /*       j += loops * listc1; */
        /*     } */
        /*     jmodval += list[j]; */
        /*     printf("#1 i=%2d, j=%2d, +1, listc1=%d, in=%d, val=%ld\n", i, j, j % listc1, list[j], val); */
        /*   } */
        /*   jmodval = 0; */
        /*   for (int j = ones2start; j < ones2end; j++) { */
        /*     val -= list[j]; */
        /*     if (j > ones2start && ones2start % listc1 == j % listc1) { */
        /*       int loops = (ones2end - ones2start) / listc1; */
        /*       val -= jmodval * loops; */
        /*       printf("#7 start=%d, end=%d, loops=%d, val=%d\n", ones2start, j, loops, jmodval); */
        /*       j += loops * listc1; */
        /*     } */
        /*     jmodval += list[j]; */
        /*     printf("#1 i=%2d, j=%2d, -1, listc1=%d, in=%d, val=%ld\n", i, j, j % listc1, list[j], val); */
        /*   } */
        /* } else { */
        /*   int firstPat = -1; */
        /*   int firstJc1 = -1; */
        /*   for (int j = 0; j < listc; j++) { */
        /*     int pat = j+1; */
        /*     pat = pat % (4*(i+1)); */

        /*     if (firstPat == -1) { */
        /*       firstPat = pat; */
        /*       firstJc1 = j % listc1; */
        /*     } else if (firstPat == pat && firstJc1 == (j % listc1)) { */
        /*       int jloops = listc / j; */
        /*       val *= jloops; */
        /*       //printf("#3 jloops=%d, left=%d, val=%ld\n", jloops, listc - j * jloops, val); */
        /*       j = j * jloops; */
        /*     } */

        /*     int a = 0; */
        /*     if (pat < i+1); */
        /*     else if (pat < 2*(i+1)) { */
        /*       val += list[j]; a = 1; */
        /*     } */
        /*     else if (pat < 3*(i+1)) a = 2; */
        /*     else if (pat < 4*(i+1)) { */
        /*       val -= list[j]; a = 3; */
        /*     } */
        /*     //printf("#1 i=%2d, j=%2d, pat=%2d, listc1=%d, in=%d, val=%ld, pat=%d, a=%d, sum=%ld\n", i, j, pat, j % listc1, list[j], val, pat / (i+1), a, sum[j]); */
        /*   } */
        }
        list2[i] = labs(val) % 10;
        //printf("#4 i=%2d, val=%ld, newval=%d\n", i, val, list2[i]);
      }

      for (int i = 0; i < listc; i++)
        list[i] = list2[i];
      }

      /* printf("%2d ", step); */
      /* for (int i = 0; i < listc; i++) { */
      /*   if (i % listc1 == 0) printf(" "); */
      /*   printf("%d", list[i]); */
      /* } */
      /* printf("\n"); */
    }


    char indexs[9];
    for (int i = 0; i < 8; i++)
      indexs[i] = '0'+list[i];
    indexs[8] = 0;

    int index = atoi(indexs);

    for (int i = 0; i < 8; i++)
      indexs[i] = '0'+list[i + index];
    int msg = atoi(indexs);

    printf("= %d (%d)\n", msg, index);
    return 0;
  }
