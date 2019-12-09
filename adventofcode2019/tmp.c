  #include <stdio.h>
  #include <stdlib.h>
  #include <stdarg.h>
  #define MAX 10000

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
    long program[MAX];
    int pc;
    int state;
    long base;
  };

  long param(struct program *program, int mode, int i) {
    switch (mode) {
    case 0: return program->program[program->program[i]];
    case 1: return program->program[i];
    case 2: return program->program[program->base + program->program[i]];
    default: die("Invalid parameter mode");
    }
    return 0;
  }

  long run(struct program *program, int input) {
    int stopper = 10000;
    long *p = program->program;
    long lastoutput = -1;
    int i = program->pc;
    int usedinput = 0;
    while (stopper-- >= 0) {
       if (i > MAX-5) die("pc outside range");
        int opcode = p[i] % 100;
        int p1mode = (p[i] / 100) % 10;
        long p1 = param(program, p1mode, i+1);
        int p2mode = (p[i] / 1000) % 10;
        long p2 = param(program, p2mode, i+2);
        int p3mode = (p[i] / 10000) % 10;
        long p3 = param(program, p3mode, i+3);
        if (opcode == 99) {
          program->state = PSTATE_STOPPED;
          program->pc = i;
          return -1;
        } else if (opcode == 1) { // +
          info("%4d: %ld %ld %ld %ld\n", i, p[i], p[i+1], p[i+2], p[i+3]);
          if (p3mode == 1) die("opcode 1 had p3mode");
          p[p[i+3]] = p1 + p2;
          i += 4;
        } else if (opcode == 2) { // *
          info("%4d: %ld %ld %ld %ld\n", i, p[i], p[i+1], p[i+2], p[i+3]);
          if (p3mode == 1) die("opcode 2 had p3mode");
          p[p[i+3]] = p1 * p2;
          i += 4;
        } else if (opcode == 3) { // input
          info("%4d: %ld %ld\n", i, p[i], p[i+1]);
          if (usedinput) {
            program->state = PSTATE_WAITING_INPUT;
            program->pc = i;
            return -2;
          }
          p[p[i+1]] = input;
          i += 2;
          usedinput = 1;
        } else if (opcode == 4) { // output
          info("%4d: %ld %ld\n", i, p[i], p[i+1]);
          lastoutput = p1;
          info("output %ld\n", lastoutput);
          i += 2;
          program->state = PSTATE_OUTPUT;
          program->pc = i;
          return lastoutput;
        } else if (opcode == 5) { // jump-if-true
          info("%4d: %ld %ld %ld\n", i, p[i], p[i+1], p[i+2]);
          if (0 != p1)
            i = p2;
          else
            i += 3;
        } else if (opcode == 6) { // jump-if-false
          info("%4d: %ld %ld %ld\n", i, p[i], p[i+1], p[i+2]);
          if (0 == p1)
            i = p2;
          else
            i += 3;
        } else if (opcode == 7) { // less than
          info("%4d: %ld %ld %ld %ld\n", i, p[i], p[i+1], p[i+2], p[i+3]);
          if (p1 < p2)
            p[p[i+3]] = 1;
          else
            p[p[i+3]] = 0;
          i += 4;
        } else if (opcode == 8) { // equal to
          info("%4d: %ld %ld %ld %ld\n", i, p[i], p[i+1], p[i+2], p[i+3]);
          if (p1 == p2)
            p[p[i+3]] = 1;
          else
            p[p[i+3]] = 0;
          i += 4;
        } else if (opcode == 9) {
          info("%4d: %ld %ld\n", i, p[i], p[i+1]);
          program->base += p1;
          i += 2;
        } else {
          info("%ld: %ld %ld %ld %ld\n", i, p[i], p[i+1], p[i+2], p[i+3]);
          die("Error: unknown opcode");
          return -2;
        }
      }
      die("Error: program outside MAX\n");
      return -3;
  }

  int main(int argc, char **args) {
      long positions[MAX];
      FILE *f = fopen("09b.txt", "r");
      for (int i = 0; i < MAX; i++) {
          if (fscanf(f, "%ld", &positions[i]) != 1) break;
          fscanf(f, ",");
      }

      struct program prog;
      prog.pc = 0;
      prog.state = PSTATE_NOT_STARTED;
      prog.base = 0;
      for (int i = 0; i < MAX; i++)
        prog.program[i] = positions[i];

      while (prog.state != PSTATE_STOPPED) {
        long output = run(&prog, 0);
        if (prog.state == PSTATE_OUTPUT)
          printf("OUTPUT %ld\n", output);
      }
      return 0;
  }
