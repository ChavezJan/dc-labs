# build & test automation

DPIP_ENDPOINT = http://localhost:8080
API_USER      = username
API_PASSWORD  = password

test:
	@echo Login
	@$(eval TOKEN := $(shell curl -u ${API_USER}:${API_PASSWORD} ${DPIP_ENDPOINT}/login | jq -r '.token'))
	@echo Logout
	@$(shell curl -H "Authorization: Bearer ${TOKEN}" ${DPIP_ENDPOINT}/logout)
	@echo Upload Image
	@$(eval IMAGE_PATH := $(shell realpath test.jpg))
	@$(shell curl -F 'data=@${IMAGE_PATH}' -H "Authorization: Bearer ${TOKEN}" ${DPIP_ENDPOINT}/upload)
	@echo Status
	@$(shell curl -H "Authorization: Bearer $(TOKEN)" ${DPIP_ENDPOINT}/status)
