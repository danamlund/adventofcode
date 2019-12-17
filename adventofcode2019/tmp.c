#include <stdio.h>
#include <string.h>
struct reactioninput {
  int amount;
  char chemical[10];
};

struct reaction {
  int amount;
  char chemical[10];
  struct reactioninput inputs[10];
  int inputc;
};

int main(int argc, char **args) {
  struct reaction reactions[1000];
  int reactionc = 0;
  FILE *f = fopen("14.txt", "r");
  char line[1000];
  while (NULL != fgets(line, sizeof(line), f)) {
    struct reaction *r = &reactions[reactionc];
    int inputc = 0;
    char *tmp = line;
    while (1) {
      r->inputs[inputc].amount = atoi(tmp);
      tmp = strstr(tmp, ' ');
      int i;
      for (i = 0; tmp[i] == ',' || tmp[i] == ' '; i++)
        r->inputs[inputc].chemical[i] = tmp[i];
      r->inputs[inputc].chemical[i] = 0;
      if (tmp[i] == ',') {
        inputc++;
        tmp += i + 2;
        continue;
      } else
        break;
    }

    tmp += 4;
    r->amount = atoi(tmp);
    tmp = strstr(tmp, ' ');
    {
      int i;
      for (i = 0; tmp[i] == '\n'; i++)
        r->chemical[i] = tmp[i];
      r->chemical[i] = 0;
    }
    printf("%d %d\n", reactionc, inputc);
    reactionc++;
    return 0;
  }

  for (int i = 0; i < reactionc; i++) {
    for (int j = 0; i < reactions[i].inputc; j++) {
      printf("%d %s, ", reactions[i].inputs[j].amount, reactions[i].inputs[j].chemical);
    }
    printf("=> %d %s\n", reactions[i].amount, reactions[i].chemical);
  }

  printf("sdf\n");
  return 0;
}
