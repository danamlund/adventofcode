  #include <stdio.h>
  #include <math.h>
  int gcd(int a, int b) { 
    if (a == 0 || b == 0) return 0;
    else if (a == b) return a;
    else if (a > b) return gcd(a - b, b);
    else/*if (a < b)*/return gcd(a, b - a);
  }
  int abs(int a) {
    return a < 0 ? -a: a;
  }
  int main(int argc, char **args) {
    int yxs[30][30] = {0};
    FILE *f = fopen("10b.txt", "r");
    int w, h;
    {
      int c, x, y;
      x = y = w = h = 0;
      while ((c = fgetc(f)) != EOF) {
        if (c == '\n') {
          y++;
          x = 0;
        } else {
          if (x > w) w = x;
          if (y > h) h = y;
          if (c == '#') yxs[y][x] = 1;
          x++;
        }
      }
      w++;
      h++;
    }

    //int x = 17, y = 23; 10.txt
    int x = 8, y = 3; // 10b.txt

    float radians[30][30] = {-1};
    int detecteds[30][30] = {0};
    int maxDetected = 0;
    int maxx, maxy;

    printf("%f\n", atan2(1.0f, 3.0f));

    //printf("w=%d, h=%d\n", w, h);
    int nth = 1;
    for (int rotations = 0; rotations < 1; rotations++) {
      float radian = 0.0f;
      float minradian = 99.0f;
      int mindist = 9999;
      int minx, miny;
          for (int yy = 0; yy < h; yy++) {
            for (int xx = 0; xx < w; xx++) {
              if (yxs[yy][xx]) {
                float rad = atan2((float) (y - yy), (float) (xx -x));
                if (rad < 0) rad = 2 * M_PI + rad;
                int dist = abs(y - yy) + abs(x - xx);
                float radianToLazer = radian - rad;
                if (radianToLazer >= 0.0f && radianToLazer < minradian) {
                  minradian = radianToLazer;
                  minx = xx;
                  miny = yy;
                }
                if (radianToLazer == minradian && dist < mindist) {
                  mindist = dist;
                  minx = xx;
                  miny = yy;
                }
              }
            }
          }

          printf("## %d: %d,%d rad=%.2f dist=%d\n", nth, minx, miny, minradian, mindist);
          yxs[miny][minx] = 0;
          nth++;
          radian = minradian;
    }

    //printf("= %d\n", maxDetected);
    return 0;
  }
