CC ?= gcc
DIST = build

$(DIST)/mpc.o: mpc.h mpc.c
	mkdir $(DIST)
	$(CC) -Wall -c mpc.c -o $(DIST)/mpc.o

foobar: $(DIST)/mpc.o example/foobar.c
	$(CC) -Wall -static example/foobar.c $(DIST)/mpc.o -o $(DIST)/foobar

clean:
	rm -rf $(DIST)