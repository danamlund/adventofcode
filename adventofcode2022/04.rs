use std::io;
use std::io::BufRead;
use std::fs::File;

fn star1() {
    let mut sum : i32 = 0;
    for line in io::BufReader::new(File::open("04.txt").unwrap()).lines().flat_map(|x|x) {
        let mut split = line.split(|c:char| !c.is_numeric());
        let a1 : i32 = split.next().unwrap().parse().unwrap();
        let a2 : i32 = split.next().unwrap().parse().unwrap();
        let b1 : i32 = split.next().unwrap().parse().unwrap();
        let b2 : i32 = split.next().unwrap().parse().unwrap();

        if a1 >= b1 && a2 <= b2 || b1 >= a1 && b2 <= a2 { sum += 1; }
        
        //println!("{a1}..{a2} {b1}..{b2}");
    }
    println!("= {sum}");
    // 607 too high
}

fn star2() {
    let mut sum : i32 = 0;
    for line in io::BufReader::new(File::open("04.txt").unwrap()).lines().flat_map(|x|x) {
        let mut split = line.split(|c:char| !c.is_numeric());
        let a1 : i32 = split.next().unwrap().parse().unwrap();
        let a2 : i32 = split.next().unwrap().parse().unwrap();
        let b1 : i32 = split.next().unwrap().parse().unwrap();
        let b2 : i32 = split.next().unwrap().parse().unwrap();

        if !(a2 < b1 || b2 < a1) { sum += 1; }
        
        //println!("{a1}..{a2} {b1}..{b2}");
    }
    println!("= {sum}");
}

fn main() {
    star1();
    star2();
}
