ELEMENTS  ?= 4


# .PHONY: mapping
# mapping:
# 	$(eval OUTPUT="mapping/mapping.pb.go")
# 	$(eval TEMPLATE="mapping/mapping.template")
# 	echo "setting $(OUTPUT) with $(TEMPLATE)"
# 
# .PHONY: switching
# switching:
# 	$(eval OUTPUT="switching/switching.pb.go")
# 	$(eval TEMPLATE="switching/switching.template")
# 	echo "setting $(OUTPUT) with $(TEMPLATE)"
# 
# .PHONY: generate
# generate:

TARGETS = switching mapping

.PHONY: all
all: $(TARGETS)

.PHONY: $(TARGETS)
$(TARGETS):
	$(call generate_go_code,$@,$(ELEMENTS))
	$(call build_go_plugin,$@)

define generate_go_code
	target=$(1); \
	elements=$(2); \
	go run generator.go -n $$elements -f $$target/$$target.template -o $$target/$$target.pb.go
endef

define build_go_plugin
	target=$(1); \
	cd $$target; \
	go build -buildmode=plugin -o $$target.so
endef


.PHONY: clean
clean:
	@echo "remove generated files"
	@find . -type f -iname "*.so" -print -delete
	@find . -type f -iname "*.pb.go" -print -delete

.PHONY: list
list:
	@find . -type f -iname "*.so" -print
