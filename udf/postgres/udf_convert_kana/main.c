#include <postgres.h>
#include <fmgr.h>
#include <stdlib.h>
#include <string.h>
#include "_cgo_export.h"

PG_MODULE_MAGIC;

PG_FUNCTION_INFO_V1(udf_convert_kana);

Datum
udf_convert_kana(PG_FUNCTION_ARGS)
{
	text  *raw_arg1 = PG_GETARG_TEXT_PP(0);
	text  *raw_arg2 = PG_GETARG_TEXT_PP(1);
	int32 raw_arg1_size = VARSIZE_ANY_EXHDR(raw_arg1);
	int32 raw_arg2_size = VARSIZE_ANY_EXHDR(raw_arg2);
	char *arg1 = (char *) palloc(raw_arg1_size + 1);
	char *arg2 = (char *) palloc(raw_arg2_size + 1);
	strncpy(arg1, VARDATA_ANY(raw_arg1), raw_arg1_size);
	strncpy(arg2, VARDATA_ANY(raw_arg2), raw_arg2_size);
    // text type is not null character terminated
	arg1[raw_arg1_size] = '\0';
	arg2[raw_arg2_size] = '\0';

	struct udf_go_convert_kana_return r = udf_go_convert_kana(arg1, arg2);
	if (r.r1 != NULL) {
		char *msg = (char *)palloc(strlen(r.r1) + 1);
		strcpy(msg, r.r1);
		free(r.r1);
		ereport(ERROR, (errcode(ERRCODE_INVALID_PARAMETER_VALUE), errmsg("%s", msg)));
	}

	int32 new_text_size = strlen(r.r0) + VARHDRSZ;
	text *new_text = (text *) palloc(new_text_size);
	SET_VARSIZE(new_text, new_text_size);
	memcpy(VARDATA(new_text), r.r0, strlen(r.r0));
	free(r.r0);

	PG_RETURN_TEXT_P(new_text);
}