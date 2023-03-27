#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINES 3000
#define MAX_LINE_LENGTH 256

char **conv_file(const char *target, int *num_lines);
void prt_str_arr(char **lines, int num_lines);

int main() {
    int num_lines = 0;
    char **lines = conv_file("input.txt", &num_lines);

    if (lines != NULL) {
        prt_str_arr(lines, num_lines);

        for (int i = 0; i < num_lines; i++) {
            free(lines[i]);
        }
        free(lines);
    }

    return 0;
}

char **conv_file(const char *target, int *num_lines) {
    FILE *file;
    char line[MAX_LINE_LENGTH];
    char **lines = malloc(MAX_LINES * sizeof(char *));
    *num_lines = 0;

    if (lines == NULL) {
        printf("Error can't allocate mem, lines is NULL \n");
        return NULL;
    }

    file = fopen(target, "r");
    if (file == NULL) {
        printf("Error unable to open file, file is NULL \n");
        free(lines);
        return NULL;
    }

    while(fgets(line, sizeof(line), file) != NULL && *num_lines < MAX_LINES) {
        lines[*num_lines] = malloc(MAX_LINE_LENGTH * sizeof(char));
        if (lines[*num_lines] == NULL) {
            printf("Error unable to allocate mem to lines index: %d \n", *num_lines);
            break;
        }
        strncpy(lines[*num_lines], line, MAX_LINE_LENGTH);
        lines[*num_lines][MAX_LINE_LENGTH - 1] = '\0';
        (*num_lines)++;
    }

    fclose(file);
    return lines;
}

void prt_str_arr(char **lines, int num_lines) {
    for (int i = 0; i < num_lines; i++) {
        printf("%s", lines[i]);
    }
}
