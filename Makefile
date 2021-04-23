.PHONY: all
all:
	make -C setup
	make -C do

clean:
	make -C setup clean
	make -C do clean
