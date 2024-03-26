CC ?= gcc
DIST = build

$(DIST)/mpc.o: mpc.h mpc.c
	mkdir $(DIST)
	$(CC) -Wall -c mpc.c -o $(DIST)/mpc.o

$(DIST)/libmpc.a: $(DIST)/mpc.o
	$(AR) rusv $(DIST)/libmpc.a $(DIST)/mpc.o

shared: $(DIST)/libmpc.a

foobar: $(DIST)/mpc.o example/foobar.c
	$(CC) -Wall -static example/foobar.c $(DIST)/mpc.o -o $(DIST)/foobar

lispy: $(DIST)/mpc.o example/lispy.c
	$(CC) -Wall -static example/lispy.c $(DIST)/mpc.o -o $(DIST)/lispy

clean:
	rm -rf $(DIST)