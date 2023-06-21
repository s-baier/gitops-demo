run: build start

build:
	mdslides --include presentation ./presentation/slides.md

start:
	python -m http.server -d slides 8000
