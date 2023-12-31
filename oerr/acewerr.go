/* This file is automatically generated by goaddgen.prg originated from
   https://github.com/alcz/openace-go
   goaddgen.txt was used as input containing simple suggestions
   on how to forward to plain C-API/CGo/SWIG */

package oerr

import "github.com/alcz/openace-go/oauto"
import openace "github.com/alcz/openace-go"

func SetCallbackOnRTE( fp uintptr )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetCallbackOnRTE( fp )
   ret = err == 0
   return ret, err
}

func SetCallbackTrace( fp uintptr )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetCallbackTrace( fp )
   ret = err == 0
   return ret, err
}

func AddCustomKey( index uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.AddCustomKey( index )
   ret = err == 0
   return ret, err
}

func AppendRecord( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.AppendRecord( wa )
   ret = err == 0
   return ret, err
}

func ApplicationExit()(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.ApplicationExit(  )
   ret = err == 0
   return ret, err
}

func AtBOF( wa uint )(_ret bool, _ret_err uint) {
   var ret uint16 = 1
   err := openace.AtBOF( wa, &ret )
   if err > 0 {
      return true, err
   }
   return ret != 0, err
}

func AtEOF( wa uint )(_ret bool, _ret_err uint) {
   var ret uint16 = 1
   err := openace.AtEOF( wa, &ret )
   if err > 0 {
      return true, err
   }
   return ret != 0, err
}

func CacheRecords( wa uint, num_recs uint16 )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.CacheRecords( wa, num_recs )
   ret = err == 0
   return ret, err
}

func BeginTransaction( conn uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.BeginTransaction( conn )
   ret = err == 0
   return ret, err
}

func ClearFilter( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.ClearFilter( wa )
   ret = err == 0
   return ret, err
}

func ClearProgressCallback()(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.ClearProgressCallback(  )
   ret = err == 0
   return ret, err
}

func ClearRelation( master_wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.ClearRelation( master_wa )
   ret = err == 0
   return ret, err
}

func ClearScope( index uint, opt uint16 )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.ClearScope( index, opt )
   ret = err == 0
   return ret, err
}

func CloseAllIndexes( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.CloseAllIndexes( wa )
   ret = err == 0
   return ret, err
}

func CloseIndex( index uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.CloseIndex( index )
   ret = err == 0
   return ret, err
}

func CloseTable( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.CloseTable( wa )
   ret = err == 0
   return ret, err
}

func CommitTransaction( conn uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.CommitTransaction( conn )
   ret = err == 0
   return ret, err
}

func Connect( server_name string )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.Connect( server_name, &ret )
   return ret, err
}

func CopyTableContents( handle uint, wa_to uint, filter_opt uint16 )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.CopyTableContents( handle, wa_to, filter_opt )
   ret = err == 0
   return ret, err
}

func CreateIndex( handle uint, filename string, tag string, expr string, cond string, while string, opt uint )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.CreateIndex( handle, filename, tag, expr, cond, while, opt, &ret )
   return ret, err
}

func CreateTable( conn uint, name string, alias string, table_type uint16, char_type uint16, lock_type uint16, check_rights uint16, memo_size uint16, fields string )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.CreateTable( conn, name, alias, table_type, char_type, lock_type, check_rights, memo_size, fields, &ret )
   return ret, err
}

func DeleteCustomKey( index uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.DeleteCustomKey( index )
   ret = err == 0
   return ret, err
}

func DeleteIndex( index uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.DeleteIndex( index )
   ret = err == 0
   return ret, err
}

func DeleteRecord( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.DeleteRecord( wa )
   ret = err == 0
   return ret, err
}

func Disconnect( conn uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.Disconnect( conn )
   ret = err == 0
   return ret, err
}

func ExtractKey( index uint )(_ret string, _ret_err uint) {
   ret, err := oauto.ExtractKeyErr( index )
   return ret, err
}

func GetBinary( wa uint, fname string, offset uint )(_ret string, _ret_err uint) {
   ret, err := oauto.GetBinaryErr( wa, fname, offset )
   return ret, err
}

func GetBinaryLength( wa uint, fname string )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.GetBinaryLength( wa, fname, &ret )
   return ret, err
}

func GetDateFormat(  )(_ret string, _ret_err uint) {
   ret, err := oauto.GetDateFormatErr(  )
   return ret, err
}

func GetDouble( wa uint, fname string )(_ret float64, _ret_err uint) {
   var ret float64 = 0
   err := openace.GetDouble( wa, fname, &ret )
   return ret, err
}

func GetErrorString( errcode uint )(_ret string, _ret_err uint) {
   ret, err := oauto.GetErrorStringErr( errcode )
   return ret, err
}

func GetField( wa uint, fname string, opt uint16 )(_ret string, _ret_err uint) {
   ret, err := oauto.GetFieldErr( wa, fname, opt )
   return ret, err
}

func GetFieldDecimals( wa uint, fname string )(_ret uint16, _ret_err uint) {
   var ret uint16 = 0
   err := openace.GetFieldDecimals( wa, fname, &ret )
   return ret, err
}

func GetFieldLength( wa uint, fname string )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.GetFieldLength( wa, fname, &ret )
   return ret, err
}

func GetFieldName( wa uint, pos uint16 )(_ret string, _ret_err uint) {
   ret, err := oauto.GetFieldNameErr( wa, pos )
   return ret, err
}

func GetFieldType( wa uint, fname string )(_ret uint16, _ret_err uint) {
   var ret uint16 = 0
   err := openace.GetFieldType( wa, fname, &ret )
   return ret, err
}

func GetHandleType( handle uint )(_ret uint16, _ret_err uint) {
   var ret uint16 = 0
   err := openace.GetHandleType( handle, &ret )
   return ret, err
}

func GetIndexCondition( index uint )(_ret string, _ret_err uint) {
   ret, err := oauto.GetIndexConditionErr( index )
   return ret, err
}

func GetIndexExpr( index uint )(_ret string, _ret_err uint) {
   ret, err := oauto.GetIndexExprErr( index )
   return ret, err
}

func GetIndexFilename( index uint, opt uint16 )(_ret string, _ret_err uint) {
   ret, err := oauto.GetIndexFilenameErr( index, opt )
   return ret, err
}

func GetIndexName( index uint )(_ret string, _ret_err uint) {
   ret, err := oauto.GetIndexNameErr( index )
   return ret, err
}

func GetIndexHandle( wa uint, tag string )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.GetIndexHandle( wa, tag, &ret )
   return ret, err
}

func GetIndexHandleByOrder( wa uint, ordnum uint16 )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.GetIndexHandleByOrder( wa, ordnum, &ret )
   return ret, err
}

func GetJulian( wa uint, fname string )(_ret int, _ret_err uint) {
   var ret int = 0
   err := openace.GetJulian( wa, fname, &ret )
   return ret, err
}

func GetKeyCount( index uint, filter_opt uint16 )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.GetKeyCount( index, filter_opt, &ret )
   return ret, err
}

func GetKeyLength( index uint )(_ret uint16, _ret_err uint) {
   var ret uint16 = 0
   err := openace.GetKeyLength( index, &ret )
   return ret, err
}

func GetKeyNum( index uint, filter_opt uint16 )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.GetKeyNum( index, filter_opt, &ret )
   return ret, err
}

func GetKeyType( index uint )(_ret uint16, _ret_err uint) {
   var ret uint16 = 0
   err := openace.GetKeyType( index, &ret )
   return ret, err
}

func GetLastError( errcode *uint )(_ret string, _ret_err uint) {
   ret, err := oauto.GetLastErrorErr( errcode )
   return ret, err
}

func GetLastTableUpdate( wa uint )(_ret string, _ret_err uint) {
   ret, err := oauto.GetLastTableUpdateErr( wa )
   return ret, err
}

func GetLogical( wa uint, fname string )(_ret bool, _ret_err uint) {
   var ret uint16 = 0
   err := openace.GetLogical( wa, fname, &ret )
   return ret != 0, err
}

func GetMemoLength( wa uint, fname string )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.GetMemoLength( wa, fname, &ret )
   return ret, err
}

func GetMilliseconds( wa uint, fname string )(_ret int, _ret_err uint) {
   var ret int = 0
   err := openace.GetMilliseconds( wa, fname, &ret )
   return ret, err
}

func GetNumFields( wa uint )(_ret uint16, _ret_err uint) {
   var ret uint16 = 0
   err := openace.GetNumFields( wa, &ret )
   return ret, err
}

func GetNumIndexes( wa uint )(_ret uint16, _ret_err uint) {
   var ret uint16 = 0
   err := openace.GetNumIndexes( wa, &ret )
   return ret, err
}

func GetNumLocks( wa uint )(_ret uint16, _ret_err uint) {
   var ret uint16 = 0
   err := openace.GetNumLocks( wa, &ret )
   return ret, err
}

func GetRecordCount( wa uint, filter_opt uint16 )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.GetRecordCount( wa, filter_opt, &ret )
   return ret, err
}

func GetRecordNum( wa uint, filter_opt uint16 )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.GetRecordNum( wa, filter_opt, &ret )
   return ret, err
}

func GetRecordLength( wa uint )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.GetRecordLength( wa, &ret )
   return ret, err
}

func GetRelKeyPos( wa uint )(_ret float64, _ret_err uint) {
   var ret float64 = 0
   err := openace.GetRelKeyPos( wa, &ret )
   return ret, err
}

func GetServerName( conn uint )(_ret string, _ret_err uint) {
   ret, err := oauto.GetServerNameErr( conn )
   return ret, err
}

func GetString( wa uint, fname string, opt uint16 )(_ret string, _ret_err uint) {
   ret, err := oauto.GetStringErr( wa, fname, opt )
   return ret, err
}

func GetTableAlias( wa uint )(_ret string, _ret_err uint) {
   ret, err := oauto.GetTableAliasErr( wa )
   return ret, err
}

func GetTableFilename( wa uint, opt uint16 )(_ret string, _ret_err uint) {
   ret, err := oauto.GetTableFilenameErr( wa, opt )
   return ret, err
}

func GetTableType( wa uint )(_ret uint16, _ret_err uint) {
   var ret uint16 = 0
   err := openace.GetTableType( wa, &ret )
   return ret, err
}

func GotoBottom( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.GotoBottom( wa )
   ret = err == 0
   return ret, err
}

func GotoRecord( wa uint, recno uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.GotoRecord( wa, recno )
   ret = err == 0
   return ret, err
}

func GotoTop( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.GotoTop( wa )
   ret = err == 0
   return ret, err
}

func IsEmpty( wa uint, fname string )(_ret bool, _ret_err uint) {
   var ret uint16 = 1
   err := openace.IsEmpty( wa, fname, &ret )
   if err > 0 {
      return true, err
   }
   return ret != 0, err
}

func IsIndexDescending( index uint )(_ret bool, _ret_err uint) {
   var ret uint16 = 0
   err := openace.IsIndexDescending( index, &ret )
   return ret != 0, err
}

func IsIndexUnique( index uint )(_ret bool, _ret_err uint) {
   var ret uint16 = 0
   err := openace.IsIndexUnique( index, &ret )
   return ret != 0, err
}

func IsRecordDeleted( wa uint )(_ret bool, _ret_err uint) {
   var ret uint16 = 1
   err := openace.IsRecordDeleted( wa, &ret )
   if err > 0 {
      return true, err
   }
   return ret != 0, err
}

func LockRecord( wa uint, recno uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.LockRecord( wa, recno )
   ret = err == 0
   return ret, err
}

func LockTable( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.LockTable( wa )
   ret = err == 0
   return ret, err
}

func MgConnect( server string, username string, password string )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.MgConnect( server, username, password, &ret )
   return ret, err
}

func MgDisconnect( handle uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.MgDisconnect( handle )
   ret = err == 0
   return ret, err
}

func MgGetServerType( handle uint )(_ret uint16, _ret_err uint) {
   var ret uint16 = 0
   err := openace.MgGetServerType( handle, &ret )
   return ret, err
}

func NullTerminateStrings( null_term uint16 )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.NullTerminateStrings( null_term )
   ret = err == 0
   return ret, err
}

func OpenTable( conn uint, name string, alias string, table_type uint16, char_type uint16, lock_type uint16, check_rights uint16, opts uint )(_ret uint, _ret_err uint) {
   var ret uint = 0
   err := openace.OpenTable( conn, name, alias, table_type, char_type, lock_type, check_rights, opts, &ret )
   return ret, err
}

func PackTable( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.PackTable( wa )
   ret = err == 0
   return ret, err
}

func RecallRecord( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.RecallRecord( wa )
   ret = err == 0
   return ret, err
}

func RefreshRecord( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.RefreshRecord( wa )
   ret = err == 0
   return ret, err
}

func Reindex( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.Reindex( wa )
   ret = err == 0
   return ret, err
}

func RollbackTransaction( conn uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.RollbackTransaction( conn )
   ret = err == 0
   return ret, err
}

func Seek( index uint, key string, data_type uint16, seek_type uint16 )(_ret bool, _ret_err uint) {
   var ret uint16 = 0
   err := openace.Seek( index, key, uint16( len( key ) ), data_type, seek_type, &ret )
   return ret != 0, err
}

func SeekLast( index uint, key string, data_type uint16 )(_ret bool, _ret_err uint) {
   var ret uint16 = 0
   err := openace.SeekLast( index, key, uint16( len( key ) ), data_type, &ret )
   return ret != 0, err
}

func SetDateFormat( format string )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetDateFormat( format )
   ret = err == 0
   return ret, err
}

func SetDefault( def string )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetDefault( def )
   ret = err == 0
   return ret, err
}

func SetDouble( wa uint, fname string, value float64 )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetDouble( wa, fname, value )
   ret = err == 0
   return ret, err
}

func SetEmpty( wa uint, fname string )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetEmpty( wa, fname )
   ret = err == 0
   return ret, err
}

func SetField( wa uint, fname string, buf string )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetField( wa, fname, buf, uint( len( buf ) ) )
   ret = err == 0
   return ret, err
}

func SetFilter( wa uint, filter string )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetFilter( wa, filter )
   ret = err == 0
   return ret, err
}

func SetJulian( wa uint, fname string, julian int )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetJulian( wa, fname, julian )
   ret = err == 0
   return ret, err
}

func SetLogical( wa uint, fname string, value uint16 )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetLogical( wa, fname, value )
   ret = err == 0
   return ret, err
}

func SetMilliseconds( wa uint, fname string, ms int )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetMilliseconds( wa, fname, ms )
   ret = err == 0
   return ret, err
}

func SetRelKeyPos( wa uint, pos float64 )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetRelKeyPos( wa, pos )
   ret = err == 0
   return ret, err
}

func SetSearchPath( path string )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetSearchPath( path )
   ret = err == 0
   return ret, err
}

func SetString( wa uint, fname string, buf string )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.SetString( wa, fname, buf, uint( len( buf ) ) )
   ret = err == 0
   return ret, err
}

func ShowDeleted( show_deleted uint16 )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.ShowDeleted( show_deleted )
   ret = err == 0
   return ret, err
}

func Skip( wa uint, num_recs int )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.Skip( wa, num_recs )
   ret = err == 0
   return ret, err
}

func UnlockRecord( wa uint, recno uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.UnlockRecord( wa, recno )
   ret = err == 0
   return ret, err
}

func UnlockTable( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.UnlockTable( wa )
   ret = err == 0
   return ret, err
}

func WriteRecord( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.WriteRecord( wa )
   ret = err == 0
   return ret, err
}

func ZapTable( wa uint )(_ret bool, _ret_err uint) {
   var ret bool
   err := openace.ZapTable( wa )
   ret = err == 0
   return ret, err
}
