  #include <stdio.h>
  #include <stdlib.h>
  #include <stdarg.h>
  #include <string.h>
  #define MAX 10000

  void die(const char *format, ...) {
      va_list arglist;
      va_start(arglist, format);
      vprintf(format, arglist);
      va_end(arglist);
      exit(0);
  }

  void info(const char * format, ...) {
      return;
      va_list arglist;
      va_start(arglist, format);
      vprintf(format, arglist);
      va_end(arglist);
  }

  #define PSTATE_NOT_STARTED 0
  #define PSTATE_STOPPED 1
  #define PSTATE_WAITING_INPUT 2
  #define PSTATE_OUTPUT 3

  struct program {
      long *program;
      int pc;
      int state;
      long base;
  };

  long paramref(struct program *program, int mode, int i) {
      if (mode == 0) {
          long val = program->program[i];
          if (val > MAX) die("param too big: mode=%d, i=%d, val=%ld", mode, i, val);
          return program->program[i];
      } else if (mode == 1) {
          return i;
      } else if (mode == 2) {
          long val = program->base + program->program[i];
          if (val > MAX) die("param too big");
          return val;
      } else {
          die("Invalid parameter mode");
      }
  }
  long param(struct program *program, int mode, int i) {
      long val = paramref(program, mode, i);
      if (val > MAX) die("param too big: mode=%d, i=%d, val=%ld", mode, i, val);
      return program->program[val];
  }

  long run(struct program *program, int input, int *readinput) {
      *readinput = 0;
      int stopper = 999999999;
      long *p = program->program;
      long lastoutput = -1;
      int i = program->pc;
      int usedinput = 0;
      while (stopper-- >= 0) {
          if (i > MAX-5) die("pc outside range");
          int opcode = p[i] % 100;
          int p1mode = (p[i] / 100) % 10;
          int p2mode = (p[i] / 1000) % 10;
          int p3mode = (p[i] / 10000) % 10;
          if (opcode == 99) {
              program->state = PSTATE_STOPPED;
              program->pc = i;
              return -1;
          } else if (opcode == 1) { // +
              info("%4d: %ld %ld %ld %ld\n", i, p[i], p[i+1], p[i+2], p[i+3]);
              p[paramref(program, p3mode, i+3)] = param(program, p1mode, i+1) + param(program, p2mode, i+2);
              i += 4;
          } else if (opcode == 2) { // *
              info("%4d: %ld %ld %ld %ld\n", i, p[i], p[i+1], p[i+2], p[i+3]);
              p[paramref(program, p3mode, i+3)] = param(program, p1mode, i+1) * param(program, p2mode, i+2);
              i += 4;
          } else if (opcode == 3) { // input
              info("%4d: %ld %ld\n", i, p[i], p[i+1]);
              if (usedinput) {
                  program->state = PSTATE_WAITING_INPUT;
                  program->pc = i;
                  return -2;
              }
              *readinput = 1;
              p[paramref(program, p1mode, i+1)] = input;
              i += 2;
              usedinput = 1;
          } else if (opcode == 4) { // output
              info("%4d: %ld %ld\n", i, p[i], p[i+1]);
              lastoutput = param(program, p1mode, i+1);
              info("output %ld\n", lastoutput);
              i += 2;
              program->state = PSTATE_OUTPUT;
              program->pc = i;
              return lastoutput;
          } else if (opcode == 5) { // jump-if-true
              info("%4d: %ld %ld %ld\n", i, p[i], p[i+1], p[i+2]);
              if (0 != param(program, p1mode, i+1))
                  i = param(program, p2mode, i+2);
              else
                  i += 3;
          } else if (opcode == 6) { // jump-if-false
              info("%4d: %ld %ld %ld\n", i, p[i], p[i+1], p[i+2]);
              if (0 == param(program, p1mode, i+1))
                  i = param(program, p2mode, i+2);
              else
                  i += 3;
          } else if (opcode == 7) { // less than
              info("%4d: %ld %ld %ld %ld\n", i, p[i], p[i+1], p[i+2], p[i+3]);
              if (param(program, p1mode, i+1) < param(program, p2mode, i+2))
                  p[paramref(program, p3mode, i+3)] = 1;
              else
                  p[paramref(program, p3mode, i+3)] = 0;
              i += 4;
          } else if (opcode == 8) { // equal to
              info("%4d: %ld %ld %ld %ld\n", i, p[i], p[i+1], p[i+2], p[i+3]);
              if (param(program, p1mode, i+1) == param(program, p2mode, i+2))
                  p[paramref(program, p3mode, i+3)] = 1;
              else
                  p[paramref(program, p3mode, i+3)] = 0;
              i += 4;
          } else if (opcode == 9) {
              info("%4d: %ld %ld\n", i, p[i], p[i+1]);
              program->base += param(program, p1mode, i+1);
              i += 2;
          } else {
              info("%ld: %ld %ld %ld %ld\n", i, p[i], p[i+1], p[i+2], p[i+3]);
              die("Error: unknown opcode %d: %ld, %ld, %ld, %ld", i, p[i], p[i+1], p[i+2], p[i+3]);
              return -2;
          }
      }
      die("Error: program outside MAX: %d\n", i);
      return -3;
  }

  #define BUF 200

  char goleft(char d) {
    switch (d) {
    case '^': return '<';
    case '<': return 'v';
    case 'v': return '>';
    case '>': return '^';
    }
  }
  char goright(char d) {
    return goleft(goleft(goleft(d)));
  }

  int main(int argc, char **args) {
      long *positions = malloc(MAX * sizeof(long));
      FILE *f = fopen("17.txt", "r");
      for (int i = 0; i < MAX; i++) {
          if (fscanf(f, "%ld", &positions[i]) != 1) break;
          fscanf(f, ",");
      }

      struct program prog;
      prog.pc = 0;
      prog.state = PSTATE_NOT_STARTED;
      prog.base = 0;
      prog.program = positions;

      char yxs[BUF][BUF];
      for (int y = 0; y < BUF; y++)
          for (int x = 0; x < BUF; x++)
              yxs[y][x] = 0;

      int rx = 1, ry = 1;
      int vx = 0, vy = 0, vdir = 0;

      /* for (int step = 0; step < 20000; step++) { */
      /*   if (ry < 0 || ry >= BUF || rx < 0 || rx >= BUF) die("outside range %d, %d", rx, ry); */
      /*   if (prog.state == PSTATE_STOPPED) break; */
      /*   int readinput = 0; */
      /*   long out = run(&prog, 0, &readinput); */
      /*   if (prog.state == PSTATE_OUTPUT) { */
      /*     if (out == '\n') { */
      /*       rx = 1; */
      /*       ry++; */
      /*     } else { */
      /*       yxs[ry][rx] = (char) out; */
      /*       if (out == '^' || out == 'v' || out == '<' || out == '>') {  */
      /*         vx = rx; vy = ry; vdir = out;  */
      /*       } */
      /*       rx++; */
      /*     } */
      /*   } */
      /* } */

      {
        FILE *f = fopen("17a.txt", "r");
        for (int out = fgetc(f); out != EOF; out = fgetc(f)) {
          if (out == '\n') {
            rx = 1;
            ry++;
          } else {
            yxs[ry][rx] = (char) out;
            if (out == '^' || out == 'v' || out == '<' || out == '>') { 
              vx = rx; vy = ry; vdir = out; 
            }
            rx++;
          }
        }
      }

      int minx = BUF, maxx = 0, miny = BUF, maxy = 0;
      for (int y = 0; y < BUF; y++) {
          for (int x = 0; x < BUF; x++) {
              if (yxs[y][x] != 0) {
                  if (x < minx) minx = x;
                  if (x > maxx) maxx = x;
                  if (y < miny) miny = y;
                  if (y > maxy) maxy = y;
              }
          }
      }
      for (int x = minx; x <= maxx; x++) { yxs[maxy+1][x] = '.'; yxs[0][x] = '.'; }
      for (int y = miny; y <= maxy; y++) { yxs[y][maxx+1] = '.'; yxs[y][0] = '.'; }
      for (int y = miny; y <= maxy; y++) {
        for (int x = minx; x <= maxx; x++) {
          printf("%c", yxs[y][x]);
        }
        printf("\n");
      }

      char moves[1000] = {0};
      int movec = 0;
      while (1) {
        char forward, left, right;
        if (vdir == '^') {
          forward = yxs[vy-1][vx]; left = yxs[vy][vx-1]; right = yxs[vy][vx+1]; }
        if (vdir == 'v') {
          forward = yxs[vy+1][vx]; left = yxs[vy][vx+1]; right = yxs[vy][vx-1]; }
        if (vdir == '<') {
          forward = yxs[vy][vx-1]; left = yxs[vy+1][vx]; right = yxs[vy-1][vx]; }
        if (vdir == '>') {
          forward = yxs[vy][vx+1]; left = yxs[vy-1][vx]; right = yxs[vy+1][vx]; }

        if (forward == '.' && left == '.' && right == '.') break;
        if (forward == '#') die("invalid forward");
        if (left == '#') { moves[movec++] = 'L'; vdir = goleft(vdir); }
        if (right == '#') { moves[movec++] = 'R'; vdir = goright(vdir); }
        //moves[movec++] = ',';
        int dx = 0, dy = 0;
        if (vdir == '^') dy = -1;
        if (vdir == 'v') dy = +1;
        if (vdir == '>') dx = +1;
        if (vdir == '<') dx = -1;
        int dist;
        for (dist = 0; yxs[vy+dy][vx+dx] != '.'; dist++) {
          vy += dy; vx += dx;
          if (vy+dy < 0 || vy+dy >= BUF || vx+dx < 0 || vx+dx >= BUF) die ("dist %d,%d\n", vx, vy);
        }
        moves[movec++] = dist >= 10 ? 'A' + dist-10 : '0' + dist;
        /* char tmp[10]; */
        /* sprintf(tmp, "%d", dist); */
        /* for (int i = 0; tmp[i] != 0; i++) */
        /*   moves[movec++] = tmp[i]; */
        /* moves[movec++] = ','; */
      }
      moves[movec++] = 0;

      int sum = 0;
      for (int y = miny+1; y <= maxy-1; y++) {
        for (int x = minx+1; x <= maxx-1; x++) {
          if (yxs[y][x] == '#' && yxs[y-1][x] == '#' && yxs[y+1][x] == '#' 
              && yxs[y][x-1] == '#' && yxs[y][x+1] == '#') {
              printf("## %d,%d\n", x, y);
              sum += x * y;
      }}}

      printf("moves: %s\n", moves);
      //R5L5L9L7L5L9L5R5L5L9L7L5L9L5R5L7L9R5R5L5L9L7L5L9L5R5L7L9R5R5L5L9R5L7L9R      

      char tmp[1000];
      for (int a = 1; a < movec; a++) {
        for (int b = 1; b < movec - a; b++) {
          for (int c = 1; c < movec - a - b; c++) {
            if (a > 10 || b > 10 || c > 10) continue;
            for (int i = 0; i < movec; i++) tmp[i] = moves[i];
            char as[100], bs[100], cs[100];
            for (int i = 0; i < a; i++) as[i] = tmp[i];
            as[a] = 0;
            while (1) {
              char *ref = strstr(tmp, as);
              if (NULL == ref) break;
              for (int i = ref-tmp; i < movec-a; i++)
                tmp[i] = tmp[i+a];
              break;
            }

            for (int i = 0; i < b; i++) bs[i] = tmp[i];
            bs[b] = 0;
            while (1) {
              char *ref = strstr(tmp, bs);
              if (NULL == ref) break;
              for (int i = ref-tmp; i < movec; i++)
                tmp[i] = tmp[i+b];
            }

            for (int i = 0; i < c; i++) cs[i] = tmp[i];
            cs[c] = 0;
            while (1) {
              char *ref = strstr(tmp, cs);
              if (NULL == ref) break;
              for (int i = ref-tmp; i < movec; i++)
                tmp[i] = tmp[i+c];
            }

            if (strlen(tmp) < 10) {
              printf("## a=%d (%s), b=%d (%s), c=%d (%s) left=%s\n", a, as, b, bs, c, cs, tmp);
              if (tmp[0] == 0) {
                printf("!\n");
                for (int i = 0; i < movec; i++) tmp[i] = moves[i];
                char *ref = tmp;
                while (ref[0] != 0) {
                  if (NULL != strstr(tmp, as)) {
                    ref += a; printf("A,");
                  }
                  if (NULL != strstr(tmp, bs)) {
                    ref += b; printf("B,");
                  }
                  if (NULL != strstr(tmp, cs)) {
                    ref += c; printf("C,");
                  }
                }
                printf("\n");
              }
            }
      }}}

      /* { */
      /* char tmp[1000]; */
      /* for (int ss1 = 0; ss1 < movec; ss1+=2) { */
      /*   for (int ss2 = ss1+4; ss2 < movec; ss2+=2) { */
      /*     for (int i = 0; i < ss2-ss1; i++) { */
      /*       tmp[i] = moves[ss1+i]; */
      /*     } */
      /*     tmp[ss2-ss1] = 0; */

      /*     int seens = 0; */
      /*     char *ref = moves; */
      /*     for (seens = 0; 1; seens++) { */
      /*       ref = strstr(ref, tmp); */
      /*       if (NULL == ref) break; */
      /*       ref++; */
      /*     } */
      /*     if (strlen(tmp) * seens >= 30) */
      /*       printf("## %d   '%s'\n", seens, tmp); */
      /*   } */
      /* } */
      /* } */

      return 0;
  }
