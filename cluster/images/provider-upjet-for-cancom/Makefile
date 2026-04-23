# ====================================================================================
# Setup Project

include ../../../build/makelib/common.mk

# ====================================================================================
#  Options

include ../../../build/makelib/imagelight.mk

# ====================================================================================
# Targets

img.build:
	@$(INFO) docker build $(IMAGE)
	@$(MAKE) BUILD_ARGS="--load" img.build.shared
	@$(OK) docker build $(IMAGE)

# img.publish is intentionally a no-op. Upstream upjet template overrides this
# target to suppress the standard imagelight.mk Docker publish pipeline. The
# provider image is published exclusively via xpkg machinery (see xpkg.mk).
img.publish:
#	@$(INFO) Skipping image publish for $(IMAGE)
#	@echo Publish is deferred to xpkg machinery
#	@$(OK) Image publish skipped for $(IMAGE)

img.build.shared:
	@cp Dockerfile $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp terraformrc.hcl $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp -r $(OUTPUT_DIR)/bin/ $(IMAGE_TEMP_DIR)/bin || $(FAIL)
	@docker buildx build $(BUILD_ARGS) \
		--platform $(IMAGE_PLATFORMS) \
		--build-arg TERRAFORM_VERSION=$(TERRAFORM_VERSION) \
		--build-arg TERRAFORM_PROVIDER_SOURCE=$(TERRAFORM_PROVIDER_SOURCE) \
		--build-arg TERRAFORM_PROVIDER_VERSION=$(TERRAFORM_PROVIDER_VERSION) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_NAME=$(TERRAFORM_PROVIDER_DOWNLOAD_NAME) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX=$(TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX) \
		--build-arg TERRAFORM_NATIVE_PROVIDER_BINARY=$(TERRAFORM_NATIVE_PROVIDER_BINARY) \
		-t $(IMAGE) \
		$(IMAGE_TEMP_DIR) || $(FAIL)

# img.promote is intentionally a no-op. Upstream upjet template overrides this
# target to suppress the standard imagelight.mk Docker promote pipeline. The
# provider image is promoted exclusively via xpkg machinery (see xpkg.mk).
img.promote:
#	@$(INFO) Skipping image promotion from $(FROM_IMAGE) to $(TO_IMAGE)
#	@echo Promote is deferred to xpkg machinery
#	@$(OK) Image promotion skipped for $(FROM_IMAGE) to $(TO_IMAGE)
