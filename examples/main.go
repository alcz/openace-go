package main

import (
	"fmt"
	"github.com/alcz/openace-go"
	"github.com/alcz/openace-go/onoerr"
	"github.com/alcz/openace-go/oerr"
	"github.com/alcz/openace-go/oauto"
)

func job( workarea uint ) {

   for {
      if oauto.AtEOF( workarea ) {
         break
      }

      fmt.Println( oauto.GetString( workarea, "LAST", openace.RTRIM ) )

      oauto.Skip( workarea, 1 )
   }
}

func main() {
        var conn_handle uint
        if 0 == openace.Connect( "dbfcdx://", &conn_handle ) {
//           openace.ZwaMgGetServerType( 0, &server_type )        
           var test_db uint
           ret := openace.OpenTable( conn_handle, 
                              "test.dbf",
                              "TEST",
                              openace.CDX,
                              0,
                              0,
                              0,
                              openace.SHARED,
                              &test_db )
           if ret == 0 {
              fmt.Println("hello job" )
           }
           job( test_db )

           fmt.Println("hello %d opendb=%d test_db=%d", conn_handle, ret, test_db )

           /* demonstrating differences between return handling (oerr, onoerr, oauto) */

           onoerr.CloseTable( test_db )

           _, errcode := oerr.CloseTable( test_db ) 

           fmt.Println("double close error -> %d", errcode )

           oauto.CloseTable( test_db ) // yes, program should panic() here
        }
}
