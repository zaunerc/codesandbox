#include <stdio.h>
#include <unistd.h>
#include <signal.h>
#include <stdbool.h>
#include <stdlib.h>

bool end_program = false;

void sig_handler(int signo)
{
    printf("Received a signal (%d) indicating that the program should stop.\n", signo);
    end_program = true;
}

int main (void)
{
    /*
     * https://stackoverflow.com/a/1716621
     * 
     * Writting to stdout is buffered by default if stdout is connected to a file.
     * Therefore, if the process is sent the SIGKILL signal application logs
     * may be lost.
     */
    //setbuf(stdout, NULL);
    
    if (signal(SIGINT, sig_handler) == SIG_ERR)
        printf("\nCan't catch SIGINT.\n");

    if (signal(SIGTERM, sig_handler) == SIG_ERR)
        printf("\nCan't catch SIGTERM.\n");

    for (int i = 0; i < 99999 && !end_program; i++) {
        
        fprintf(stdout, "STDOUT: Counter value: %d\n", i);
        //fprintf(stderr, "STDERR: Counter value: %d\n", i);
       
        /*
        if (i == 4) {
            fprintf(stderr, "STDERR 1: Calling exit(1)\n");
            fprintf(stdout, "STDOUT 2: Calling exit(1)\n");
            fprintf(stderr, "STDERR 3: Calling exit(1)\n");
            fprintf(stdout, "STDOUT 4: Calling exit(1)\n");
            exit(1);
        }
        */
        
        //sleep(1);
    }

    printf("Counter is going to exit.\n");
    return 0;
}
