package openace

const FALSE uint16 = 0
const TRUE uint16 = 1
const DEFAULT uint16 = 0

const REG_OWNER_LEN int = 36
const REVISION_LEN int = 16
const INST_DATE_LEN int = 16
const OEM_CHAR_NAME_LEN int = 16
const ANSI_CHAR_NAME_LEN int = 16
const SERIAL_NUM_LEN int = 16

const EXCLUSIVE uint = 0x00000001
const READONLY uint = 0x00000002
const SHARED uint = 0x00000004
const CLIPPER_MEMOS uint = 0x00000008
const TABLE_PERM_READ uint = 0x00000010
const TABLE_PERM_UPDATE uint = 0x00000020
const TABLE_PERM_INSERT uint = 0x00000040
const TABLE_PERM_DELETE uint = 0x00000080
const REINDEX_ON_COLLATION_MISMATCH uint = 0x00000100
const IGNORE_COLLATION_MISMATCH uint = 0x00000200
const FREE_TABLE uint = 0x00001000
const TEMP_TABLE uint = 0x00002000
const DICTIONARY_BOUND_TABLE uint = 0x00004000

const CACHE_READS uint = 0x20000000
const CACHE_WRITES uint = 0x40000000

const ASCENDING uint = 0x00000000
const UNIQUE uint = 0x00000001
const COMPOUND uint = 0x00000002
const CUSTOM uint = 0x00000004
const DESCENDING uint = 0x00000008
const USER_DEFINED uint = 0x00000010

const RAWKEY uint16 = 1
const STRINGKEY uint16 = 2
const DOUBLEKEY uint16 = 4
const WSTRINGKEY uint16 = 8

const SOFTSEEK uint16 = 0x0001
const HARDSEEK uint16 = 0x0002
const SEEKGT uint16 = 0x0004

const TOP uint16 = 1
const BOTTOM uint16 = 2

const RESPECTFILTERS uint16 = 1
const IGNOREFILTERS uint16 = 2
const RESPECTSCOPES uint16 = 3

const DATABASE_TABLE uint16 = 0
const NTX uint16 = 1
const CDX uint16 = 2
const ADT uint16 = 3
const VFP uint16 = 4

const TYPE_UNKNOWN uint = 0
const LOGICAL uint = 1
const NUMERIC uint = 2
const DATE uint = 3
const STRING uint = 4
const MEMO uint = 5
const BINARY uint = 6
const IMAGE uint = 7
const VARCHAR uint = 8
const COMPACTDATE uint = 9
const DOUBLE uint = 10
const INTEGER uint = 11
const SHORTINT uint = 12
const TIME uint = 13
const TIMESTAMP uint = 14
const AUTOINC uint = 15
const RAW uint = 16
const CURDOUBLE uint = 17
const MONEY uint = 18
const LONGLONG uint = 19
const CISTRING uint = 20
const ROWVERSION uint = 21
const MODTIME uint = 22
const VARCHAR_FOX uint = 23
const VARBINARY_FOX uint = 24
const SYSTEM_FIELD uint = 25
const NCHAR uint = 26
const NVARCHAR uint = 27
const NMEMO uint = 28

const NONE uint16 = 0x0000
const LTRIM uint16 = 0x0001
const RTRIM uint16 = 0x0002
const TRIM uint16 = 0x0003
const GET_UTF8 uint16 = 0x0004
const DONT_CHECK_CONV_ERR uint16 = 0x0008
const GET_FORMAT_ANSI uint16 = 0x0010
const GET_FORMAT_WEB uint16 = 0x0030

const CONNECTION uint = 1
const TABLE uint = 2
const INDEX_ORDER uint = 3
const STATEMENT uint = 4
const CURSOR uint = 5
const DATABASE_CONNECTION uint = 6
const FTS_INDEX_ORDER uint = 8
const MGMT_NT uint = 2
const MGMT_LOCAL uint = 3
const MGMT_LINUX uint = 6
const MGMT_NT_64 uint = 7
const MGMT_LINUX_64 uint = 8
