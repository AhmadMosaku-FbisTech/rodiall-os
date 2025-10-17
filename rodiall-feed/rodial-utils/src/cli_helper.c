/*
 * Rodial Utils - CLI Helper (C Core)
 * Provides lightweight CLI parsing and logging used across
 * Rodial OS daemons and tools.
 */

#include "cli_helper.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

void rodial_print_banner(const char *service, const char *version) {
    time_t t = time(NULL);
    struct tm *tm_info = localtime(&t);

    printf("\n🛰️  %s (%s)\n", service, version);
    printf("───────────────────────────────────────────────\n");

    char buffer[64];
    strftime(buffer, sizeof(buffer), "%a %b %d %H:%M:%S %Y", tm_info);
    printf("Build Time: %s\n", buffer);
    printf("Rodial OS © 2025 — All rights reserved.\n\n");
}

void rodial_log_info(const char *msg) {
    printf("[INFO] %s\n", msg);
}

void rodial_log_error(const char *msg, int fatal) {
    fprintf(stderr, "[ERROR] %s\n", msg);
    if (fatal) exit(EXIT_FAILURE);
}

void rodial_log_debug(const char *msg) {
#ifdef DEBUG
    printf("[DEBUG] %s\n", msg);
#endif
}

int rodial_file_exists(const char *path) {
    FILE *file = fopen(path, "r");
    if (file) {
        fclose(file);
        return 1;
    }
    return 0;
}
