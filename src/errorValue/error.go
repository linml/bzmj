package errorValue

const (
	ERET_OK                                 = 0
	ERET_INVALID_PARAM                      = 1
	ERET_SYS_ERR                            = 2
	ERET_TIME_OUT                           = 3
	ERET_INVALID_UIN                        = 4
	ERET_DATA_NOT_READY                     = 5
	ERET_DBFATAL_ERROR                      = 6
	ERET_EMPTY_RSP                          = 7
	ERET_VIP_LV_LIMIT                       = 8
	ERET_NOT_IN_WHITELIST                   = 9
	ERET_INVALID_CHANNEL                    = 10
	ERET_NOT_LOGIN                          = 11
	ERET_CREATE_PLAYER_FAIL                 = 12
	ERET_DATA_SYS_ERR                       = 13
	ERET_DATA_INVALID_PARAM                 = 14
	ERET_DATA_DELETED                       = 15
	ERET_DATA_NOT_FOUND                     = 16
	ERET_PARAM                              = 17
	ERET_HASH_IS_FULL                       = 18
	ERET_PROTO_NO_SUCH_PACKET               = 19
	ERET_HASH_INSERT                        = 20
	ERET_GAMESVRD_START                     = 1000
	ERET_INVALID_ROOM_ID                    = 1001
	ERET_NOT_IN_TABLE                       = 1002
	ERET_NOT_IN_ROOM                        = 1003
	ERET_TABLE_NOT_INIT                     = 1004
	ERET_INVALID_SEAT_ID                    = 1005
	ERET_DEAL                               = 1006
	ERET_NO_OPERATE_MASK                    = 1007
	ERET_TAG_MJ                             = 1008
	ERET_USERINFO_NOT_FOUND                 = 1009
	ERET_HU                                 = 1010
	ERET_USER_IN_GAMING                     = 1011
	ERET_ROOM_MONEY_LIMIT                   = 1012
	ERET_REPEAT_NICK                        = 1013
	ERET_NO_SEX                             = 1014
	ERET_NO_IDENTIFY                        = 1015
	ERET_NO_SUB_FLOWER                      = 1016
	ERET_NO_CHALLENGEMODE_RANK_RECORD       = 1017
	ERET_NO_CHALLENGEMODE_RECORD            = 1018
	ERET_NO_CHALLENGEMODE                   = 1019
	ERET_TIMER_CREATE_FAIL                  = 1020
	ERET_NO_ROOM_OWNER                      = 1021
	ERET_NO_ENOUGH_GOLD                     = 1022
	ERET_ROOM_FULL                          = 1023
	ERET_TABLE_FULL                         = 1024
	ERET_TABLE_IN_GAME                      = 1025
	ERET_ROOM_DISSOLVE                      = 1026
	ERET_ROOM_IN_GAME                       = 1027
	ERET_ROOM_NOT_ALLOW_RENEW               = 1028
	ERET_INVALID_ROOM_TYPE                  = 1029
	ERET_INVALID_ROOM_BEI_SHU               = 1030
	ERET_INVALID_ROOM_MATCH_NUM             = 1031
	ERET_TABLE_STATE                        = 1032
	ERET_USER_CROWDED_OFFLINE               = 1033
	ERET_NO_SUCH_EMAIL                      = 1034
	ERET_EMAIL_REWARD_ALREADY_OPENED        = 1035
	ERET_SIGN_IN_DATE_OOR                   = 1036
	ERET_SIGN_IN_ALREADY_SIGNED             = 1037
	ERET_NO_SIGN_IN_REWARD_CFG              = 1038
	ERET_NO_WEIXIN_LOGIN_BASEINFO_NO        = 1039
	ERET_NO_SUCH_TASK                       = 1040
	ERET_TASK_NOT_FINISHED                  = 1041
	ERET_TASK_REWARD_ALREADY_RECEIVED       = 1042
	ERET_ACCOUNT_WX_PARSE_BASEINFO          = 1043
	ERET_NO_ENOUGH_DIAMOND                  = 1044
	ERET_GENERATE_CUSTOM_ROOM_ID            = 1045
	ERET_INVALID_CURRENCY_TYPE              = 1046
	ERET_NO_DIAMOND_TO_GOLD_CFG             = 1047
	ERET_SUIT_NOT_UNIFORM                   = 1048
	ERET_INVAILD_SUIT                       = 1049
	ERET_INVAILD_MJ                         = 1050
	ERET_TABLE_DISSOLVE                     = 1051
	ERET_INVALID_ROOM_DI_ZHU                = 1052
	ERET_NOT_ROOM_OWNER                     = 1053
	ERET_SYSTEM_NOT_ALLOW                   = 1054
	ERET_NOT_ALLOW_RENEW                    = 1055
	ERET_TABLE_UID_SEATID                   = 1056
	ERET_USER_NOT_ON_MATCH                  = 1057
	ERET_NO_SUCH_CLUB                       = 1058
	ERET_NO_SUCH_USER                       = 1059
	ERET_USER_ALREADY_IN_CLUB               = 1060
	ERET_USER_NOT_IN_CLUB                   = 1061
	ERET_USER_ALREADY_IN_OTHER_CLUB         = 1062
	ERET_CLUB_MEMBER_FULL                   = 1063
	ERET_MATCH_USER_NUM                     = 1064
	ERET_USER_NOT_CLUB_CAPTAIN              = 1065
	ERET_NO_SUCH_CLUB_APPLICANT             = 1066
	ERET_NO_SUCH_CLUB_MEMBER                = 1067
	ERET_NO_ALLOW_DISSOLVE_APPLY            = 1068
	ERET_ROOM_PASSWORD                      = 1069
	ERET_USER_IS_CLUB_CAPTAIN               = 1070
	ERET_USER_MIN_COIN_VALUE_NO_ENOUGH      = 1071
	ERET_HAS_CLUB_MATCH                     = 1072
	ERET_HAS_CLUB_MATCH_MAX_VALUE           = 1073
	ERET_HAS_CLUB_MATCH_NO_ALLOW_MATCH      = 1074
	ERET_HAS_CLUB_MATCH_STARTING            = 1075
	ERET_HAS_CLUB_MATCH_DELETE_NO_ALLOW     = 1076
	ERET_CLUB_ALREADY_BAN                   = 1077
	ERET_TABLE_USER_STATE                   = 1078
	ERET_CLUB_NAME_DUPLICATE                = 1079
	ERET_INVALID_MIN_CURRENCY_VALUE         = 1080
	ERET_MJ_OPERATE_IN_WAIT_DISSOLVE        = 1081
	ERET_USER_STATUS_DISABLE                = 1082
	ERET_ACCOUNTSVRD_START                  = 2000
	ERET_ACCOUNT_NOT_FOUND                  = 2001
	ERET_ACCOUNT_TOKEN_NOT_FOUND            = 2002
	ERET_ACCOUNT_UID_GENERATE               = 2003
	ERET_ACCOUNT_TOKEN_GENERATE             = 2004
	ERET_ACCOUNT_UPDATE_TOKEN               = 2005
	ERET_ACCOUNT_PHONENUM_NOT_HAVE          = 2006
	ERET_ACCOUNT_PHONENUM_ERR               = 2007
	ERET_ACCOUNT_AUTH_PASSWORD_ERR          = 2008
	ERET_ACCOUNT_CHANGE_PASSWORD_ERR        = 2009
	ERET_ACCOUNT_LOGIN_WEIXIN_FAIL          = 2010
	ERET_ACCOUNT_WX_RETURN_ERR              = 2011
	ERET_ACCOUNT_WX_NEED_AUTH               = 2012
	ERET_RECHARGESVRD_ASSIGNSERVERID_FAIL   = 4000
	ERET_RECHARGESVRD_PRODUCTCFG_FAILURE    = 4001
	ERET_RECHARGESVRD_PRODUCTCFG_INFOERR    = 4002
	ERET_RECHARGESVRD_THIRDRECPIPT_EXIST    = 4003
	ERET_RECHARGESVRD_ADD_ROLL_FAIL         = 4004
	ERET_RECHARGESVRD_GENERA_ORDERID_FAIL   = 4005
	ERET_RECHARGESVRD_GET_ORDER_FAIL        = 4006
	ERET_GAMELOGSVRD_FROMJSON_FAIL          = 5000
	ERET_GAMELOGSVRD_NO_DOCUMENT            = 5001
	ERET_GAMELOGSVRD_INSERT_INTO_MONGO_FAIL = 5002
)
