# ==============================================================================
# Makefile helper functions for swagger
#

.PHONY: swagger.serve
swagger.serve: tools.verify.swagger ## 启动 swagger 在线文档（监听端口：65534）.
	@swagger serve -F=swagger --no-open --port 65534 $(ROOT_DIR)/api/openapi/openapi.yaml

.PHONY: swagger.docker
swagger.docker: ## 通过 docker 启动 swagger 在线文档（监听端口：65534）.
	@docker rm swaggerui -f && docker run -d --rm --name swaggerui \
       -p 65534:8080 \
       -v $(ROOT_DIR)/api/openapi:/tmp \
       -e SWAGGER_JSON=/tmp/openapi.yaml \
       -e PERSIST_AUTHORIZATION=true \
       swaggerapi/swagger-ui