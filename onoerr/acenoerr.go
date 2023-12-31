/* This file is automatically generated by goaddgen.prg originated from
   https://github.com/alcz/openace-go
   goaddgen.txt was used as input containing simple suggestions
   on how to forward to plain C-API/CGo/SWIG */

package onoerr

import "github.com/alcz/openace-go/oauto"
import openace "github.com/alcz/openace-go"

func SetCallbackOnRTE( fp uintptr )(_ret bool) {
   err := openace.SetCallbackOnRTE( fp )
   if err > 0 {
      return false
   }
   return true
}

func SetCallbackTrace( fp uintptr )(_ret bool) {
   err := openace.SetCallbackTrace( fp )
   if err > 0 {
      return false
   }
   return true
}

func AddCustomKey( index uint )(_ret bool) {
   err := openace.AddCustomKey( index )
   if err > 0 {
      return false
   }
   return true
}

func AppendRecord( wa uint )(_ret bool) {
   err := openace.AppendRecord( wa )
   if err > 0 {
      return false
   }
   return true
}

func ApplicationExit()(_ret bool) {
   err := openace.ApplicationExit(  )
   if err > 0 {
      return false
   }
   return true
}

func AtBOF( wa uint )(_ret bool) {
   var ret uint16 = 1
   if 0 == openace.AtBOF( wa, &ret ) {
      return ret != 0
   }
   return true
}

func AtEOF( wa uint )(_ret bool) {
   var ret uint16 = 1
   if 0 == openace.AtEOF( wa, &ret ) {
      return ret != 0
   }
   return true
}

func CacheRecords( wa uint, num_recs uint16 )(_ret bool) {
   err := openace.CacheRecords( wa, num_recs )
   if err > 0 {
      return false
   }
   return true
}

func BeginTransaction( conn uint )(_ret bool) {
   err := openace.BeginTransaction( conn )
   if err > 0 {
      return false
   }
   return true
}

func ClearFilter( wa uint )(_ret bool) {
   err := openace.ClearFilter( wa )
   if err > 0 {
      return false
   }
   return true
}

func ClearProgressCallback()(_ret bool) {
   err := openace.ClearProgressCallback(  )
   if err > 0 {
      return false
   }
   return true
}

func ClearRelation( master_wa uint )(_ret bool) {
   err := openace.ClearRelation( master_wa )
   if err > 0 {
      return false
   }
   return true
}

func ClearScope( index uint, opt uint16 )(_ret bool) {
   err := openace.ClearScope( index, opt )
   if err > 0 {
      return false
   }
   return true
}

func CloseAllIndexes( wa uint )(_ret bool) {
   err := openace.CloseAllIndexes( wa )
   if err > 0 {
      return false
   }
   return true
}

func CloseIndex( index uint )(_ret bool) {
   err := openace.CloseIndex( index )
   if err > 0 {
      return false
   }
   return true
}

func CloseTable( wa uint )(_ret bool) {
   err := openace.CloseTable( wa )
   if err > 0 {
      return false
   }
   return true
}

func CommitTransaction( conn uint )(_ret bool) {
   err := openace.CommitTransaction( conn )
   if err > 0 {
      return false
   }
   return true
}

func Connect( server_name string )(_ret uint) {
   var ret uint = 0
   openace.Connect( server_name, &ret )
   return ret
}

func CopyTableContents( handle uint, wa_to uint, filter_opt uint16 )(_ret bool) {
   err := openace.CopyTableContents( handle, wa_to, filter_opt )
   if err > 0 {
      return false
   }
   return true
}

func CreateIndex( handle uint, filename string, tag string, expr string, cond string, while string, opt uint )(_ret uint) {
   var ret uint = 0
   openace.CreateIndex( handle, filename, tag, expr, cond, while, opt, &ret )
   return ret
}

func CreateTable( conn uint, name string, alias string, table_type uint16, char_type uint16, lock_type uint16, check_rights uint16, memo_size uint16, fields string )(_ret uint) {
   var ret uint = 0
   openace.CreateTable( conn, name, alias, table_type, char_type, lock_type, check_rights, memo_size, fields, &ret )
   return ret
}

func DeleteCustomKey( index uint )(_ret bool) {
   err := openace.DeleteCustomKey( index )
   if err > 0 {
      return false
   }
   return true
}

func DeleteIndex( index uint )(_ret bool) {
   err := openace.DeleteIndex( index )
   if err > 0 {
      return false
   }
   return true
}

func DeleteRecord( wa uint )(_ret bool) {
   err := openace.DeleteRecord( wa )
   if err > 0 {
      return false
   }
   return true
}

func Disconnect( conn uint )(_ret bool) {
   err := openace.Disconnect( conn )
   if err > 0 {
      return false
   }
   return true
}

func ExtractKey( index uint )(_ret string) {
   ret, _ := oauto.ExtractKeyErr( index )
   return ret
}

func GetBinary( wa uint, fname string, offset uint )(_ret string) {
   ret, _ := oauto.GetBinaryErr( wa, fname, offset )
   return ret
}

func GetBinaryLength( wa uint, fname string )(_ret uint) {
   var ret uint = 0
   openace.GetBinaryLength( wa, fname, &ret )
   return ret
}

func GetDateFormat(  )(_ret string) {
   ret, _ := oauto.GetDateFormatErr(  )
   return ret
}

func GetDouble( wa uint, fname string )(_ret float64) {
   var ret float64 = 0
   openace.GetDouble( wa, fname, &ret )
   return ret
}

func GetErrorString( errcode uint )(_ret string) {
   ret, _ := oauto.GetErrorStringErr( errcode )
   return ret
}

func GetField( wa uint, fname string, opt uint16 )(_ret string) {
   ret, _ := oauto.GetFieldErr( wa, fname, opt )
   return ret
}

func GetFieldDecimals( wa uint, fname string )(_ret uint16) {
   var ret uint16 = 0
   openace.GetFieldDecimals( wa, fname, &ret )
   return ret
}

func GetFieldLength( wa uint, fname string )(_ret uint) {
   var ret uint = 0
   openace.GetFieldLength( wa, fname, &ret )
   return ret
}

func GetFieldName( wa uint, pos uint16 )(_ret string) {
   ret, _ := oauto.GetFieldNameErr( wa, pos )
   return ret
}

func GetFieldType( wa uint, fname string )(_ret uint16) {
   var ret uint16 = 0
   openace.GetFieldType( wa, fname, &ret )
   return ret
}

func GetHandleType( handle uint )(_ret uint16) {
   var ret uint16 = 0
   openace.GetHandleType( handle, &ret )
   return ret
}

func GetIndexCondition( index uint )(_ret string) {
   ret, _ := oauto.GetIndexConditionErr( index )
   return ret
}

func GetIndexExpr( index uint )(_ret string) {
   ret, _ := oauto.GetIndexExprErr( index )
   return ret
}

func GetIndexFilename( index uint, opt uint16 )(_ret string) {
   ret, _ := oauto.GetIndexFilenameErr( index, opt )
   return ret
}

func GetIndexName( index uint )(_ret string) {
   ret, _ := oauto.GetIndexNameErr( index )
   return ret
}

func GetIndexHandle( wa uint, tag string )(_ret uint) {
   var ret uint = 0
   openace.GetIndexHandle( wa, tag, &ret )
   return ret
}

func GetIndexHandleByOrder( wa uint, ordnum uint16 )(_ret uint) {
   var ret uint = 0
   openace.GetIndexHandleByOrder( wa, ordnum, &ret )
   return ret
}

func GetJulian( wa uint, fname string )(_ret int) {
   var ret int = 0
   openace.GetJulian( wa, fname, &ret )
   return ret
}

func GetKeyCount( index uint, filter_opt uint16 )(_ret uint) {
   var ret uint = 0
   openace.GetKeyCount( index, filter_opt, &ret )
   return ret
}

func GetKeyLength( index uint )(_ret uint16) {
   var ret uint16 = 0
   openace.GetKeyLength( index, &ret )
   return ret
}

func GetKeyNum( index uint, filter_opt uint16 )(_ret uint) {
   var ret uint = 0
   openace.GetKeyNum( index, filter_opt, &ret )
   return ret
}

func GetKeyType( index uint )(_ret uint16) {
   var ret uint16 = 0
   openace.GetKeyType( index, &ret )
   return ret
}

func GetLastError( errcode *uint )(_ret string) {
   ret, _ := oauto.GetLastErrorErr( errcode )
   return ret
}

func GetLastTableUpdate( wa uint )(_ret string) {
   ret, _ := oauto.GetLastTableUpdateErr( wa )
   return ret
}

func GetLogical( wa uint, fname string )(_ret bool) {
   var ret uint16 = 0
   if 0 == openace.GetLogical( wa, fname, &ret ) {
      return ret != 0
   }
   return false
}

func GetMemoLength( wa uint, fname string )(_ret uint) {
   var ret uint = 0
   openace.GetMemoLength( wa, fname, &ret )
   return ret
}

func GetMilliseconds( wa uint, fname string )(_ret int) {
   var ret int = 0
   openace.GetMilliseconds( wa, fname, &ret )
   return ret
}

func GetNumFields( wa uint )(_ret uint16) {
   var ret uint16 = 0
   openace.GetNumFields( wa, &ret )
   return ret
}

func GetNumIndexes( wa uint )(_ret uint16) {
   var ret uint16 = 0
   openace.GetNumIndexes( wa, &ret )
   return ret
}

func GetNumLocks( wa uint )(_ret uint16) {
   var ret uint16 = 0
   openace.GetNumLocks( wa, &ret )
   return ret
}

func GetRecordCount( wa uint, filter_opt uint16 )(_ret uint) {
   var ret uint = 0
   openace.GetRecordCount( wa, filter_opt, &ret )
   return ret
}

func GetRecordNum( wa uint, filter_opt uint16 )(_ret uint) {
   var ret uint = 0
   openace.GetRecordNum( wa, filter_opt, &ret )
   return ret
}

func GetRecordLength( wa uint )(_ret uint) {
   var ret uint = 0
   openace.GetRecordLength( wa, &ret )
   return ret
}

func GetRelKeyPos( wa uint )(_ret float64) {
   var ret float64 = 0
   openace.GetRelKeyPos( wa, &ret )
   return ret
}

func GetServerName( conn uint )(_ret string) {
   ret, _ := oauto.GetServerNameErr( conn )
   return ret
}

func GetString( wa uint, fname string, opt uint16 )(_ret string) {
   ret, _ := oauto.GetStringErr( wa, fname, opt )
   return ret
}

func GetTableAlias( wa uint )(_ret string) {
   ret, _ := oauto.GetTableAliasErr( wa )
   return ret
}

func GetTableFilename( wa uint, opt uint16 )(_ret string) {
   ret, _ := oauto.GetTableFilenameErr( wa, opt )
   return ret
}

func GetTableType( wa uint )(_ret uint16) {
   var ret uint16 = 0
   openace.GetTableType( wa, &ret )
   return ret
}

func GotoBottom( wa uint )(_ret bool) {
   err := openace.GotoBottom( wa )
   if err > 0 {
      return false
   }
   return true
}

func GotoRecord( wa uint, recno uint )(_ret bool) {
   err := openace.GotoRecord( wa, recno )
   if err > 0 {
      return false
   }
   return true
}

func GotoTop( wa uint )(_ret bool) {
   err := openace.GotoTop( wa )
   if err > 0 {
      return false
   }
   return true
}

func IsEmpty( wa uint, fname string )(_ret bool) {
   var ret uint16 = 1
   if 0 == openace.IsEmpty( wa, fname, &ret ) {
      return ret != 0
   }
   return true
}

func IsIndexDescending( index uint )(_ret bool) {
   var ret uint16 = 0
   if 0 == openace.IsIndexDescending( index, &ret ) {
      return ret != 0
   }
   return false
}

func IsIndexUnique( index uint )(_ret bool) {
   var ret uint16 = 0
   if 0 == openace.IsIndexUnique( index, &ret ) {
      return ret != 0
   }
   return false
}

func IsRecordDeleted( wa uint )(_ret bool) {
   var ret uint16 = 1
   if 0 == openace.IsRecordDeleted( wa, &ret ) {
      return ret != 0
   }
   return true
}

func LockRecord( wa uint, recno uint )(_ret bool) {
   err := openace.LockRecord( wa, recno )
   if err > 0 {
      return false
   }
   return true
}

func LockTable( wa uint )(_ret bool) {
   err := openace.LockTable( wa )
   if err > 0 {
      return false
   }
   return true
}

func MgConnect( server string, username string, password string )(_ret uint) {
   var ret uint = 0
   openace.MgConnect( server, username, password, &ret )
   return ret
}

func MgDisconnect( handle uint )(_ret bool) {
   err := openace.MgDisconnect( handle )
   if err > 0 {
      return false
   }
   return true
}

func MgGetServerType( handle uint )(_ret uint16) {
   var ret uint16 = 0
   openace.MgGetServerType( handle, &ret )
   return ret
}

func NullTerminateStrings( null_term uint16 )(_ret bool) {
   err := openace.NullTerminateStrings( null_term )
   if err > 0 {
      return false
   }
   return true
}

func OpenTable( conn uint, name string, alias string, table_type uint16, char_type uint16, lock_type uint16, check_rights uint16, opts uint )(_ret uint) {
   var ret uint = 0
   openace.OpenTable( conn, name, alias, table_type, char_type, lock_type, check_rights, opts, &ret )
   return ret
}

func PackTable( wa uint )(_ret bool) {
   err := openace.PackTable( wa )
   if err > 0 {
      return false
   }
   return true
}

func RecallRecord( wa uint )(_ret bool) {
   err := openace.RecallRecord( wa )
   if err > 0 {
      return false
   }
   return true
}

func RefreshRecord( wa uint )(_ret bool) {
   err := openace.RefreshRecord( wa )
   if err > 0 {
      return false
   }
   return true
}

func Reindex( wa uint )(_ret bool) {
   err := openace.Reindex( wa )
   if err > 0 {
      return false
   }
   return true
}

func RollbackTransaction( conn uint )(_ret bool) {
   err := openace.RollbackTransaction( conn )
   if err > 0 {
      return false
   }
   return true
}

func Seek( index uint, key string, data_type uint16, seek_type uint16 )(_ret bool) {
   var ret uint16 = 0
   if 0 == openace.Seek( index, key, uint16( len( key ) ), data_type, seek_type, &ret ) {
      return ret != 0
   }
   return false
}

func SeekLast( index uint, key string, data_type uint16 )(_ret bool) {
   var ret uint16 = 0
   if 0 == openace.SeekLast( index, key, uint16( len( key ) ), data_type, &ret ) {
      return ret != 0
   }
   return false
}

func SetDateFormat( format string )(_ret bool) {
   err := openace.SetDateFormat( format )
   if err > 0 {
      return false
   }
   return true
}

func SetDefault( def string )(_ret bool) {
   err := openace.SetDefault( def )
   if err > 0 {
      return false
   }
   return true
}

func SetDouble( wa uint, fname string, value float64 )(_ret bool) {
   err := openace.SetDouble( wa, fname, value )
   if err > 0 {
      return false
   }
   return true
}

func SetEmpty( wa uint, fname string )(_ret bool) {
   err := openace.SetEmpty( wa, fname )
   if err > 0 {
      return false
   }
   return true
}

func SetField( wa uint, fname string, buf string )(_ret bool) {
   err := openace.SetField( wa, fname, buf, uint( len( buf ) ) )
   if err > 0 {
      return false
   }
   return true
}

func SetFilter( wa uint, filter string )(_ret bool) {
   err := openace.SetFilter( wa, filter )
   if err > 0 {
      return false
   }
   return true
}

func SetJulian( wa uint, fname string, julian int )(_ret bool) {
   err := openace.SetJulian( wa, fname, julian )
   if err > 0 {
      return false
   }
   return true
}

func SetLogical( wa uint, fname string, value uint16 )(_ret bool) {
   err := openace.SetLogical( wa, fname, value )
   if err > 0 {
      return false
   }
   return true
}

func SetMilliseconds( wa uint, fname string, ms int )(_ret bool) {
   err := openace.SetMilliseconds( wa, fname, ms )
   if err > 0 {
      return false
   }
   return true
}

func SetRelKeyPos( wa uint, pos float64 )(_ret bool) {
   err := openace.SetRelKeyPos( wa, pos )
   if err > 0 {
      return false
   }
   return true
}

func SetSearchPath( path string )(_ret bool) {
   err := openace.SetSearchPath( path )
   if err > 0 {
      return false
   }
   return true
}

func SetString( wa uint, fname string, buf string )(_ret bool) {
   err := openace.SetString( wa, fname, buf, uint( len( buf ) ) )
   if err > 0 {
      return false
   }
   return true
}

func ShowDeleted( show_deleted uint16 )(_ret bool) {
   err := openace.ShowDeleted( show_deleted )
   if err > 0 {
      return false
   }
   return true
}

func Skip( wa uint, num_recs int )(_ret bool) {
   err := openace.Skip( wa, num_recs )
   if err > 0 {
      return false
   }
   return true
}

func UnlockRecord( wa uint, recno uint )(_ret bool) {
   err := openace.UnlockRecord( wa, recno )
   if err > 0 {
      return false
   }
   return true
}

func UnlockTable( wa uint )(_ret bool) {
   err := openace.UnlockTable( wa )
   if err > 0 {
      return false
   }
   return true
}

func WriteRecord( wa uint )(_ret bool) {
   err := openace.WriteRecord( wa )
   if err > 0 {
      return false
   }
   return true
}

func ZapTable( wa uint )(_ret bool) {
   err := openace.ZapTable( wa )
   if err > 0 {
      return false
   }
   return true
}
