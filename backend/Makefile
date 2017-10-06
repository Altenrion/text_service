BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean

APPNAME=text_service_server
EXENAME=github.com/altenrion/$(APPNAME)
IMAGENAME=text_service_server

PORT=8080

build:
	@echo "Prepare building...setting OS & build..."
	@echo "----------------------------------------"

	GOOG=linux go build -o $(APPNAME)
	docker build -t $(IMAGENAME) .
	@echo "----------------------------------------"

cleanbinary:
	@echo "Removing binary from local machine! "
	@echo "----------------------------------------"
	@rm -f $(APPNAME)
	@echo "Binary removed successfully! "
	@echo "----------------------------------------"

pack: build cleanbinary


localrun:
#	kill $(sudo lsof -t -i:$(PORT))
	docker run -p 8080:8080 $(APPNAME)


cloudpush:
	@echo "Change Docker tags ..."
	@echo "----------------------------------------"
	docker tag $(IMAGENAME) gcr.io/text-service/$(IMAGENAME)
	@echo "----------------------------------------"
	@echo "push Docker container ..."
	@echo "----------------------------------------"
	gcloud docker -- push gcr.io/text-service/$(IMAGENAME)
	@echo "----------------------------------------"
	@echo "Container pushed to cloud successfully! "
	@echo "----------------------------------------"



push:
	@echo "Change Docker tags ..."
	@echo "----------------------------------------"
	docker tag -f $(IMAGENAME) cc-allplus-go:5000/$(IMAGENAME):latest
	@echo "----------------------------------------"
	@echo "push Docker container ..."
	@echo "----------------------------------------"
	docker push cc-allplus-go:5000/$(IMAGENAME):latest
	@echo "----------------------------------------"
	@echo "Container pushed successfully! "
	@echo "----------------------------------------"

some:
	@echo "------------------"
	@echo "Just some space..."
	@echo "------------------"

deploy: pack push some
