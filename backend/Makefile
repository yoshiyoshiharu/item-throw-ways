include .env

local:
	sam build
	sam local start-api --docker-network go-academy --parameter-overrides\
			DbUser=${DB_USER_DEV}\
			DbPassword=${DB_PASSWORD_DEV}\
			DbName=${DB_NAME_DEV}\
			DbHost=${DB_HOST_DEV}\
			HiraganaTranslationAppId=${HIRAGANA_TRANSLATION_APP_ID_DEV}
local_update_items:
	sam build
	sam local invoke UpdateItemsFunction --docker-network go-academy --parameter-overrides\
			DbUser=${DB_USER_DEV}\
			DbPassword=${DB_PASSWORD_DEV}\
			DbName=${DB_NAME_DEV}\
			DbHost=${DB_HOST_DEV}\
			HiraganaTranslationAppId=${HIRAGANA_TRANSLATION_APP_ID_DEV}
local_update_collect_weekdays:
	sam build
	sam local invoke UpdateCollectWeekdaysFunction --docker-network go-academy --parameter-overrides\
			DbUser=${DB_USER_DEV}\
			DbPassword=${DB_PASSWORD_DEV}\
			DbName=${DB_NAME_DEV}\
			DbHost=${DB_HOST_DEV}\
			HiraganaTranslationAppId=${HIRAGANA_TRANSLATION_APP_ID_DEV}
deploy:
	sam build
	sam deploy --parameter-overrides\
			DbUser=${DB_USER}\
			DbPassword=${DB_PASSWORD}\
			DbName=${DB_NAME}\
			DbHost=${DB_HOST}\
			HiraganaTranslationAppId=${HIRAGANA_TRANSLATION_APP_ID}

