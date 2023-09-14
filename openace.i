%module openace

typedef unsigned int HB_U32;
typedef signed int HB_I32;
typedef unsigned short HB_USHORT;
typedef char HB_UCHAR;
typedef int HB_BOOL;
typedef void ACE_UCHAR_OUT;

#define HB_EXPORT

%rename("%(strip:[Zwa])s") "";

%{
#include <hbdefs.h>
#include <hbapi.h>
#include <hbapirdd.h>

#include "include/openace.h"
%}

%include "include/openace.h"
