package main

const program_app_c = `#include <stdio.h>

int main(int argc, char *argv[])
{
    printf("Hello World\n");
    
    return 0;
}
`

const program_app_cpp = `#include <iostream>

using std::cout;
using std::endl;

int main(int argc, char *argv[])
{
    cout << "Hello World" << endl;
    
    return 0;
}
`

const program_header = `#ifndef {{.Program}}_H
#define {{.Program}}_H

#ifndef __cplusplus
    #include <stdbool.h>
#endif

#ifdef __cplusplus
extern "C" {
#endif

bool is_even(int n);

#ifdef __cplusplus
}
#endif

#endif  // {{.Program}}_H
`

const program_lib_c = `#include <stdbool.h>
#include "{{.Program}}.h"

bool is_even(int n)
{
    return n % 2 == 0;
}
`

const program_lib_cpp = `#include "{{.Program}}.hpp"

bool is_even(int n)
{
    return n % 2 == 0;
}
`

const programLibDef = `LIBRARY {{.Program}}
EXPORTS
    is_even
`
