use std::io;
use std::io::BufRead;
use std::fs::File;

fn star1() {
    let mut stacks : Vec<Vec<char>> = Vec::new();
    
    let mut readstacks = true;
    for line in io::BufReader::new(File::open("05.txt").unwrap()).lines().flat_map(|x|x) {
        if readstacks {
            for i in 0.. {
                let char_index = 1+i*4;
                if char_index >= line.len() { break; }
                let c = line.chars().nth(char_index).unwrap();
                if c.is_numeric() { break; }
                if c != ' ' {
                    while i >= stacks.len() {
                        stacks.push(Vec::new());
                    }
                    stacks[i].push(c);
                }
                //println!("{i} = {char_index} = {c}");
            }
            if line == "" {
                readstacks = false;
                for i in 0..stacks.len() {
                    stacks[i].reverse();
                }
                //println!("{:?}", stacks);
            }
            continue;
        }
        let split : Vec<&str> = line.split(" ").collect();
        let amount : i32 = split[1].parse().unwrap();
        let from : usize = split[3].parse().unwrap();
        let to : usize = split[5].parse().unwrap();
        

        let mut moved : Vec<char> = Vec::new();
        for _ in 0..amount {
            moved.push(stacks[from-1].pop().unwrap());
        }
        stacks[to-1].append(&mut moved);
        //println!("{from} -> {to} ({amount})");
        //println!("{:?}", stacks);
    }

    let mut tops : Vec<char> = Vec::new();
    for i in 0..stacks.len() {
        if !stacks[i].is_empty() {
            tops.push(stacks[i].pop().unwrap());
        }
    }

    println!("{}", tops.iter().collect::<String>());
}

fn star2() {
    let mut stacks : Vec<Vec<char>> = Vec::new();
    
    let mut readstacks = true;
    for line in io::BufReader::new(File::open("05.txt").unwrap()).lines().flat_map(|x|x) {
        if readstacks {
            for i in 0.. {
                let char_index = 1+i*4;
                if char_index >= line.len() { break; }
                let c = line.chars().nth(char_index).unwrap();
                if c.is_numeric() { break; }
                if c != ' ' {
                    while i >= stacks.len() {
                        stacks.push(Vec::new());
                    }
                    stacks[i].push(c);
                }
                //println!("{i} = {char_index} = {c}");
            }
            if line == "" {
                readstacks = false;
                for i in 0..stacks.len() {
                    stacks[i].reverse();
                }
                //println!("{:?}", stacks);
            }
            continue;
        }
        let split : Vec<&str> = line.split(" ").collect();
        let amount : i32 = split[1].parse().unwrap();
        let from : usize = split[3].parse().unwrap();
        let to : usize = split[5].parse().unwrap();
        

        let mut moved : Vec<char> = Vec::new();
        for _ in 0..amount {
            moved.push(stacks[from-1].pop().unwrap());
        }
        moved.reverse();
        stacks[to-1].append(&mut moved);
        //println!("{from} -> {to} ({amount})");
        //println!("{:?}", stacks);
    }

    let mut tops : Vec<char> = Vec::new();
    for i in 0..stacks.len() {
        if !stacks[i].is_empty() {
            tops.push(stacks[i].pop().unwrap());
        }
    }

    println!("{}", tops.iter().collect::<String>());
}

fn main() {
    star1();
    star2();
}
