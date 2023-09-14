/* This file is automatically generated by goaddgen.prg originated from
   https://github.com/alcz/openace-go
   goaddgen.txt was used as input containing simple suggestions
   on how to forward to plain C-API/CGo/SWIG */

package oauto

/*
#include <stdlib.h>
*/
import "C"
import openace "github.com/alcz/openace-go"
import "errors"
import "fmt"


func SetCallbackOnRTE( fp uintptr ) {
   err := openace.SetCallbackOnRTE( fp )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetCallbackOnRTE() error %d", err) ) )
   }
}

func SetCallbackTrace( fp uintptr ) {
   err := openace.SetCallbackTrace( fp )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetCallbackTrace() error %d", err) ) )
   }
}

func AddCustomKey( index uint ) {
   err := openace.AddCustomKey( index )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("AddCustomKey() error %d", err) ) )
   }
}

func AppendRecord( wa uint ) {
   err := openace.AppendRecord( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("AppendRecord() error %d", err) ) )
   }
}

func ApplicationExit() {
   err := openace.ApplicationExit(  )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("ApplicationExit() error %d", err) ) )
   }
}

func AtBOF( wa uint )(_ret bool) {
   var ret uint16 = 1
   err := openace.AtBOF( wa, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("AtBOF() error %d", err) ) )
   }
   return ret != 0
}

func AtEOF( wa uint )(_ret bool) {
   var ret uint16 = 1
   err := openace.AtEOF( wa, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("AtEOF() error %d", err) ) )
   }
   return ret != 0
}

func CacheRecords( wa uint, num_recs uint16 ) {
   err := openace.CacheRecords( wa, num_recs )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("CacheRecords() error %d", err) ) )
   }
}

func BeginTransaction( conn uint ) {
   err := openace.BeginTransaction( conn )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("BeginTransaction() error %d", err) ) )
   }
}

func ClearFilter( wa uint ) {
   err := openace.ClearFilter( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("ClearFilter() error %d", err) ) )
   }
}

func ClearProgressCallback() {
   err := openace.ClearProgressCallback(  )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("ClearProgressCallback() error %d", err) ) )
   }
}

func ClearRelation( master_wa uint ) {
   err := openace.ClearRelation( master_wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("ClearRelation() error %d", err) ) )
   }
}

func ClearScope( index uint, opt uint16 ) {
   err := openace.ClearScope( index, opt )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("ClearScope() error %d", err) ) )
   }
}

func CloseAllIndexes( wa uint ) {
   err := openace.CloseAllIndexes( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("CloseAllIndexes() error %d", err) ) )
   }
}

func CloseIndex( index uint ) {
   err := openace.CloseIndex( index )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("CloseIndex() error %d", err) ) )
   }
}

func CloseTable( wa uint ) {
   err := openace.CloseTable( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("CloseTable() error %d", err) ) )
   }
}

func CommitTransaction( conn uint ) {
   err := openace.CommitTransaction( conn )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("CommitTransaction() error %d", err) ) )
   }
}

func Connect( server_name string )(_ret uint) {
   var ret uint = 0
   err := openace.Connect( server_name, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("Connect() error %d", err) ) )
   }
   return ret
}

func CopyTableContents( handle uint, wa_to uint, filter_opt uint16 ) {
   err := openace.CopyTableContents( handle, wa_to, filter_opt )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("CopyTableContents() error %d", err) ) )
   }
}

func CreateIndex( handle uint, filename string, tag string, expr string, cond string, while string, opt uint )(_ret uint) {
   var ret uint = 0
   err := openace.CreateIndex( handle, filename, tag, expr, cond, while, opt, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("CreateIndex() error %d", err) ) )
   }
   return ret
}

func CreateTable( conn uint, name string, alias string, table_type uint16, char_type uint16, lock_type uint16, check_rights uint16, memo_size uint16, fields string )(_ret uint) {
   var ret uint = 0
   err := openace.CreateTable( conn, name, alias, table_type, char_type, lock_type, check_rights, memo_size, fields, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("CreateTable() error %d", err) ) )
   }
   return ret
}

func DeleteCustomKey( index uint ) {
   err := openace.DeleteCustomKey( index )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("DeleteCustomKey() error %d", err) ) )
   }
}

func DeleteIndex( index uint ) {
   err := openace.DeleteIndex( index )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("DeleteIndex() error %d", err) ) )
   }
}

func DeleteRecord( wa uint ) {
   err := openace.DeleteRecord( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("DeleteRecord() error %d", err) ) )
   }
}

func Disconnect( conn uint ) {
   err := openace.Disconnect( conn )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("Disconnect() error %d", err) ) )
   }
}

func ExtractKeyErr( index uint )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.ExtractKey( index, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.ExtractKey( index, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func ExtractKey( index uint )(_ret string) {
   ret, err := ExtractKeyErr( index )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("ExtractKey() error %d", err) ) )
   }
   return ret
}

func GetBinaryErr( wa uint, fname string, offset uint )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint
   
   ret := openace.GetBinary( wa, fname, offset, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetBinary( wa, fname, offset, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetBinary( wa uint, fname string, offset uint )(_ret string) {
   ret, err := GetBinaryErr( wa, fname, offset )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetBinary() error %d", err) ) )
   }
   return ret
}

func GetBinaryLength( wa uint, fname string )(_ret uint) {
   var ret uint = 0
   err := openace.GetBinaryLength( wa, fname, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetBinaryLength() error %d", err) ) )
   }
   return ret
}

func GetDateFormatErr(  )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetDateFormat( 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetDateFormat( (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetDateFormat(  )(_ret string) {
   ret, err := GetDateFormatErr(  )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetDateFormat() error %d", err) ) )
   }
   return ret
}

func GetDouble( wa uint, fname string )(_ret float64) {
   var ret float64 = 0
   err := openace.GetDouble( wa, fname, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetDouble() error %d", err) ) )
   }
   return ret
}

func GetErrorStringErr( errcode uint )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetErrorString( errcode, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetErrorString( errcode, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetErrorString( errcode uint )(_ret string) {
   ret, err := GetErrorStringErr( errcode )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetErrorString() error %d", err) ) )
   }
   return ret
}

func GetFieldErr( wa uint, fname string, opt uint16 )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint
   
   ret := openace.GetField( wa, fname, 0, &buflen, opt )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetField( wa, fname, (uintptr)( bufptr ), &buflen, opt )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetField( wa uint, fname string, opt uint16 )(_ret string) {
   ret, err := GetFieldErr( wa, fname, opt )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetField() error %d", err) ) )
   }
   return ret
}

func GetFieldDecimals( wa uint, fname string )(_ret uint16) {
   var ret uint16 = 0
   err := openace.GetFieldDecimals( wa, fname, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetFieldDecimals() error %d", err) ) )
   }
   return ret
}

func GetFieldLength( wa uint, fname string )(_ret uint) {
   var ret uint = 0
   err := openace.GetFieldLength( wa, fname, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetFieldLength() error %d", err) ) )
   }
   return ret
}

func GetFieldNameErr( wa uint, pos uint16 )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetFieldName( wa, pos, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetFieldName( wa, pos, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetFieldName( wa uint, pos uint16 )(_ret string) {
   ret, err := GetFieldNameErr( wa, pos )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetFieldName() error %d", err) ) )
   }
   return ret
}

func GetFieldType( wa uint, fname string )(_ret uint16) {
   var ret uint16 = 0
   err := openace.GetFieldType( wa, fname, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetFieldType() error %d", err) ) )
   }
   return ret
}

func GetHandleType( handle uint )(_ret uint16) {
   var ret uint16 = 0
   err := openace.GetHandleType( handle, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetHandleType() error %d", err) ) )
   }
   return ret
}

func GetIndexConditionErr( index uint )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetIndexCondition( index, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetIndexCondition( index, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetIndexCondition( index uint )(_ret string) {
   ret, err := GetIndexConditionErr( index )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetIndexCondition() error %d", err) ) )
   }
   return ret
}

func GetIndexExprErr( index uint )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetIndexExpr( index, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetIndexExpr( index, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetIndexExpr( index uint )(_ret string) {
   ret, err := GetIndexExprErr( index )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetIndexExpr() error %d", err) ) )
   }
   return ret
}

func GetIndexFilenameErr( index uint, opt uint16 )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetIndexFilename( index, opt, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetIndexFilename( index, opt, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetIndexFilename( index uint, opt uint16 )(_ret string) {
   ret, err := GetIndexFilenameErr( index, opt )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetIndexFilename() error %d", err) ) )
   }
   return ret
}

func GetIndexNameErr( index uint )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetIndexName( index, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetIndexName( index, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetIndexName( index uint )(_ret string) {
   ret, err := GetIndexNameErr( index )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetIndexName() error %d", err) ) )
   }
   return ret
}

func GetIndexHandle( wa uint, tag string )(_ret uint) {
   var ret uint = 0
   err := openace.GetIndexHandle( wa, tag, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetIndexHandle() error %d", err) ) )
   }
   return ret
}

func GetIndexHandleByOrder( wa uint, ordnum uint16 )(_ret uint) {
   var ret uint = 0
   err := openace.GetIndexHandleByOrder( wa, ordnum, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetIndexHandleByOrder() error %d", err) ) )
   }
   return ret
}

func GetJulian( wa uint, fname string )(_ret int) {
   var ret int = 0
   err := openace.GetJulian( wa, fname, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetJulian() error %d", err) ) )
   }
   return ret
}

func GetKeyCount( index uint, filter_opt uint16 )(_ret uint) {
   var ret uint = 0
   err := openace.GetKeyCount( index, filter_opt, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetKeyCount() error %d", err) ) )
   }
   return ret
}

func GetKeyLength( index uint )(_ret uint16) {
   var ret uint16 = 0
   err := openace.GetKeyLength( index, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetKeyLength() error %d", err) ) )
   }
   return ret
}

func GetKeyNum( index uint, filter_opt uint16 )(_ret uint) {
   var ret uint = 0
   err := openace.GetKeyNum( index, filter_opt, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetKeyNum() error %d", err) ) )
   }
   return ret
}

func GetKeyType( index uint )(_ret uint16) {
   var ret uint16 = 0
   err := openace.GetKeyType( index, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetKeyType() error %d", err) ) )
   }
   return ret
}

func GetLastErrorErr( errcode *uint )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetLastError( errcode, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetLastError( errcode, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetLastError( errcode *uint )(_ret string) {
   ret, err := GetLastErrorErr( errcode )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetLastError() error %d", err) ) )
   }
   return ret
}

func GetLastTableUpdateErr( wa uint )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetLastTableUpdate( wa, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetLastTableUpdate( wa, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetLastTableUpdate( wa uint )(_ret string) {
   ret, err := GetLastTableUpdateErr( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetLastTableUpdate() error %d", err) ) )
   }
   return ret
}

func GetLogical( wa uint, fname string )(_ret bool) {
   var ret uint16 = 0
   err := openace.GetLogical( wa, fname, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetLogical() error %d", err) ) )
   }
   return ret != 0
}

func GetMemoLength( wa uint, fname string )(_ret uint) {
   var ret uint = 0
   err := openace.GetMemoLength( wa, fname, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetMemoLength() error %d", err) ) )
   }
   return ret
}

func GetMilliseconds( wa uint, fname string )(_ret int) {
   var ret int = 0
   err := openace.GetMilliseconds( wa, fname, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetMilliseconds() error %d", err) ) )
   }
   return ret
}

func GetNumFields( wa uint )(_ret uint16) {
   var ret uint16 = 0
   err := openace.GetNumFields( wa, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetNumFields() error %d", err) ) )
   }
   return ret
}

func GetNumIndexes( wa uint )(_ret uint16) {
   var ret uint16 = 0
   err := openace.GetNumIndexes( wa, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetNumIndexes() error %d", err) ) )
   }
   return ret
}

func GetNumLocks( wa uint )(_ret uint16) {
   var ret uint16 = 0
   err := openace.GetNumLocks( wa, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetNumLocks() error %d", err) ) )
   }
   return ret
}

func GetRecordCount( wa uint, filter_opt uint16 )(_ret uint) {
   var ret uint = 0
   err := openace.GetRecordCount( wa, filter_opt, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetRecordCount() error %d", err) ) )
   }
   return ret
}

func GetRecordNum( wa uint, filter_opt uint16 )(_ret uint) {
   var ret uint = 0
   err := openace.GetRecordNum( wa, filter_opt, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetRecordNum() error %d", err) ) )
   }
   return ret
}

func GetRecordLength( wa uint )(_ret uint) {
   var ret uint = 0
   err := openace.GetRecordLength( wa, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetRecordLength() error %d", err) ) )
   }
   return ret
}

func GetRelKeyPos( wa uint )(_ret float64) {
   var ret float64 = 0
   err := openace.GetRelKeyPos( wa, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetRelKeyPos() error %d", err) ) )
   }
   return ret
}

func GetServerNameErr( conn uint )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetServerName( conn, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetServerName( conn, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetServerName( conn uint )(_ret string) {
   ret, err := GetServerNameErr( conn )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetServerName() error %d", err) ) )
   }
   return ret
}

func GetStringErr( wa uint, fname string, opt uint16 )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint
   
   ret := openace.GetString( wa, fname, 0, &buflen, opt )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetString( wa, fname, (uintptr)( bufptr ), &buflen, opt )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetString( wa uint, fname string, opt uint16 )(_ret string) {
   ret, err := GetStringErr( wa, fname, opt )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetString() error %d", err) ) )
   }
   return ret
}

func GetTableAliasErr( wa uint )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetTableAlias( wa, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetTableAlias( wa, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetTableAlias( wa uint )(_ret string) {
   ret, err := GetTableAliasErr( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetTableAlias() error %d", err) ) )
   }
   return ret
}

func GetTableFilenameErr( wa uint, opt uint16 )(_buf_ret string, _err_ret uint) {
   var buf string
   var buflen uint16
   
   ret := openace.GetTableFilename( wa, opt, 0, &buflen )
   
   /* intentional error path */
   if ret == openace.ECA_INSUFFICIENT_BUFFER && buflen > 0 {
      bufptr := C.malloc( C.uint( buflen + 1 ) )
      
      ret := openace.GetTableFilename( wa, opt, (uintptr)( bufptr ), &buflen )
      
      buf = C.GoStringN( (*C.char)( bufptr ), C.int( buflen ) )
      C.free( bufptr )
      return buf, ret
   }
   return "", ret
}

func GetTableFilename( wa uint, opt uint16 )(_ret string) {
   ret, err := GetTableFilenameErr( wa, opt )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetTableFilename() error %d", err) ) )
   }
   return ret
}

func GetTableType( wa uint )(_ret uint16) {
   var ret uint16 = 0
   err := openace.GetTableType( wa, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GetTableType() error %d", err) ) )
   }
   return ret
}

func GotoBottom( wa uint ) {
   err := openace.GotoBottom( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GotoBottom() error %d", err) ) )
   }
}

func GotoRecord( wa uint, recno uint ) {
   err := openace.GotoRecord( wa, recno )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GotoRecord() error %d", err) ) )
   }
}

func GotoTop( wa uint ) {
   err := openace.GotoTop( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("GotoTop() error %d", err) ) )
   }
}

func IsEmpty( wa uint, fname string )(_ret bool) {
   var ret uint16 = 1
   err := openace.IsEmpty( wa, fname, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("IsEmpty() error %d", err) ) )
   }
   return ret != 0
}

func IsIndexDescending( index uint )(_ret bool) {
   var ret uint16 = 0
   err := openace.IsIndexDescending( index, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("IsIndexDescending() error %d", err) ) )
   }
   return ret != 0
}

func IsIndexUnique( index uint )(_ret bool) {
   var ret uint16 = 0
   err := openace.IsIndexUnique( index, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("IsIndexUnique() error %d", err) ) )
   }
   return ret != 0
}

func IsRecordDeleted( wa uint )(_ret bool) {
   var ret uint16 = 1
   err := openace.IsRecordDeleted( wa, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("IsRecordDeleted() error %d", err) ) )
   }
   return ret != 0
}

func LockRecord( wa uint, recno uint ) {
   err := openace.LockRecord( wa, recno )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("LockRecord() error %d", err) ) )
   }
}

func LockTable( wa uint ) {
   err := openace.LockTable( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("LockTable() error %d", err) ) )
   }
}

func MgConnect( server string, username string, password string )(_ret uint) {
   var ret uint = 0
   err := openace.MgConnect( server, username, password, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("MgConnect() error %d", err) ) )
   }
   return ret
}

func MgDisconnect( handle uint ) {
   err := openace.MgDisconnect( handle )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("MgDisconnect() error %d", err) ) )
   }
}

func MgGetServerType( handle uint )(_ret uint16) {
   var ret uint16 = 0
   err := openace.MgGetServerType( handle, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("MgGetServerType() error %d", err) ) )
   }
   return ret
}

func NullTerminateStrings( null_term uint16 ) {
   err := openace.NullTerminateStrings( null_term )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("NullTerminateStrings() error %d", err) ) )
   }
}

func OpenTable( conn uint, name string, alias string, table_type uint16, char_type uint16, lock_type uint16, check_rights uint16, opts uint )(_ret uint) {
   var ret uint = 0
   err := openace.OpenTable( conn, name, alias, table_type, char_type, lock_type, check_rights, opts, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("OpenTable() error %d", err) ) )
   }
   return ret
}

func PackTable( wa uint ) {
   err := openace.PackTable( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("PackTable() error %d", err) ) )
   }
}

func RecallRecord( wa uint ) {
   err := openace.RecallRecord( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("RecallRecord() error %d", err) ) )
   }
}

func RefreshRecord( wa uint ) {
   err := openace.RefreshRecord( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("RefreshRecord() error %d", err) ) )
   }
}

func Reindex( wa uint ) {
   err := openace.Reindex( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("Reindex() error %d", err) ) )
   }
}

func RollbackTransaction( conn uint ) {
   err := openace.RollbackTransaction( conn )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("RollbackTransaction() error %d", err) ) )
   }
}

func Seek( index uint, key string, data_type uint16, seek_type uint16 )(_ret bool) {
   var ret uint16 = 0
   err := openace.Seek( index, key, uint16( len( key ) ), data_type, seek_type, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("Seek() error %d", err) ) )
   }
   return ret != 0
}

func SeekLast( index uint, key string, data_type uint16 )(_ret bool) {
   var ret uint16 = 0
   err := openace.SeekLast( index, key, uint16( len( key ) ), data_type, &ret )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SeekLast() error %d", err) ) )
   }
   return ret != 0
}

func SetDateFormat( format string ) {
   err := openace.SetDateFormat( format )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetDateFormat() error %d", err) ) )
   }
}

func SetDefault( def string ) {
   err := openace.SetDefault( def )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetDefault() error %d", err) ) )
   }
}

func SetDouble( wa uint, fname string, value float64 ) {
   err := openace.SetDouble( wa, fname, value )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetDouble() error %d", err) ) )
   }
}

func SetEmpty( wa uint, fname string ) {
   err := openace.SetEmpty( wa, fname )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetEmpty() error %d", err) ) )
   }
}

func SetField( wa uint, fname string, buf string ) {
   err := openace.SetField( wa, fname, buf, uint( len( buf ) ) )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetField() error %d", err) ) )
   }
}

func SetFilter( wa uint, filter string ) {
   err := openace.SetFilter( wa, filter )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetFilter() error %d", err) ) )
   }
}

func SetJulian( wa uint, fname string, julian int ) {
   err := openace.SetJulian( wa, fname, julian )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetJulian() error %d", err) ) )
   }
}

func SetLogical( wa uint, fname string, value uint16 ) {
   err := openace.SetLogical( wa, fname, value )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetLogical() error %d", err) ) )
   }
}

func SetMilliseconds( wa uint, fname string, ms int ) {
   err := openace.SetMilliseconds( wa, fname, ms )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetMilliseconds() error %d", err) ) )
   }
}

func SetRelKeyPos( wa uint, pos float64 ) {
   err := openace.SetRelKeyPos( wa, pos )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetRelKeyPos() error %d", err) ) )
   }
}

func SetSearchPath( path string ) {
   err := openace.SetSearchPath( path )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetSearchPath() error %d", err) ) )
   }
}

func SetString( wa uint, fname string, buf string ) {
   err := openace.SetString( wa, fname, buf, uint( len( buf ) ) )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("SetString() error %d", err) ) )
   }
}

func ShowDeleted( show_deleted uint16 ) {
   err := openace.ShowDeleted( show_deleted )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("ShowDeleted() error %d", err) ) )
   }
}

func Skip( wa uint, num_recs int ) {
   err := openace.Skip( wa, num_recs )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("Skip() error %d", err) ) )
   }
}

func UnlockRecord( wa uint, recno uint ) {
   err := openace.UnlockRecord( wa, recno )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("UnlockRecord() error %d", err) ) )
   }
}

func UnlockTable( wa uint ) {
   err := openace.UnlockTable( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("UnlockTable() error %d", err) ) )
   }
}

func WriteRecord( wa uint ) {
   err := openace.WriteRecord( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("WriteRecord() error %d", err) ) )
   }
}

func ZapTable( wa uint ) {
   err := openace.ZapTable( wa )
   if err > 0 {
      panic( errors.New( fmt.Sprintf("ZapTable() error %d", err) ) )
   }
}
