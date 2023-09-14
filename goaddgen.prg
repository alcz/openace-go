/*
    goaddgen.prg    -- spits out Go wrappers to SWIG/CGo imports
                       when ran, will make oerr/*, onoerr/*, oauto/*
                       packages

    license is MIT, see LICENSE

    Copyright (c) 2023 Aleksander Czajczynski
*/


#include "hbclass.ch"

#define _PKG_NOERR 1
#define _PKG_ERR   2
#define _PKG_AUTO  3

#pragma -b +

PROCEDURE Main
   LOCAL aLines, aFuncs, c, a, n, nSpec, cSpec, hFunc

   aLines := hb_atokens( StrTran( hb_memoRead( "goaddgen.txt" ), hb_bchar( 13 ) ), ;
                                                                 hb_bchar( 10 ) )
   IF Len( aLines ) > 0
      aFuncs := Array( Len( aLines ) )
   ENDIF

   FOR EACH c IN aLines
      c := AllTrim( c )
      IF ! Empty( c )
         IF At( "//", c ) == 0 
            OutErr( "line", c:__enumIndex, "no return specifier // -", c )
            LOOP
         ENDIF

         IF ( n := At( "(", c ) ) == 0 
            OutErr( "line", c:__enumIndex, "invalid func definition, no ( found -", c )
            LOOP
         ENDIF

         hFunc := { => }
         hFunc["name"] := Left( c, n - 1 )
         c := SubStr( c, n )

         IF ( n := At( ")", c ) ) == 0 
            OutErr( "line", c:__enumIndex, "invalid func definition, no ) found -", c )
            LOOP
         ENDIF

         nSpec := At( "//", c ) + 2
         cSpec := LTrim( SubStr( c, nSpec ) )
         IF At( ":", cSpec ) > 0
            a := hb_atokens( cSpec, ":" )
            hFunc["retspec"] := a[ 1 ]
            IF a[ 1 ] == "string"
               hFunc["retvar"] := a[ 2 ]
               hFunc["retopt"] := ""
            ELSE
               hFunc["retvar"] := ""
               hFunc["retopt"] := a[ 2 ]
            ENDIF
         ELSE
            hFunc["retspec"] := cSpec
            hFunc["retvar"] := ""
            hFunc["retopt"] := ""
         ENDIF

         hFunc["args_orig"] := Left( c, n )
         hFunc["args_rewritten"] := RewriteArgs( hFunc )
         aFuncs[ c:__enumIndex ]  := hFunc

      ENDIF
   NEXT

   Gen( _PKG_NOERR, aFuncs, "onoerr", "acenoerr.go" )
   Gen( _PKG_ERR,   aFuncs, "oerr",   "acewerr.go" )
   Gen( _PKG_AUTO,  aFuncs, "oauto",  "aceauto.go" )
   
   RETURN

PROCEDURE Gen( nPkgType, aFuncs, cPkg, cOutput )
   LOCAL cOut, oWrt := GoWriter():New()
   LOCAL h, oAuto := GoWriter():New()
   LOCAL bPkgAuto := { |o,cName|
      Header( o )
      o:WriteLine("package " + cName)
      o:WriteLine([])
      o:WriteLine("/*")
      o:WriteLine("#include <stdlib.h>")
      o:WriteLine("*/")
      o:WriteLine([import "C"])
      o:WriteLine([import openace "github.com/alcz/openace-go"])
      o:WriteLine([import "errors"])
      o:WriteLine([import "fmt"])
      o:WriteLine([]) 
   }
   LOCAL bPkgWrap := { |o,cName|
      Header( o )
      o:WriteLine("package " + cName)
      o:WriteLine([])
      o:WriteLine([import "github.com/alcz/openace-go/oauto"])
      o:WriteLine([import openace "github.com/alcz/openace-go"])
   }

   SWITCH nPkgType
      CASE _PKG_NOERR
         Eval( bPkgWrap, oWrt, cPkg )
         EXIT
      CASE _PKG_ERR
         Eval( bPkgWrap, oWrt, cPkg )
         EXIT
      CASE _PKG_AUTO
         Eval( bPkgAuto, oAuto, cPkg )
   ENDSWITCH

   FOR EACH h IN aFuncs
      IF h == NIL
         EXIT
      ENDIF

      IF nPkgType == _PKG_AUTO .AND. ; /* execute AutoGen() once */
         h["retspec"] == "string"
         AutoGen( oAuto, h )
      ENDIF

      IF h["retspec"] == "todo"
         LOOP
      ENDIF

      SWITCH nPkgType
         CASE _PKG_NOERR 
            oWrt:WriteLine( "" )
            oWrt:PushBlock( "func " + h["name"] + h["args_rewritten"] + "(_ret " + RetSpec( h ) + ")" )
            SWITCH h["retspec"] 
               CASE "string"
                  oWrt:WriteLine( "ret, _ := oauto." + h["name"] + "Err" + h["call_args"] )
                  oWrt:WriteLine( "return ret" )
                  EXIT
               CASE "logic"
                  oWrt:WriteLine( "err := openace." + h["name"] + h["call_args"] )
                  oWrt:PushBlock( "if err > 0" )
                  IF h["retopt"] == "erristrue"
                     oWrt:WriteLine("return true /* !!! error reported as true here */ " )
                  ELSE
                     oWrt:WriteLine("return false")
                  ENDIF
                  oWrt:PopBlock()
                  oWrt:WriteLine("return true")
                  EXIT
              CASE "logicu16last"
                 IF h["retopt"] == "erristrue"
                    oWrt:WriteLine( "var ret uint16 = 1" )
                 ELSE
                    oWrt:WriteLine( "var ret uint16 = 0" )
                 ENDIF
                 oWrt:PushBlock( "if 0 == openace." + h["name"] + StrTran( h["call_args_int"], "&last", "&ret" ) )
                 oWrt:WriteLine( "return ret != 0" )
                 oWrt:PopBlock()
                 IF h["retopt"] == "erristrue"
                    oWrt:WriteLine( "return true" )
                 ELSE
                    oWrt:WriteLine( "return false" )
                 ENDIF
 //              oWrt:WriteLine( "ret, _ := openace." + h["name"] + h["call_args"] )
                 EXIT
              CASE "uint16last"
              CASE "uintlast"
              CASE "intlast"
              CASE "float64last"
                 oWrt:WriteLine( "var ret " + RetSpec( h ) +" = 0" )
                 oWrt:WriteLine( /*"_ := */ "openace." + h["name"] + StrTran( h["call_args_int"], "&last", "&ret" ) )
                 oWrt:WriteLine( "return ret" )
            ENDSWITCH
            oWrt:PopBlock()
            EXIT
         CASE _PKG_ERR
            oWrt:WriteLine( "" )
            oWrt:PushBlock( "func " + h["name"] + h["args_rewritten"] + "(_ret " + RetSpec( h ) + ", _ret_err uint)" )
            SWITCH h["retspec"] 
               CASE "string"
                  oWrt:WriteLine( "ret, err := oauto." + h["name"] + "Err" + h["call_args"] )
                  oWrt:WriteLine( "return ret, err" )
                  EXIT
               CASE "logic"
                  oWrt:WriteLine( "var ret bool" )
                  oWrt:WriteLine( "err := openace." + h["name"] + h["call_args"] )
                  // oWrt:PushBlock( "if err > 0" )
                  IF h["retopt"] == "erristrue"
                     oWrt:WriteLine("ret = err > 0 /* !!! error reported as true here */ " )
                  ELSE
                     oWrt:WriteLine("ret = err == 0")
                  ENDIF
 //                 oWrt:PopBlock()
                  oWrt:WriteLine("return ret, err")
                  EXIT
              CASE "logicu16last"
//                 oWrt:WriteLine( "var ret bool" )
                 IF h["retopt"] == "erristrue"
                    oWrt:WriteLine( "var ret uint16 = 1" )
                 ELSE
                    oWrt:WriteLine( "var ret uint16 = 0" )
                 ENDIF
                 oWrt:WriteLine( "err := openace." + h["name"] + StrTran( h["call_args_int"], "&last", "&ret" ) )
                 IF h["retopt"] == "erristrue"
                    oWrt:PushBlock( "if err > 0" )
                    oWrt:WriteLine( "return true, err" )
                    oWrt:PopBlock()
                 ENDIF
                 oWrt:WriteLine( "return ret != 0, err" )
 //              oWrt:WriteLine( "ret, _ := openace." + h["name"] + h["call_args"] )
                 EXIT
              CASE "uint16last"
              CASE "uintlast"
              CASE "intlast"
              CASE "float64last"
                 oWrt:WriteLine( "var ret " + RetSpec( h ) +" = 0" )
                 oWrt:WriteLine( "err := openace." + h["name"] + StrTran( h["call_args_int"], "&last", "&ret" ) )
                 oWrt:WriteLine( "return ret, err" )
            ENDSWITCH
            oWrt:PopBlock()
            EXIT
         CASE _PKG_AUTO
            oAuto:WriteLine( "" )
            IF h["retspec"]  == "logic" /* errors makes a function with retval nosense here */
               oAuto:PushBlock( "func " + h["name"] + h["args_rewritten"] )
            ELSE
               oAuto:PushBlock( "func " + h["name"] + h["args_rewritten"] + "(_ret " + RetSpec( h ) + ")" )
            ENDIF
            SWITCH h["retspec"] 
               CASE "string"
                  oAuto:WriteLine( "ret, err := " + h["name"] + "Err" + h["call_args"] )
                  oAuto:PushBlock( "if err > 0" )
                  oAuto:WriteLine( [panic( errors.New( fmt.Sprintf("] + h["name"] + [() error %d", err) ) )] )
                  oAuto:PopBlock()
                  oAuto:WriteLine( "return ret" )
                  EXIT
               CASE "logic"
                  oAuto:WriteLine( "err := openace." + h["name"] + h["call_args"] )
                  oAuto:PushBlock( "if err > 0" )
                  oAuto:WriteLine( [panic( errors.New( fmt.Sprintf("] + h["name"] + [() error %d", err) ) )] )
                  oAuto:PopBlock()
                  EXIT
              CASE "logicu16last"
                 IF h["retopt"] == "erristrue"
                    oAuto:WriteLine( "var ret uint16 = 1" )
                 ELSE
                    oAuto:WriteLine( "var ret uint16 = 0" )
                 ENDIF
                 oAuto:WriteLine( "err := openace." + h["name"] + StrTran( h["call_args_int"], "&last", "&ret" ) )
                 oAuto:PushBlock( "if err > 0" )
                 oAuto:WriteLine( [panic( errors.New( fmt.Sprintf("] + h["name"] + [() error %d", err) ) )] )
                 oAuto:PopBlock()
                 oAuto:WriteLine( "return ret != 0" )
                 EXIT
              CASE "uint16last"
              CASE "uintlast"
              CASE "intlast"
              CASE "float64last"
                 oAuto:WriteLine( "var ret " + RetSpec( h ) +" = 0" )
                 oAuto:WriteLine( "err := openace." + h["name"] + StrTran( h["call_args_int"], "&last", "&ret" ) )
                 oAuto:PushBlock( "if err > 0" )
                 oAuto:WriteLine( [panic( errors.New( fmt.Sprintf("] + h["name"] + [() error %d", err) ) )] )
                 oAuto:PopBlock()
                 oAuto:WriteLine( "return ret" )
            ENDSWITCH
            oAuto:PopBlock()
            EXIT
      ENDSWITCH
   NEXT

   IF nPkgType == _PKG_AUTO
      hb_memoWrit( cPkg + hb_ps() + cOutput, oAuto:Buf )
   ELSE
      hb_memoWrit( cPkg + hb_ps() + cOutput, oWrt:Buf )
   ENDIF

   RETURN

PROCEDURE RewriteArgs( hFunc )
   LOCAL c, a, nDel, lProc := .F., aCallArgs, aCallArgsInt

   a := hb_atokens( AllTrim( SubStr( hb_strShrink( AllTrim( hFunc["args_orig"] ) ), 2 ) ), "," )

   aCallArgs := AClone( a )
   AEval( aCallArgs, { |x,n| aCallArgs[ n ] := hb_atokens( LTrim( x ), " " )[ 1 ] } )
   aCallArgsInt := AClone( aCallArgs )

   IF ( nDel := AScan( a, { |x| Left( AllTrim( x ), 8 ) == "len uint" } ) ) > 0
      IF Right( c := RTrim( a[ nDel - 1 ] ), 7 ) == " string"
         c := hb_strShrink( c, 7 )
         IF Right( RTrim( a[ nDel ] ), 2 ) == "16"
            aCallArgs[ nDel ] := "uint16( len(" + c + " ) )"
         ELSE
            aCallArgs[ nDel ] := "uint( len(" + c + " ) )"
         ENDIF
         aCallArgsInt[ nDel ] := aCallArgs[ nDel ]
         HB_ADel( a, nDel, .T. )
         lProc := .T.
      ENDIF
   ENDIF

   nDel := NIL

   IF Right( hFunc["retspec"], 4 ) == "last"
      ASize( a, Len( a ) - 1 )
      aCallArgsInt[ Len( aCallArgsInt ) ] := "&last"
      lProc := .T.
   ELSEIF hFunc["retspec"] == "string"
      AEval( a, { |x,n| IIF( Left( LTrim( x ), Len( hFunc["retvar"] ) + 1 ) == ;
                                      hFunc["retvar"] + " ", nDel := n, NIL ) } )
      IF nDel # NIL
         HB_ADel( a, nDel, .T. )
         HB_ADel( a, nDel, .T. ) /* after string param there is linked length, also deleted */
         HB_ADel( aCallArgs, nDel, .T. )
         HB_ADel( aCallArgs, nDel, .T. )
         aCallArgsInt[ nDel     ] := "buf" /* currently hardcoded in AutoGen() proc */
         aCallArgsInt[ nDel + 1 ] := "&buflen"
         lProc := .T.
      ENDIF
   ENDIF

   c := "( "
   AEval( aCallArgs, { |x,n| c += x + IIF( n < Len( aCallArgs ), ", ", "" ) } )
   c += " )"
   hFunc["call_args"] := c

   c := "( "
   AEval( aCallArgsInt, { |x,n| c += x + IIF( n < Len( aCallArgsInt ), ", ", "" ) } )
   c += " )"
   hFunc["call_args_int"] := c

   IF lProc
      c := "( "
      AEval( a, { |x,n| c += x + IIF( n < Len( a ), ",", "" ) } )
      c += " )"
   ELSE
      c := hFunc["args_orig"]
   ENDIF

   RETURN c

FUNCTION RetSpec( hFunc )
   LOCAL cIn := hFunc["retspec"]
   IF Left( cIn, 5 ) == "logic"
      RETURN "bool"
   ELSEIF Right( cIn, 4 ) == "last"
      RETURN hb_strShrink( cIn, 4 )
   ENDIF
   RETURN cIn
   
CLASS TrivialWriter
   PROTECTED:
   VAR nIndent INIT 0
   VAR cEoL INIT HB_EoL()
   EXPORTED:
   METHOD WriteLine( c )
   METHOD WriteIndented( c )
   METHOD PushBlock( c )
   METHOD PopBlock( c )
ENDCLASS

METHOD WriteLine( c ) CLASS TrivialWriter
   RETURN ::WriteIndented( c )

METHOD WriteIndented( c ) CLASS TrivialWriter
   RETURN ::Out( Space( ::nIndent ) + c + ::cEoL )

METHOD PushBlock( cHeader ) CLASS TrivialWriter
   IF !Empty( cHeader )
      ::WriteIndented( cHeader )
   ENDIF
   ::WriteIndented("{")
   ::nIndent += 3
   RETURN self

METHOD PopBlock() CLASS TrivialWriter
   ::nIndent -= 3
   ::WriteIndented("}")
   RETURN self

CLASS VarWriter INHERIT TrivialWriter
   PROTECTED:
   VAR cBuf INIT ""
   EXPORTED:
   METHOD Out( c ) INLINE ::cBuf += c, Len( c ) 
   METHOD OutEoL() INLINE ::cBuf += ::cEol, Len( ::cEoL )
   ACCESS Buf INLINE ::cBuf
ENDCLASS

CLASS GoWriter INHERIT VarWriter
   EXPORTED:
   METHOD PushBlock( c )
ENDCLASS

METHOD PushBlock( c ) CLASS GoWriter
   IF !Empty( c )
      ::WriteIndented( c + " {" )
   ELSE
      ::WriteIndented("{")
   ENDIF
   ::nIndent += 3
   RETURN self

PROCEDURE AutoGen( o, h )
   o:WriteLine([])
   o:PushBlock( "func " + h["name"] + "Err" + h["args_rewritten"] + "(_buf_ret string, _err_ret uint)" )
   o:WriteLine([var buf string])
   IF At("len *uint16", h["args_orig"] ) > 0
      o:WriteLine([var buflen uint16])
   ELSE
      o:WriteLine([var buflen uint])
   ENDIF
   o:WriteLine([])
   o:WriteLine([ret := openace.] + h["name"] + StrTran( h["call_args_int"], " buf,", " 0," ) )
   o:WriteLine([])
   o:WriteLine("/* intentional error path */")
   o:PushBlock("if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0")
   o:WriteLine(   "bufptr := C.malloc( C.uint( buflen + 1 ) )")
   o:WriteLine([])
   o:WriteLine(   [ret := openace.] + h["name"] + StrTran( h["call_args_int"], " buf,", " (uintptr)( bufptr )," ) )
   o:WriteLine([])
   o:WriteLine(   [buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )] )  
   o:WriteLine(   [C.free( bufptr )] )
   o:WriteLine(   [return buf, ret] )
   o:PopBlock()
   o:WriteLine([return "", ret]) 
   o:PopBlock()

   RETURN

#define GENINFO  GENINFO1 + HB_EoL() + GENINFO2 + HB_EoL() + GENINFO3 + HB_EoL() + GENINFO4 + HB_EoL()
#define GENINFO1 "/* This file is automatically generated by goaddgen.prg originated from"
#define GENINFO2 "   https://github.com/alcz/openace-go"
#define GENINFO3 "   goaddgen.txt was used as input containing simple suggestions"
#define GENINFO4 "   on how to forward to plain C-API/CGo/SWIG */"

STATIC PROCEDURE Header( o )
   o:WriteLine( GENINFO )
   RETURN
