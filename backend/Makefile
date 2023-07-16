include .env

local:
	sam build
	sam local start-api
deploy:
	sam build
	sam deploy --parameter-overrides\
			DbUser=${DB_USER}\
			DbPassword=${DB_PASSWORD}\
			DbName=${DB_NAME}\
			DbHost=${DB_HOST}\
			HiraganaTranslationAppId=${HIRAGANA_TRANSLATION_APP_ID}


