use std::io;
use std::fs::File;
use std::io::BufRead;
use std::collections::HashMap;
use std::collections::HashSet;

fn star1() {
    let mut sum : i64 = 0;
    for line in io::BufReader::new(File::open("03.txt").unwrap()).lines().flat_map(|x|x) {
        //println!("{}", line);

        let pocket1 = line.get(..line.len()/2).unwrap();
        let pocket2 = line.get(line.len()/2..).unwrap();

        for p1 in pocket1.chars() {
            if pocket2.contains(p1) {
                //println!("!{}", p1);
                //sum += i64::max(0, p1 as i64 - 'a' as i64) + i64::max(0, p1 as i64 - 'A' as i64);
                //let sum0 = sum;
                sum += match p1 { b if 'a' <= b && b <= 'z' => 1 + b as i64 - 'a' as i64,
                                  b if 'A' <= b && b <= 'Z' => 27 + b as i64 - 'A' as i64,
                                  _ => panic!()};
                //println!("!{} = {}", p1, sum-sum0);
                break;
            }
        }
        
        //println!("{} {}", pocket1, pocket2);
    }

    println!("= {}", sum);
}

fn star2() {
    let mut sum : i64 = 0;
    let mut bags = Vec::new();
    for line in io::BufReader::new(File::open("03.txt").unwrap()).lines().flat_map(|x|x) {
        bags.push(line.clone());
        if bags.len() != 3 { continue; }
        //println!("{} {} {}", bags[0], bags[1], bags[2]);
        //println!("{}", line);

        let mut counts = HashMap::new();
        for bag in bags.iter() {
            for letter in bag.chars().collect::<HashSet<char>>() {
                counts.insert(letter.clone(), 1 + counts.get(&letter).unwrap_or(&0));
            }
        }

        for (k,v) in counts {
            if v == 3 {
                sum += match k { b if 'a' <= b && b <= 'z' => 1 + b as i64 - 'a' as i64,
                                 b if 'A' <= b && b <= 'Z' => 27 + b as i64 - 'A' as i64,
                                 _ => panic!()};
                //println!("{}", k);
            }
        }
        
        bags.clear();
    }

    println!("= {}", sum);
}

fn main() {
    star1();
    star2();
}
