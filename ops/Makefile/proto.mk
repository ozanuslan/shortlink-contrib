# PROTO TASKS ==========================================================================================================

proto-lint: ## Check lint
	@buf ls-files
	@buf lint
	@buf breaking --against ops/proto-lock.json

proto-lock: ## Lock proto dependencies
	@buf build -o ops/proto-lock.json

proto-generate: ## Generate proto-files
	# Link service --------------------------------------------------------------------------------
	@buf generate \
		--path=internal/services/link/domain \
		--path=internal/services/link/infrastructure \
		--template=ops/proto/link/buf.gen.yaml \
		--config=ops/proto/link/buf.yaml

	@buf generate \
		--path=internal/services/link/domain \
		--path=internal/services/link/infrastructure \
		--template=ops/proto/link/buf.gen.tag.yaml \
		--config=ops/proto/link/buf.yaml

	# Metadata service -----------------------------------------------------------------------------
	@buf generate \
		--path=internal/services/metadata/domain \
		--path=internal/services/metadata/infrastructure \
		--template=ops/proto/metadata/buf.gen.yaml \
		--config=ops/proto/metadata/buf.yaml

	# Proxy service --------------------------------------------------------------------------------
	@buf generate \
		--path=internal/services/proxy/src/proto/domain \
		--path=internal/services/proxy/src/proto/infrastructure \
		--template=ops/proto/proxy/buf.gen.yaml \
		--config=ops/proto/proxy/buf.yaml

	# Billing service -------------------------------------------------------------------------------
	@buf generate \
		--path=internal/services/billing/domain \
		--path=internal/services/billing/infrastructure \
		--template=ops/proto/billing/buf.gen.yaml \
		--config=ops/proto/billing/buf.yaml

	# Referral service ------------------------------------------------------------------------------
	@buf generate \
		--path=internal/services/referral/src/domain \
		--template=ops/proto/referral/buf.gen.yaml \
		--config=ops/proto/referral/buf.yaml

	# Eventsourcing service -------------------------------------------------------------------------
	@buf generate \
		--path=internal/pkg/eventsourcing/v1 \
		--template=ops/proto/eventsourcing/buf.gen.yaml \
		--config=ops/proto/eventsourcing/buf.yaml

	# Shortdb service -------------------------------------------------------------------------------
	@buf generate \
		--path=internal/services/shortdb/parser/v1 \
		--path=internal/services/shortdb/domain/query/v1 \
		--path=internal/services/shortdb/domain/index/v1 \
		--path=internal/services/shortdb/domain/page/v1 \
		--path=internal/services/shortdb/domain/table/v1 \
		--path=internal/services/shortdb/domain/field/v1 \
		--path=internal/services/shortdb/domain/database/v1 \
		--path=internal/services/shortdb/domain/session/v1 \
		--template=ops/proto/shortdb/buf.gen.yaml \
		--config=ops/proto/shortdb/buf.yaml

	# API-gateway service ---------------------------------------------------------------------------
	@buf generate \
		--path=internal/services/api-gateway/gateways/grpc-web/infrastructure/server/v1 \
		--template=ops/proto/api-gateway/buf.gen.yaml \
		--config=ops/proto/api-gateway/buf.yaml
