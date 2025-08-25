#include <unistd.h>
#include <errno.h>
#include <sys/lock.h>
#include <stdio.h>
#include <sys/cdefs.h>

extern char _heap_start[];
extern char _heap_end[];

static char *__brk = _heap_start;

void *sbrk(ptrdiff_t incr)
{
        if (incr < 0)
        {
                if ((size_t)(__brk - _heap_start) < (size_t)(-incr))
                {
                        errno = ENOMEM;
                        return (void *)-1;
                }
        }
        else
        {
                if ((size_t)(_heap_end - __brk) < (size_t)incr)
                {
                        errno = ENOMEM;
                        return (void *)-1;
                }
        }
        void *ret = __brk;
        __brk += incr;
        return ret;
}

struct __lock
{
        char unused;
};

struct __lock __lock___libc_recursive_mutex;

void __retarget_lock_init(_LOCK_T *lock)
{
        (void)lock;
}

void __retarget_lock_init_recursive(_LOCK_T *lock)
{
        (void)lock;
}

void __retarget_lock_close(_LOCK_T lock)
{
        (void)lock;
}

void __retarget_lock_close_recursive(_LOCK_T lock)
{
        (void)lock;
}

void __retarget_lock_acquire(_LOCK_T lock)
{
        (void)lock;
}

void __retarget_lock_acquire_recursive(_LOCK_T lock)
{
        (void)lock;
}

void __retarget_lock_release(_LOCK_T lock)
{
        (void)lock;
}

void __retarget_lock_release_recursive(_LOCK_T lock)
{
        (void)lock;
}

// Defined in the runtime package. Writes to the default console (usually, the
// first UART or an USB-CDC device).
int runtime_putchar(char c, FILE *file);

// Define stdin, stdout, and stderr as a single object.
// This object must not reside in ROM.
static FILE __stdio = FDEV_SETUP_STREAM(runtime_putchar, NULL, NULL, _FDEV_SETUP_WRITE);

int runtime_putchar(char c, FILE *file)
{
        (void)file;
        extern void writeByte(char);
        writeByte(c);
        return 0;
}

// Define the underlying structs for stdin, stdout, and stderr.
FILE *const stdin = &__stdio;
__strong_reference(stdin, stdout);
__strong_reference(stdin, stderr);
