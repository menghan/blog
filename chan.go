package main

import (
	"log"
	"time"
)

func main() {
	// bchanEx1()
	// bchanEx2()
	bchanEx3()
	// bchanEx4()
	// bchanEx5()
	// fnbchanEx1()
	// fnbchanEx2()
	// fnbchanEx3()
	// nbchanEx1()
	// nbchanEx2()
}

// 测试先发
func bchanEx1() {
	bch := make(chan int)
	f := func (n int) {
		log.Println("to send", n)
		bch <- n
		log.Println("sent", n)
	}
	go f(1)
	go f(2)
	time.Sleep(time.Second)
	log.Println("to recv")
	<-bch
	log.Println("recved")
}

// 测试先收, select发
func bchanEx2() {
	bch := make(chan int)
	f := func (n int) {
		time.Sleep(time.Second)
		log.Println("to send", n)
		select {
		case bch <- n:
			log.Println("sent", n)
		default:
			log.Println("send failed", n)
		}
	}
	go f(2)
	go f(1)
	log.Println("to recv")
	log.Println(<-bch, "recved")
}

// 测试先发，select发
func bchanEx3() {
	bch := make(chan int)
	f := func (n int) {
		log.Println("to send", n)
		select {
		case bch <- n:
			log.Println("sent", n)
		default:
			log.Println("send failed", n)
		}
	}
	go f(2)
	go f(1)
	time.Sleep(time.Second)
	log.Println("to recv")
	log.Println(<-bch, "recved")
}

// 测试先收，select收，发
func bchanEx4() {
	bch := make(chan int)
	f := func (n int) {
		time.Sleep(time.Second)
		log.Println("to send", n)
		bch <- n
		log.Println("sent", n)
	}
	go f(1)
	log.Println("to recv")
	select {
	case i := <-bch:
		log.Println("recved", i)
	default:
		log.Println("recv failed")
	}
}

// 测试先收，select收，不发
func bchanEx5() {
	bch := make(chan int)
	log.Println("to recv")
	select {
	case i := <-bch:
		log.Println("recved", i)
	default:
		log.Println("recv failed")
	}
}

// 测试先发
func fnbchanEx1() {
	fnbch := make(chan int, 1)
	fnbch <- 0
	go func () {
		log.Println("to send")
		fnbch <- 1
		log.Println("sent")
	}()
	time.Sleep(time.Second)
	log.Println("to recv")
	<-fnbch
	log.Println("recved")
	log.Println("to recv again")
	<-fnbch
	log.Println("recved again")
}

// 测试先收, select发
func fnbchanEx2() {
	fnbch := make(chan int, 1)
	fnbch <- 0
	f := func (n int) {
		time.Sleep(time.Second)
		log.Println("to send", n)
		select {
		case fnbch <- n:
			log.Println("sent", n)
		default:
			log.Println("send failed", n)
		}
	}
	go f(2)
	go f(1)
	log.Println("to recv")
	log.Println(<-fnbch, "recved")
	log.Println("to recv again")
	log.Println(<-fnbch, "recved again")
}

// 测试先发, select发
func fnbchanEx3() {
	fnbch := make(chan int, 1)
	f := func (n int) {
		log.Println("to send", n)
		select {
		case fnbch <- n:
			log.Println("sent", n)
		default:
			log.Println("send failed", n)
		}
	}
	go f(2)
	go f(1)
	time.Sleep(time.Second)
	log.Println("to recv")
	log.Println(<-fnbch, "recved")
}

// // 测试发之后收
// func nbchanEx1() {
//         nbch := make(chan int, 1)
//         f := func (n int) {
//                 log.Println("to send", n)
//                 nbch <- n
//                 log.Println("sent", n)
//         }
//         go f(1)
//         go f(2)
//         time.Sleep(time.Second)
//         log.Println("to recv")
//         <-nbch
//         log.Println("recved")
// }

