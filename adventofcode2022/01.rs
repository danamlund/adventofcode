use std::fs::File;
use std::io;
use std::io::BufRead;

fn star1() {
   let mut sum : i64 = 0;
   let mut maxsum = sum;
   for line in io::BufReader::new(File::open("01.txt").expect("")).lines().flat_map(|l| l) {
      if line == "" {
        if sum > maxsum { maxsum = sum; }
        sum = 0;
      } else {
        let calories : i64 = line.parse().expect("");
        sum += calories;
      }
   }

   println!("= {}", maxsum);
}

fn star2() {
   let mut sum : i64 = 0;
   let mut top3 : Vec<i64> = Vec::new();
   for line in io::BufReader::new(File::open("01.txt").expect("")).lines().flat_map(|l| l) {
      if line == "" {
        top3.push(sum);
        if top3.len() > 3 {
          top3.sort_by(|a,b| b.cmp(a));
          top3.resize(3, 0);
        }
        sum = 0;
      } else {
        let calories : i64 = line.parse().expect("");
        sum += calories;
      }
   }

   let sum : i64 = top3.iter().sum();
   println!("= {}", sum);
   //println!("= {}", top3.iter().sum::<i64>()); // top from others
}

fn main() {
   star1();
   star2();
}
