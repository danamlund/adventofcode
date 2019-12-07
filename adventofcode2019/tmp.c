  #include <stdio.h>
  #include <stdlib.h>
  #include <stdarg.h>
  #define MAX 100000

  void die(const char *s) {
    printf("%s\n", s);
    exit(0);
  }

  void info(const char * format, ...) {
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
    int program[MAX];
    int pc;
    int state;
  };

  int run(struct program *program, int input) {
    int *p = program->program;
    int lastoutput = -1;
    int i = program->pc;
    int usedinput = 0;
    while (i < MAX) {
        int opcode = p[i] % 100;
        int p1mode = (p[i] / 100) % 10;
        int p2mode = (p[i] / 1000) % 10;
        int p3mode = (p[i] / 10000) % 10;
        if (opcode == 99) {
          program->state = PSTATE_STOPPED;
          program->pc = i;
          return -1;
        } else if (opcode == 1) { // +
          info("%4d: %d %d %d %d\n", i, p[i], p[i+1], p[i+2], p[i+3]);
          if (p3mode == 1) die("opcode 1 had p3mode");
          p[p[i+3]] = (p1mode ? p[i+1] : p[p[i+1]]) + (p2mode ? p[i+2] : p[p[i+2]]);
          i += 4;
        } else if (opcode == 2) { // *
          info("%4d: %d %d %d %d\n", i, p[i], p[i+1], p[i+2], p[i+3]);
          if (p3mode == 1) die("opcode 2 had p3mode");
          p[p[i+3]] = (p1mode ? p[i+1] : p[p[i+1]]) * (p2mode ? p[i+2] : p[p[i+2]]);
          i += 4;
        } else if (opcode == 3) { // input
          info("%4d: %d %d\n", i, p[i], p[i+1]);
          if (usedinput) {
            program->state = PSTATE_WAITING_INPUT;
            program->pc = i;
            return -2;
          }
          p[p[i+1]] = input;
          i += 2;
          usedinput = 1;
        } else if (opcode == 4) { // output
          info("%4d: %d %d\n", i, p[i], p[i+1]);
          lastoutput = p1mode ? p[i+1] : p[p[i+1]];
          info("output %d\n", lastoutput);
          i += 2;

          program->state = PSTATE_OUTPUT;
          program->pc = i;
          return lastoutput;
        } else if (opcode == 5) { // jump-if-true
          info("%4d: %d %d %d\n", i, p[i], p[i+1], p[i+2]);
          if (0 != (p1mode ? p[i+1] : p[p[i+1]]))
            i = p2mode ? p[i+2] : p[p[i+2]];
          else
            i += 3;
        } else if (opcode == 6) { // jump-if-false
          info("%4d: %d %d %d\n", i, p[i], p[i+1], p[i+2]);
          if (0 == (p1mode ? p[i+1] : p[p[i+1]]))
            i = p2mode ? p[i+2] : p[p[i+2]];
          else
            i += 3;
        } else if (opcode == 7) { // less than
          info("%4d: %d %d %d %d\n", i, p[i], p[i+1], p[i+2], p[i+3]);
          if ((p1mode ? p[i+1] : p[p[i+1]]) < (p2mode ? p[i+2] : p[p[i+2]]))
            p[p[i+3]] = 1;
          else
            p[p[i+3]] = 0;
          i += 4;
        } else if (opcode == 8) { // equal to
          info("%4d: %d %d %d %d\n", i, p[i], p[i+1], p[i+2], p[i+3]);
          if ((p1mode ? p[i+1] : p[p[i+1]]) == (p2mode ? p[i+2] : p[p[i+2]]))
            p[p[i+3]] = 1;
          else
            p[p[i+3]] = 0;
          i += 4;
        } else {
          info("%d: %d %d %d %d\n", i, p[i], p[i+1], p[i+2], p[i+3]);
          die("Error: unknown opcode");
          return -2;
        }
      }
      die("Error: program outside MAX\n");
      return -3;
  }

  int main(int argc, char **args) {
      int positions[MAX];
      FILE *f = fopen("07b.txt", "r");
      for (int i = 0; i < MAX; i++) {
          if (fscanf(f, "%d", &positions[i]) != 1) break;
          fscanf(f, ",");
      }

      int max = 0;
      /* for (int p1 = 5; p1 <= 9; p1++) { */
      /*   for (int p2 = 0; p2 <= 9; p2++) { */
      /*     if (p2 == p1) continue; */
      /*     for (int p3 = 5; p3 <= 9; p3++) { */
      /*       if (p3 == p2 || p3 == p1) continue; */
      /*       for (int p4 = 5; p4 <= 9; p4++) { */
      /*         if (p4 == p3 || p4 == p2 || p4 == p1) continue; */
      /*         for (int p5 = 5; p5 <= 9; p5++) { */
      /*           if (p5 == p4 || p5 == p3 || p5 == p2 || p5 == p1) continue; */
                int p1 = 9;
                int p2 = 8;
                int p3 = 7;
                int p4 = 6;
                int p5 = 5;
                struct program progs[6];
                for (int j = 0; j < 5; j++) {
                  progs[j].pc = 0;
                  progs[j].state = PSTATE_NOT_STARTED;
                  for (int i = 0; i < MAX; i++) 
                    progs[j].program[i] = positions[i];
                }

                printf("1\n");
                run(&progs[1], p1);
                printf("2\n");
                run(&progs[2], p2);
                printf("3\n");
                run(&progs[3], p3);
                printf("4\n");
                run(&progs[4], p4);
                printf("5\n");
                run(&progs[5], p5);

                int output = 0;
                while (progs[1].state != PSTATE_STOPPED) {
                  output = run(&progs[1], output);
                  output = run(&progs[2], output);
                  output = run(&progs[3], output);
                  output = run(&progs[4], output);
                  output = run(&progs[5], output);
                }
                if (output > max) max = output;
      /* }}}}} */
      printf("= %d\n", max);
      return 0;
  }
