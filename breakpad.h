#pragma once

#if __cplusplus
extern "C"
{
#endif

    extern void breakpad_init(const char *str);
    extern void onDumpCallback();

#if __cplusplus
}
#endif
