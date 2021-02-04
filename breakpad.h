#pragma once

#if __cplusplus
extern "C"
{
#endif

    extern void breakpad_init(const char *str);
    extern void onDumpCallback(void);

#if __cplusplus
}
#endif
