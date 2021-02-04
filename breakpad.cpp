#include "breakpad.h"
#include <client/linux/handler/exception_handler.h>

static google_breakpad::ExceptionHandler *eh = nullptr;

static bool dumpCallback(const google_breakpad::MinidumpDescriptor &descriptor, void *context, bool succeeded)
{
    printf("Dump path: %s\n", descriptor.path());
    onDumpCallback();
    return succeeded;
}

void breakpad_init(const char *outpath)
{
    google_breakpad::MinidumpDescriptor descriptor(outpath);
    eh = new google_breakpad::ExceptionHandler(descriptor, NULL, dumpCallback, NULL, true, -1);
}
