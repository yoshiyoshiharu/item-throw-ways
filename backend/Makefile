include .env

local_start_api:
	sam build --no-cached
	sam local start-api --docker-network item-throw-ways --parameter-overrides\
			DbUser=${DB_USER_DEV}\
			DbPassword=${DB_PASSWORD_DEV}\
			DbName=${DB_NAME_DEV}\
			DbHost=${DB_HOST_DEV}\
			HiraganaTranslationAppId=${HIRAGANA_TRANSLATION_APP_ID_DEV}
local_update_items:
	sam build --no-cached
	sam local invoke UpdateItemsFunction --docker-network item-throw-ways --parameter-overrides\
			DbUser=${DB_USER_DEV}\
			DbPassword=${DB_PASSWORD_DEV}\
			DbName=${DB_NAME_DEV}\
			DbHost=${DB_HOST_DEV}\
			HiraganaTranslationAppId=${HIRAGANA_TRANSLATION_APP_ID_DEV}
local_update_area_collect_weekdays:
	sam build --no-cached
	sam local invoke UpdateAreaCollectWeekdaysFunction --docker-network item-throw-ways --parameter-overrides\
			DbUser=${DB_USER_DEV}\
			DbPassword=${DB_PASSWORD_DEV}\
			DbName=${DB_NAME_DEV}\
			DbHost=${DB_HOST_DEV}\
			HiraganaTranslationAppId=${HIRAGANA_TRANSLATION_APP_ID_DEV}
local_db_init:
	mysql -h localhost -u ${DB_USER_DEV} -p -P 3306 -D ${DB_NAME_DEV} --protocol=tcp < ./db/init.sql
  
deploy:
	sam build --no-cached
	sam deploy --parameter-overrides\
			DbUser=${DB_USER}\
			DbPassword=${DB_PASSWORD}\
			DbName=${DB_NAME}\
			DbHost=${DB_HOST}\
			HiraganaTranslationAppId=${HIRAGANA_TRANSLATION_APP_ID}

