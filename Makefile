run:
	go build -i ./cmd/myleague
	./myleague

generateAPI:
	./hack/gen-api.sh

plot:
	./plot/plot.py
