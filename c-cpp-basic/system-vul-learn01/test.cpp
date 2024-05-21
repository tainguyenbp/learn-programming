#include <stdio.h>

int main(int argc, char **argv)
{
  volatile int important_data = 0;
  char user_input[10];

  gets(user_input);  // Vulnerable function

  if(important_data != 0) {
    printf("Warning !!!, the 'important_data' was changed\n");
  } else {
    printf("the 'important_data' was not changed\n");
  }
}
