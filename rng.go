package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags)
	//Declare Variables
	var numcount int
	var valuemin int
	var valuemax int
	var waittime int64
	//Declare Flags
	flag.IntVar(&numcount, "ncount", 100, "Number of random numbers generated per run.")
	flag.IntVar(&valuemin, "nmin", 1, "Minimum value of a single random number.")
	flag.IntVar(&valuemax, "nmax", 1000, "Maximum value of a single random number.")
	flag.Int64Var(&waittime, "nwait", 100, "Time in milliseconds between making each number.")
	flag.Parse()
	//Prepare Log File/Functions
	log.SetPrefix("")
	log.SetFlags(0)
	f, err := os.OpenFile("RNG.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	//Start Program
	t0 := time.Now()
	fmt.Println("Welcome to the Random Number Generator v1.0 by redstone2019.")
	fmt.Println("This will create a list of random numbers based on -valuemin, -valuemax, and numcount.")
	time.Sleep(2000 * time.Millisecond)
	print("Starting")
	time.Sleep(500 * time.Millisecond)
	print(".")
	time.Sleep(500 * time.Millisecond)
	print(".")
	time.Sleep(500 * time.Millisecond)
	print(".")
	println("")
	time.Sleep(500 * time.Millisecond)
	//Random Number Generator
	rand.Seed(time.Now().UnixNano())
	log.Println("New Run:")
	for numcount > 0 {
		println("")
		println(numcount, "numbers to go")
		print("Calculating...")
		random := (rand.Intn(valuemax-valuemin+1) + valuemin)
		print(" ", random)
		log.Println(random)
		println("")
		numcount -= 1
		time.Sleep(time.Duration(waittime) * time.Millisecond)
	}
	t1 := time.Now()
	//Completion Message
	println("Done, check RNG.log for your ", numcount, " random numbers")
	fmt.Printf("That took %v \n", t1.Sub(t0))
	log.Printf("That took %v \n", t1.Sub(t0))
	//Program End Countdown
	x := 5
	for x > 0 {
		print(x, "...")
		time.Sleep(1 * time.Second)
		x -= 1
	}
}
