TARGET = aviCommandTester
RELEASE_DIR = release

GIT_HOOKS := .git/hooks/applied

#ALL: $(GIT_HOOKS) bridge

.PHONY: bridge
bridge:
	@make -C $@

$(GIT_HOOKS):
	@scripts/install-git-hooks.sh
	@echo
