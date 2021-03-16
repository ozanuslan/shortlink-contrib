# SKAFFOLD TASKS =======================================================================================================
skaffold-init: ## Run local minikube and set default params
	@make minikube-up
	-kubectl create namespace shortlink

skaffold-up: ## Run local minikube and set default params
	@skaffold dev

skaffold-debug: ## Run local minikube and set default params with debug mode
	@skaffold debug

skaffold-down: ## Down dev-workspace