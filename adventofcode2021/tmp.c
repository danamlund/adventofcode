#include <stdio.h>
#include <stdlib.h>
int main() {
  FILE *f = fopen("15t.txt", "r");

  char yxs[105][105] = {0};
  char line[1024];
  int size = 0;
  while (fgets(line, sizeof(line), f)) {
    for (int i = 0; line[i] != '\n'; i++)
      yxs[size][i] = line[i]-'0';
    size++;
  }
  if (0) {
  for (int y = 0; y < size; y++) {
    for (int x = 0; x < size; x++) {
      printf("%c", '0'+yxs[y][x]);
    }
    printf("\n");
  }
  }

  int yxscore[105*5][105*5] = {0};
  for (int y = 0; y < size*5; y++) {
    for (int x = 0; x < size*5; x++) {
      if (x == 0 && y == 0) continue;
      int score = 999999;
      if (y > 0 && yxscore[y-1][x] < score) score = yxscore[y-1][x];
      if (x > 0 && yxscore[y][x-1] < score) score = yxscore[y][x-1];
      int gridscore = (yxs[y%size][x%size] + x/size + y/size) % 10;
      yxscore[y][x] = score + gridscore;
    }
  }

  printf("%d\n", yxscore[5*size-1][5*size-1]);
  return 0;
}
