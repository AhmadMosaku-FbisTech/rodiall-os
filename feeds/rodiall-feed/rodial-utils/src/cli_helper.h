#ifndef RODIAL_CLI_HELPER_H
#define RODIAL_CLI_HELPER_H

#ifdef __cplusplus
extern "C" {
#endif

void rodial_print_banner(const char *service, const char *version);
void rodial_log_info(const char *msg);
void rodial_log_error(const char *msg, int fatal);
void rodial_log_debug(const char *msg);
int  rodial_file_exists(const char *path);

#ifdef __cplusplus
}
#endif

#endif // RODIAL_CLI_HELPER_H
/*
 * Rodial Utils - CLI Helper (C Core)
 * Provides lightweight CLI parsing and logging used across
 * Rodial OS daemons and tools.
 */