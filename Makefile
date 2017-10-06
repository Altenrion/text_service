cloudconf:
	gcloud config set account landerfeld@gmail.com
	gcloud config set project text-service
	gcloud config set compute/zone europe-west1-b
	gcloud config set compute/region europe-west1

cloudinitproject:
	gcloud config set project text-service

clusterlist:
	gcloud container clusters list

deployments:
	kubectl create -f backend/kubernetes.yaml